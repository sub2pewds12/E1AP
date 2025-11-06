import os
import re
import logging
import shutil
from asn1_parser import ASN1Parser 
from asn1_patches import ASN1Patcher
from typing import Dict, List, Any, Tuple, Optional
from go_templates import render_sequence_struct, render_choice_struct, render_enum_struct, render_extension_struct_only, render_list_struct
from method_templates import render_pdu_methods, render_internal_struct_methods, render_enum_methods, render_list_methods, render_extension_methods
from common_types import (
    ASN1Procedure,
    EnumDefinition,
    SequenceDefinition,
    ChoiceDefinition,
    AliasDefinition,
    StringDefinition,
    IntegerDefinition,
    ConstantDefinition,
    ListDefinition,
    BuiltinDefinition,
)

logger = logging.getLogger(__name__)


class GoCodeGenerator:
    def __init__(
        self,
        definitions: Dict[str, Any],
        failures: List[Dict[str, Any]],
        procedures: Dict[str, ASN1Procedure],
        message_to_procedure_map: Dict[str, str],
        output_dir: str,
        patcher: "ASN1Patcher",
        acronyms: List[str],      # Correctly placed before parser
        parser: "ASN1Parser"      # Correctly placed at the end
    ):
        self.definitions = definitions
        self.failures = failures
        self.procedures = procedures
        self.message_to_procedure_map = message_to_procedure_map
        self.output_dir = output_dir
        self.patcher = patcher
        self.acronyms = set(acronyms) # This line will now work correctly
        self.parser = parser         # This line will now work correctly
        self.generated_files = set()
        self.name_overrides = {
            "UE-associatedLogicalE1-ConnectionListResItem": "UEAssociatedLogicalE1ConnectionItemRes",
            "UE-associatedLogicalE1-ConnectionListResAckItem": "UEAssociatedLogicalE1ConnectionItemResAck",
        }
        self.parser.pascal_case_converter = self._standard_string
        self.parser.go_type_resolver = self._resolve_go_type

    def _standard_string(self, input_string: str) -> str:
        if not input_string:
            return ""
        if input_string in self.name_overrides:
            return self.name_overrides[input_string]
        
        overrides = {
            "ProtocolIE-SingleContainer": "ProtocolIESingleContainer",
            "ProtocolIE-Container": "ProtocolIEContainer",
            "ProtocolExtensionContainer": "ProtocolExtensionContainer",
            "OBJECT IDENTIFIER": "ObjectIdentifier",
            "CriticalityDiagnostics-IE-List": "CriticalityDiagnosticsIEList",
        }
        if input_string in overrides:
            return overrides[input_string]

        s = re.sub(r"^(id-)", "", input_string)
        parts = re.split(r"[-_\s]", s)

        processed_parts = []
        for part in parts:
            if not part:
                continue
            if part.upper() in self.acronyms: 
                processed_parts.append(part.upper())
            else:
                processed_parts.append(part[0].upper() + part[1:])

        return "".join(processed_parts)


    def _format_go_value(self, value: str) -> str:
        if not value:
            return "0"
        try:
            return str(int(value))
        except (ValueError, TypeError):
            return self._standard_string(value)

    def _go_name_to_snake_case(self, go_name: str) -> str:
        """Converts a PascalCase string to snake_case for filenames."""
        s1 = re.sub("(.)([A-Z][a-z]+)", r"\1_\2", go_name)
        return re.sub("([a-z0-9])([A-Z])", r"\1_\2", s1).lower()

    def _classify_definition(self, item: Any) -> str:
        """Classifies a definition as 'constant', 'simple', or 'complex'."""

        if "{" in item.name or "}" in item.name:
            return "unknown"

        if isinstance(item, ConstantDefinition):
            return "constant"

        if isinstance(
            item, (SequenceDefinition, ChoiceDefinition, ListDefinition, EnumDefinition)
        ):
            return "complex"

        if isinstance(item, (IntegerDefinition, StringDefinition, AliasDefinition)):
            return "simple"

        return "unknown"

    def _resolve_go_type(self, type_name: str) -> Tuple[str, Optional[Any]]:
        """
        Resolves an ASN.1 type name to its final, concrete Go type and returns
        the concrete ASN1Definition object.
        """
        manual_types = [
            "ProtocolExtensionContainer",
            "PrivateIEContainer",
            "ProtocolIESingleContainer",
            "ProtocolIEField",
            "PrivateIEField",
            "ProtocolExtensionField",
        ]
        if type_name in manual_types:
            return type_name, None

        if type_name == "ANY":
            return "aper.OctetString", self.definitions.get("OCTET STRING")

        container_mappings = {
            "ProtocolExtensionContainer": "ProtocolExtensionContainer",
            "PrivateIEContainer": "PrivateIEContainer",
            "ProtocolIESingleContainer": "ProtocolIESingleContainer",
        }
        if type_name in container_mappings:
            go_type = container_mappings[type_name]

            dummy_def = BuiltinDefinition(name=type_name, source_file="manual")
            return go_type, dummy_def

        definition = self.definitions.get(type_name)
        if not definition:
            return self._standard_string(type_name), None
        if isinstance(definition, (IntegerDefinition, StringDefinition)) and \
            not isinstance(definition, BuiltinDefinition) and \
            not isinstance(definition, ConstantDefinition):
                return self._standard_string(definition.name), definition

        visited = {type_name}

        while isinstance(definition, AliasDefinition):
            next_type_name = definition.alias_of
            
            # Infinite loop prevention
            if next_type_name in visited:
                logger.warning(f"Circular alias dependency detected for '{type_name}' at '{next_type_name}'. Breaking.")
                # Fallback to just returning the last known type name
                return self._standard_string(next_type_name), None
            visited.add(next_type_name)

            # Get the next definition in the chain
            next_def = self.definitions.get(next_type_name)
            if not next_def:
                logger.warning(f"Alias '{next_type_name}' for type '{definition.name}' not found. Breaking chain.")
                # Fallback to returning the last known type name
                return self._standard_string(next_type_name), None
            
            # Update the definition for the next loop iteration
            definition = next_def

        if isinstance(definition, (IntegerDefinition, ConstantDefinition)):
            return "aper.Integer", definition
        
        elif (
            isinstance(definition, StringDefinition)
            and definition.string_type == "BIT STRING"
        ):
            return "aper.BitString", definition
        
        elif isinstance(definition, StringDefinition):
            return "aper.OctetString", definition
        
        elif isinstance(definition, BuiltinDefinition):
            if definition.name in ["VisibleString", "PrintableString", "UTF8String"]:
                return "aper.OctetString", self.definitions.get("OCTET STRING")
            if definition.name == "OBJECT IDENTIFIER":
                return "string", definition
            if definition.name == "NULL":
                return "bool", definition
        elif isinstance(
            definition,
            (EnumDefinition, SequenceDefinition, ChoiceDefinition, ListDefinition),
        ):
            return self._standard_string(definition.name), definition

        logger.warning(
            f"Could not resolve a specific Go type for ASN.1 type '{type(definition).__name__}'. Defaulting to its name."
        )
        return self._standard_string(definition.name), definition

    def _generate_header_comment(self, item: Any, go_name: str) -> str:
        """Builds the standard comment block for a generated Go definition."""
        header = f"// {go_name} From: {item.source_file}:{item.source_line}\n"
        if hasattr(item, "type"):
            header += f"// ASN.1 Data Type: {item.type}\n"
        return header

    def generate_files(self):
        os.makedirs(self.output_dir, exist_ok=True)

        classified_defs = {"constant": [], "simple": [], "complex": []}

        for item in self.definitions.values():
            if not isinstance(item, dict):
                classification = self._classify_definition(item)
                if classification != "unknown":
                    classified_defs[classification].append(item)

        self._generate_constants_file(classified_defs["constant"])
        self._generate_common_types_file(classified_defs["simple"])
        self._generate_complex_type_files(classified_defs["complex"])
        
        stale_files_found = 0
        for filename_on_disk in os.listdir(self.output_dir):
            if not filename_on_disk.endswith(".go"):
                continue

            if (
                filename_on_disk not in self.generated_files
                and filename_on_disk not in ["e1ap_manual_types.go", "e1ap_pdu.go"]
            ):
                file_to_delete = os.path.join(self.output_dir, filename_on_disk)
                os.remove(file_to_delete)
                logger.info(f"Cleaned up stale generated file: {filename_on_disk}")
                stale_files_found += 1

        if stale_files_found == 0:
            logger.info("No stale generated files to clean up.")

        self.patcher.patch_generated_files(self.output_dir)
        self._format_generated_code()

    def _format_generated_code(self):
        """Runs gofmt and goimports on the output directory."""
        logger.info("Formatting generated Go code...")
        try:

            os.system(f"goimports -w {self.output_dir}")
            os.system(f"gofmt -w {self.output_dir}")
            logger.info("Formatting complete.")
        except Exception as e:
            logger.warning(
                f"Could not format Go code. Please run 'goimports -w' and 'gofmt -w' on the '{self.output_dir}' directory manually. Error: {e}"
            )

    def _generate_constants_file(self, constant_items: List[Any]):
        """Generates a single file for all constant definitions."""

        if not constant_items:
            logger.info("No constants found to generate.")
            return

        integer_consts, proc_code_consts, protocol_ie_consts = [], [], []
        proc_code_type_name = self._standard_string("ProcedureCode")
        protocol_ie_id_type_name = self._standard_string("ProtocolIE-ID")

        for item in constant_items:
            # Use the existing pascal-case converter
            base_name = self._standard_string(item.name)
            
            # Clean the "Id" prefix that comes from ASN.1 names like "id-e1Setup"
            if base_name.startswith("Id"):
                base_name = base_name[2:]

            value = self._format_go_value(item.min_val)

            # Check the ASN.1 type of the constant to categorize it
            if item.const_type == "ProcedureCode":
                # This will generate names like: ProcedureCodeE1Setup
                proc_code_consts.append((f"ProcedureCode{base_name}", value))
            elif item.const_type == "ProtocolIE-ID":
                # This will generate names like: ProtocolIEIDGNBCUCPUEE1APID
                protocol_ie_consts.append((f"ProtocolIEID{base_name}", value))
            else:
                # For all other integer constants
                integer_consts.append((base_name, value))

        sorter = lambda item: (
            int(item[1]) if item[1] and item[1].isdigit() else float("inf")
        )

        integer_consts.sort(key=sorter)
        proc_code_consts.sort(key=sorter)
        protocol_ie_consts.sort(key=sorter)

        go_code = ""

        if integer_consts:
            go_code += "const (\n"
            max_len = max(len(name) for name, _ in integer_consts)
            for name, value in integer_consts:
                padding = " " * (max_len - len(name))
                go_code += f"\t{name}{padding} int64 = {value}\n"
            go_code += ")\n\n"

        if proc_code_consts:
            go_code += "const (\n"
            max_len = max(len(name) for name, _ in proc_code_consts)
            for name, value in proc_code_consts:
                padding = " " * (max_len - len(name))
                go_code += f"\t{name}{padding} {proc_code_type_name} = {value}\n"
            go_code += ")\n\n"

        if protocol_ie_consts:
            go_code += "const (\n"
            max_len = max(len(name) for name, _ in protocol_ie_consts)
            for name, value in protocol_ie_consts:
                padding = " " * (max_len - len(name))
                go_code += f"\t{name}{padding} {protocol_ie_id_type_name} = {value}\n"
            go_code += ")\n\n"

        total_consts = (
            len(integer_consts) + len(proc_code_consts) + len(protocol_ie_consts)
        )
        if total_consts > 0:
            file_path = os.path.join(self.output_dir, "e1ap_constants.go")
            with open(file_path, "w", encoding="utf-8") as f:
                f.write(f"package e1ap_ies\n\n{go_code}")
            self.generated_files.add("e1ap_constants.go")
            logger.info(
                f"SUCCESS: Wrote {total_consts} constants to 'e1ap_constants.go'."
            )
        else:
            logger.info("No constants found to generate.")

    def _generate_common_types_file(self, simple_items: List[Any]):
        """
        Generates a single file for simple, reusable types as structs with Encode/Decode methods.
        """
        if not simple_items:
            logger.info("No simple types found to generate in common file.")
            return

        go_code = 'package e1ap_ies\n\nimport (\n\t"fmt"\n\t"io"\n\t"github.com/lvdund/ngap/aper"\n)\n\n'

        sorted_items = sorted(simple_items, key=lambda x: self._standard_string(x.name))
        total_count = 0

        for item in sorted_items:
            go_name = self._standard_string(item.name)
            
            underlying_type = ""
            encode_logic = ""
            decode_logic = ""
            
            if isinstance(item, IntegerDefinition) or (isinstance(item, AliasDefinition) and "INTEGER" in item.alias_of.upper()):
                underlying_type = "aper.Integer"

                # This is the corrected, robust logic from your code
                min_val_str = getattr(item, 'min_val', None)
                max_val_str = getattr(item, 'max_val', None)

                min_val = self._format_go_value(min_val_str)
                max_val = self._format_go_value(max_val_str)
                
                is_extensible = str(getattr(item, 'is_extensible', False)).lower()

                encode_logic = f"return w.WriteInteger(int64(s.Value), &aper.Constraint{{Lb: {min_val}, Ub: {max_val}}}, {is_extensible})"
                decode_logic = f"""
        val, err := r.ReadInteger(&aper.Constraint{{Lb: {min_val}, Ub: {max_val}}}, {is_extensible})
        if err != nil {{
            return err
        }}
        s.Value = aper.Integer(val)
        return nil"""

            elif isinstance(item, StringDefinition) or (isinstance(item, AliasDefinition) and "STRING" in item.alias_of.upper()):
                # Applying the same corrected, robust logic here
                min_val_str = getattr(item, 'min_val', None)
                max_val_str = getattr(item, 'max_val', None)

                min_val = self._format_go_value(min_val_str)
                max_val = self._format_go_value(max_val_str)

                is_extensible = str(getattr(item, 'is_extensible', False)).lower()
                string_type_name = item.string_type if isinstance(item, StringDefinition) else item.alias_of
                
                if "BIT STRING" in string_type_name.upper():
                    underlying_type = "aper.BitString"
                    encode_logic = f"return w.WriteBitString(s.Value.Bytes, uint(s.Value.NumBits), &aper.Constraint{{Lb: {min_val}, Ub: {max_val}}}, {is_extensible})"
                    decode_logic = f"""
        var err error
        s.Value.Bytes, s.Value.NumBits, err = r.ReadBitString(&aper.Constraint{{Lb: {min_val}, Ub: {max_val}}}, {is_extensible})
        return err"""
                else: # OCTET STRING
                    underlying_type = "aper.OctetString"
                    encode_logic = f"return w.WriteOctetString([]byte(s.Value), &aper.Constraint{{Lb: {min_val}, Ub: {max_val}}}, {is_extensible})"
                    decode_logic = f"""
        var err error
        s.Value, err = r.ReadOctetString(&aper.Constraint{{Lb: {min_val}, Ub: {max_val}}}, {is_extensible})
        return err"""

            if underlying_type:
                total_count += 1
                go_code += self._generate_header_comment(item, go_name)
                go_code += f"type {go_name} struct {{\n\tValue {underlying_type}\n}}\n\n"
                go_code += f"// Encode implements the aper.AperMarshaller interface.\n"
                go_code += f"func (s *{go_name}) Encode(w *aper.AperWriter) error {{\n\t{encode_logic}\n}}\n\n"
                go_code += f"// Decode implements the aper.AperUnmarshaller interface.\n"
                go_code += f"func (s *{go_name}) Decode(r *aper.AperReader) error {{\n\t{decode_logic}\n}}\n\n"
                
        if total_count > 0:
            file_path = os.path.join(self.output_dir, "e1ap_common_types.go")
            with open(file_path, "w", encoding="utf-8") as f:
                f.write(go_code)
            self.generated_files.add("e1ap_common_types.go")

        logger.info(f"SUCCESS: Wrote {total_count} common type definitions to 'e1ap_common_types.go'.")

    def _generate_complex_type_files(self, complex_items: List[Any]):
       """Generates a dedicated file for each complex type using a robust two-pass strategy."""
       if not complex_items:
           logger.info("No complex types found to generate individual files for.")
           return


       type_counts = {"SEQUENCE": 0, "CHOICE": 0, "ENUMERATED": 0, "LIST": 0}
        
       files_to_write = {}

       logger.info(
           "Complex types pass 1: Generating SEQUENCE, CHOICE, and ENUMERATED files..."
       )
       for item in complex_items:
           if isinstance(item, SequenceDefinition):
                for member in item.ies:
                    if member.type == "ProtocolExtensionContainer" and hasattr(member, 'extension_set_name'):
                        extension_set_name = member.extension_set_name
                        extension_set = self.parser.extension_sets.get(extension_set_name)
                        
                        if extension_set:
                            # 1. Determine names
                            ext_go_name = self._standard_string(item.name) + "Extensions"
                            ext_filename = ext_go_name + ".go"
                            ext_filepath = os.path.join(self.output_dir, ext_filename)

                            # 2. Generate the struct code
                            ext_struct_code = render_extension_struct_only(
                                go_name=ext_go_name,
                                extension_set=extension_set,
                                pascal_case_converter=self._standard_string,
                            )
                            
                            # 3. Generate the method code using the template
                            ext_method_code, ext_imports = render_extension_methods(ext_go_name)
                            
                            # 4. Assemble the full file content, including imports, struct, and methods
                            ext_file_content = "package e1ap_ies\n\n"
                            if ext_imports:
                                ext_file_content += "import (\n"
                                for imp in sorted(list(ext_imports)):
                                    ext_file_content += f'\t"{imp}"\n'
                                ext_file_content += ")\n\n"
                            ext_file_content += f"{ext_struct_code}\n\n{ext_method_code}"

                            # 5. Write the complete file directly to disk
                            with open(ext_filepath, "w", encoding="utf-8") as f:
                                f.write(ext_file_content)

                            self.generated_files.add(ext_filename)
                            logger.info(f"SUCCESS: Wrote type-safe extensions file to '{ext_filename}'.")
                            
                            # We only need one extension struct per sequence, so we can stop looking.
                            break


           if isinstance(item, (SequenceDefinition, ChoiceDefinition, EnumDefinition, ListDefinition)):
               go_name = self._standard_string(item.name)
               if go_name in ["InitiatingMessage", "SuccessfulOutcome", "UnsuccessfulOutcome", "E1APPDU"]:
                   logger.info(f"Skipping generation for abstract PDU wrapper: {go_name}")
                   continue
               filename = go_name + ".go"
               file_path = os.path.join(self.output_dir, filename)
              
               # Step 1: Generate the struct (your _generate_struct_for_item helper
               # must now be able to handle ListDefinition).
               struct_code = self._generate_struct_for_item(item)
              
               # Step 2: Generate the methods
               method_code, required_imports = "", set()
               pdu_suffixes = [
                   "Request", "Response", "Failure", "Command", "Complete",
                   "Indication", "Acknowledge", "Setup", "Transfer", "Notification",
                   "PDU" # for E1APPDU itself
               ]
               is_pdu = item.name in self.message_to_procedure_map
              
               if is_pdu:
                   method_code, pdu_imports = render_pdu_methods(go_name, item, self.parser, self.message_to_procedure_map, self.procedures)
                   required_imports.update(pdu_imports)
               elif isinstance(item, (SequenceDefinition, ChoiceDefinition)):
                   method_code, internal_imports = render_internal_struct_methods(go_name, item, self.parser)
                   required_imports.update(internal_imports)
               elif isinstance(item, EnumDefinition):
                   method_code, enum_imports = render_enum_methods(go_name, item)
                   required_imports.update(enum_imports)
               elif isinstance(item, ListDefinition): # <-- This new block handles lists
                   method_code, list_imports = render_list_methods(go_name, item, self.parser)
                   required_imports.update(list_imports)
              
               # Step 3: Assemble and write
               file_content = "package e1ap_ies\n\n"
               if required_imports:
                   file_content += "import (\n"
                   for imp in sorted(list(required_imports)):
                       file_content += f'\t"{imp}"\n'
                   file_content += ")\n\n"
               file_content += f"{struct_code}\n\n{method_code}"
              
               with open(file_path, "w", encoding="utf-8") as f:
                   f.write(file_content)
              
               self.generated_files.add(filename)
               item_type_str = type(item).__name__.replace("Definition", "").upper()
               if item_type_str in type_counts:
                   type_counts[item_type_str] += 1
          
       total_complex_count = sum(type_counts.values())
       if total_complex_count > 0:
           for def_type, count in type_counts.items():
               if count > 0:
                   logger.info(
                       f"SUCCESS: Wrote {count} {def_type} definitions to individual files."
                   )

    def _generate_struct_for_item(self, item: Any) -> str:
        go_name = self._standard_string(item.name)
        if isinstance(item, SequenceDefinition):
            return render_sequence_struct(go_name, item, self._standard_string, self._resolve_go_type, self._format_go_value, self.parser.extension_sets)
        if isinstance(item, ChoiceDefinition):
            return render_choice_struct(go_name, item, self._standard_string, self._resolve_go_type)
        if isinstance(item, EnumDefinition):
            return render_enum_struct(go_name, item.enum_values, self._standard_string)
        if isinstance(item, ListDefinition):
            # This is the new part for the helper, make sure it's here.
            return render_list_struct(go_name, item, self._standard_string)
        return ""


    def run_full_diagnostic(self):
        print("\n" + "=" * 80)
        print(" " * 25 + "ASN.1 PARSER DIAGNOSTIC REPORT")
        print("=" * 80)

        successes = {
            "CONSTANT": [],
            "INTEGER": [],
            "ENUMERATED": [],
            "SEQUENCE": [],
            "CHOICE": [],
            "LIST": [],
            "BIT STRING": [],
            "OCTET STRING": [],
            "PrintableString": [],
            "VisibleString": [],
            "UTF8String": [],
            "NULL": [],
        }

        for item in self.definitions.values():
            if item.is_constant:
                successes["CONSTANT"].append(item)
            elif item.type in successes:
                successes[item.type].append(item)

        total_success_count = 0

        report_format = "  {context:<50} {name}"

        for def_type in [
            "CONSTANT",
            "INTEGER",
            "ENUMERATED",
            "BIT STRING",
            "OCTET STRING",
            "PrintableString",
            "VisibleString",
            "UTF8String",
            "NULL",
            "SEQUENCE",
            "CHOICE",
            "LIST",
        ]:
            items = successes.get(def_type)
            if items:
                sorted_items = sorted(items, key=lambda x: x.name)
                print(
                    f"\n--- [ SUCCESS ] PARSED {def_type} DEFINITIONS ({len(sorted_items)}) ---"
                )
                for item in sorted_items:
                    context = f"({item.source_file}:{item.source_line})"
                    print(report_format.format(context=context, name=item.name))
                total_success_count += len(sorted_items)

        print("\n" + "-" * 40)
        print(
            f"TOTAL SUCCESSFUL DEFINITIONS: {total_success_count} (out of {len(self.definitions)} total)"
        )
        print("-" * 40)

        print(
            f"\n--- [ FAILURE ] SKIPPED ABSTRACT/UNHANDLED DEFINITIONS ({len(self.failures)}) ---"
        )
        sorted_failures = sorted(self.failures, key=lambda x: x["name"])
        for failure in sorted_failures:
            context = f"({failure['file']}:{failure['line']})"
            print(report_format.format(context=context, name=failure["name"]))

        print("\n" + "=" * 80)
        print(
            "ACTION REQUIRED: Verify the [ SUCCESS ] counts against your manual analysis."
        )
        print(
            "The [ FAILURE ] list should only contain abstract schemas like E1AP-PROTOCOL-IES."
        )
        print("=" * 80)
