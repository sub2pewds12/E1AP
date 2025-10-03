import os
import re
import logging
import shutil
from typing import Dict, List, Any, Tuple, Optional
from common_types import ASN1Procedure 

logger = logging.getLogger(__name__)


class GoCodeGenerator:
    def __init__(
        self,
        definitions: Dict[str, Any],
        failures: List[Dict[str, Any]],
        procedures: Dict[str, ASN1Procedure],
        message_to_procedure_map: Dict[str, str],
        output_dir: str,
    ):
        self.definitions = definitions
        self.failures = failures
        self.procedures = procedures
        self.message_to_procedure_map = message_to_procedure_map
        self.output_dir = output_dir
        self.generated_files = set()

    def _standard_string(self, input_string: str) -> str:
        if not input_string:
            return ""
        
        overrides = {
            "ProtocolIE-SingleContainer": "ProtocolIESingleContainer",
            "ProtocolIE-Container": "ProtocolIEContainer",
            "ProtocolExtensionContainer": "ProtocolExtensionContainer",
            "OBJECT IDENTIFIER": "ObjectIdentifier",
            "CriticalityDiagnostics-IE-List": "CriticalityDiagnosticsIEList", 
        }
        if input_string in overrides:
            return overrides[input_string]
        
        known_acronyms = {
            "ID", "IE", "PDU", "E1AP", "GNB", "CU", "CP", "UP", "QOS",
            "DRB", "IP", "ASN", "TEID", "URI", "SCTP", "GTP"
        }

        s = re.sub(r"^(id-)", "", input_string)
        parts = re.split(r"[-_\s]", s)

        processed_parts = []
        for part in parts:
            if not part:
                continue
            if part.upper() in known_acronyms:
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
        s1 = re.sub('(.)([A-Z][a-z]+)', r'\1_\2', go_name)
        return re.sub('([a-z0-9])([A-Z])', r'\1_\2', s1).lower()
    
    def _classify_definition(self, item: Any) -> str:
        """Classifies a definition as 'constant', 'simple', or 'complex'."""
        if isinstance(item, dict):
            return 'unknown'

        if item.is_constant:
            return 'constant'
        
        if item.type in ["SEQUENCE", "CHOICE", "LIST", "ENUMERATED"]:
            return 'complex'
            
        if item.type in ["INTEGER", "BIT STRING", "OCTET STRING", "ALIAS", "BASE_TYPE", "PrintableString", "VisibleString", "UTF8String", "NULL", "ANY"]:
            return 'simple'
            
        return 'unknown'
    
    def _resolve_go_type(self, type_name: str) -> Tuple[str, Optional[Any]]:
        """
        Resolves an ASN.1 type name to its final, concrete Go type and returns
        the concrete ASN1Definition object.
        """
        if type_name == "ANY":
            return "aper.OctetString", self.definitions.get("ANY")

        definition = self.definitions.get(type_name)

        if not definition or isinstance(definition, dict):
            return self._standard_string(type_name), None

        visited = {type_name}
        while definition and definition.alias_of:
            next_alias = definition.alias_of
            if next_alias in visited:
                logger.warning(f"Circular alias dependency detected for '{type_name}' at '{next_alias}'.")
                break
            visited.add(next_alias)
            next_def = self.definitions.get(next_alias)
            if not next_def or isinstance(next_def, dict):
                logger.warning(f"Alias '{next_alias}' for type '{definition.name}' not found in definitions.")
                break
            definition = next_def

        if definition.type in ["INTEGER", "BASE_TYPE"]:
            return "aper.Integer", definition
        elif definition.type == "BIT STRING":
            return "aper.BitString", definition
        elif definition.type in ["OCTET STRING", "PrintableString", "VisibleString", "UTF8String", "ANY"]:
            return "aper.OctetString", definition
        elif definition.type == "OBJECT IDENTIFIER": 
            return "string", definition
        elif definition.type == "NULL":
            return "bool", definition 
        elif definition.type in ["ENUMERATED", "SEQUENCE", "CHOICE", "LIST"]:
            return self._standard_string(definition.name), definition
        
        logger.warning(f"Could not resolve a specific Go type for ASN.1 type '{definition.type}'. Defaulting to its name.")
        return self._standard_string(definition.name), definition

    def _generate_header_comment(self, item: Any, go_name: str) -> str:
        """Builds the standard comment block for a generated Go definition."""
        header = f"// {go_name} From: {item.source_file}:{item.source_line}\n"
        if hasattr(item, 'type'):
             header += f"// ASN.1 Data Type: {item.type}\n"
        return header
    
    def generate_files(self):
        os.makedirs(self.output_dir, exist_ok=True)

        classified_defs = { 'constant': [], 'simple': [], 'complex': [] }
        
        # Classify all valid definitions, ignoring metadata dictionaries
        for item in self.definitions.values():
            if not isinstance(item, dict):
                classification = self._classify_definition(item)
                if classification != 'unknown':
                    classified_defs[classification].append(item)

        # --- Generation Stages ---
        self._generate_constants_file(classified_defs['constant'])
        self._generate_common_types_file(classified_defs['simple'])
        self._generate_complex_type_files(classified_defs['complex'])
        
        # --- File Cleanup ---
        stale_files_found = 0
        for filename_on_disk in os.listdir(self.output_dir):
            if not filename_on_disk.endswith(".go"):
                continue

            # Do not delete hand-written helper files
            if filename_on_disk not in self.generated_files and filename_on_disk not in ["e1ap_manual_types.go", "e1ap_pdu.go"]: 
                file_to_delete = os.path.join(self.output_dir, filename_on_disk)
                os.remove(file_to_delete)
                logger.info(f"Cleaned up stale generated file: {filename_on_disk}")
                stale_files_found += 1
        
        if stale_files_found == 0:
            logger.info("No stale generated files to clean up.")

        # --- Final Formatting ---
        self._format_generated_code()

    def _format_generated_code(self):
        """Runs gofmt and goimports on the output directory."""
        logger.info("Formatting generated Go code...")
        try:
            # Using os.system for simplicity; subprocess is more robust for production
            os.system(f"goimports -w {self.output_dir}")
            os.system(f"gofmt -w {self.output_dir}")
            logger.info("Formatting complete.")
        except Exception as e:
            logger.warning(f"Could not format Go code. Please run 'goimports -w' and 'gofmt -w' on the '{self.output_dir}' directory manually. Error: {e}")

    def _generate_constants_file(self, constant_items: List[Any]):
        integer_consts, proc_code_consts, protocol_ie_consts = [], [], []
        proc_code_type_name = self._standard_string("ProcedureCode")
        protocol_ie_id_type_name = self._standard_string("ProtocolIE-ID")

        for item in self.definitions.values():
            if not item.is_constant:
                continue

            base_name = self._standard_string(item.name)
            if not base_name:
                continue

            value = self._format_go_value(item.min_val)

            if item.type == "ProcedureCode":
                proc_code_consts.append((f"ProcedureCode_{base_name}", value))
            elif item.type == "ProtocolIE-ID":
                protocol_ie_consts.append((f"ProtocolIEID_{base_name}", value))
            elif item.type == "INTEGER":
                integer_consts.append((base_name, value))

        sorter = (
            lambda item: int(item[1]) if item[1] and item[1].isdigit() else float("inf")
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
        """Generates a single file for all simple, reusable type definitions with validation."""
        if not simple_items:
            logger.info("No simple types found to generate in common file.")
            return

        # Start with package and necessary imports for validation errors
        go_code = "package e1ap_ies\n\nimport (\n\t\"fmt\"\n\t\"github.com/lvdund/ngap/aper\"\n)\n\n"
        
        from collections import defaultdict
        type_counts = defaultdict(int)
        total_count = 0
        
        sorted_items = sorted(simple_items, key=lambda x: self._standard_string(x.name))
        
        for item in sorted_items:
            go_name = self._standard_string(item.name)
            go_base_type, concrete_def = self._resolve_go_type(item.name)
            
            final_item = concrete_def if concrete_def else item
            
            is_redundant_builtin = item.is_builtin and go_name.lower() == go_base_type.lower()
            if go_name.lower() != go_base_type.lower() and not is_redundant_builtin:
                
                go_code += self._generate_header_comment(final_item, go_name)
                go_code += f"type {go_name} {go_base_type}\n\n"
                
                has_constraints = final_item.min_val is not None
                
                # Only generate validation for types where it makes sense.
                is_validatable_integer = (final_item.type == "INTEGER" and final_item.max_val != "18446744073709551615")
                is_validatable_octet_string = final_item.type == "OCTET STRING"

                if has_constraints and (is_validatable_integer or is_validatable_octet_string):
                    min_val_unformatted = final_item.min_val
                    max_val_unformatted = final_item.max_val if final_item.max_val is not None else final_item.min_val
                    
                    min_val_str = self._format_go_value(min_val_unformatted)
                    max_val_str = self._format_go_value(max_val_unformatted)

                    try:
                        int(min_val_unformatted)
                        min_val_assign = min_val_str
                    except (ValueError, TypeError):
                        min_val_assign = f"{go_name}({min_val_str})"

                    try:
                        int(max_val_unformatted)
                        max_val_assign = max_val_str
                    except (ValueError, TypeError):
                        max_val_assign = f"{go_name}({max_val_str})"

                    if min_val_str != max_val_str:
                        go_code += "const (\n"
                        go_code += f"\tMin{go_name} {go_name} = {min_val_assign}\n"
                        go_code += f"\tMax{go_name} {go_name} = {max_val_assign}\n"
                        go_code += ")\n\n"

                    go_code += f"// Validate checks if the value is within the specified range.\n"
                    go_code += f"func (v {go_name}) Validate() error {{\n"
                    
                    min_check = f"Min{go_name}" if min_val_str != max_val_str else min_val_assign
                    max_check = f"Max{go_name}" if min_val_str != max_val_str else max_val_assign

                    check_logic = ""
                    error_range_str = ""
                    
                    if final_item.type == "OCTET STRING":
                        check_logic = f"len(v) < int({min_check}) || len(v) > int({max_check})"
                        error_range_str = f"length ({min_check}..{max_check})"
                    else: # INTEGER
                        check_logic = f"v < {min_check} || v > {max_check}"
                        error_range_str = f"range ({min_check}..{max_check})"

                    go_code += f"\tif {check_logic} {{\n"
                    go_code += f'\t\treturn fmt.Errorf("value for {go_name} is outside the valid {error_range_str}")\n'
                    go_code += "\t}\n"
                    go_code += "\treturn nil\n"
                    go_code += "}\n\n"

                total_count += 1
                type_counts[final_item.type] += 1
        
        if total_count > 0:
            file_path = os.path.join(self.output_dir, "e1ap_common_types.go")
            with open(file_path, "w", encoding="utf-8") as f:
                f.write(go_code)
            self.generated_files.add("e1ap_common_types.go")

            for def_type, count in sorted(type_counts.items()):
                if count > 0:
                    plural = "s" if count > 1 else ""
                    logger.info(f"SUCCESS: Wrote {count} {def_type} type definition{plural} with validation to 'e1ap_common_types.go'.")
        else:
            logger.info("No simple type aliases needed to be generated.")

    def _generate_complex_type_files(self, complex_items: List[Any]):
        """Generates a dedicated file for each complex type using a robust two-pass strategy."""
        if not complex_items:
            logger.info("No complex types found to generate individual files for.")
            return

        type_counts = { "SEQUENCE": 0, "CHOICE": 0, "ENUMERATED": 0, "LIST": 0 }
        
        # --- PASS 1: Generate all non-list types ---
        logger.info("Complex types pass 1: Generating SEQUENCE, CHOICE, and ENUMERATED files...")
        for item in complex_items:
            if item.type in ["SEQUENCE", "CHOICE", "ENUMERATED"]:
                go_name = self._standard_string(item.name)
                filename = self._go_name_to_snake_case(go_name) + ".go"
                file_path = os.path.join(self.output_dir, filename)
                
                item_code = self._generate_header_comment(item, go_name)
                item_code += self._generate_code_for_item(item)
                
                file_content = "package e1ap_ies\n\n"
                
                imports = set()
                # ENUMERATED and CHOICE types always need the aper library.
                if item.type in ["ENUMERATED", "CHOICE"]:
                    imports.add('"github.com/lvdund/ngap/aper"')

                # PDU SEQUENCE types need all three imports for their Encode/Decode methods.
                is_pdu = item.type == "SEQUENCE" and item.ies and 'id' in item.ies[0]
                if is_pdu:
                    imports.add('"fmt"')
                    imports.add('"io"')
                    imports.add('"github.com/lvdund/ngap/aper"')

                if imports:
                    file_content += "import (\n"
                    for imp in sorted(list(imports)):
                        file_content += f"\t{imp}\n"
                    file_content += ")\n\n"
                
                file_content += item_code
                with open(file_path, "w", encoding="utf-8") as f:
                    f.write(file_content)
                self.generated_files.add(filename)
                
                if item.type in type_counts:
                    type_counts[item.type] += 1

        # --- PASS 2: Generate all LIST types ---
        # These files will now correctly reference the item files created in Pass 1.
        logger.info("Complex types pass 2: Generating LIST files...")
        for item in complex_items:
            if item.type == "LIST":
                if not item.of_type:
                    logger.warning(f"Skipping LIST {item.name} because its element type is unknown.")
                    continue

                go_name = self._standard_string(item.name)
                filename = self._go_name_to_snake_case(go_name) + ".go"
                file_path = os.path.join(self.output_dir, filename)

                of_type_go_name = self._standard_string(item.of_type)
                
                item_code = self._generate_header_comment(item, go_name)
                item_code += f"type {go_name} []{of_type_go_name}\n"

                file_content = f"package e1ap_ies\n\n{item_code}"
                with open(file_path, "w", encoding="utf-8") as f:
                    f.write(file_content)
                self.generated_files.add(filename)

                if item.type in type_counts:
                    type_counts[item.type] += 1
        
        # Final logging
        total_complex_count = sum(type_counts.values())
        if total_complex_count > 0:
            for def_type, count in type_counts.items():
                if count > 0:
                    logger.info(f"SUCCESS: Wrote {count} {def_type} definitions to individual files.")


    def _generate_code_for_item(self, item: Any) -> str:
        """Generates the Go code string for a single complex ASN.1 definition,
           including advanced Encode/Decode methods for PDUs."""
        go_name = self._standard_string(item.name)
        item_code = ""

        if item.type == "ENUMERATED":
            item_code += "const (\n"
            for i, val in enumerate(item.enum_values):
                enum_member_name = self._standard_string(val).replace("_", "")
                enum_name = f"{go_name}{enum_member_name}"
                item_code += f"\t{enum_name} aper.Enumerated = {i}\n"
            item_code += ")\n\n"
            item_code += f"type {go_name} struct {{\n"
            item_code += "\tValue aper.Enumerated\n"
            item_code += "}\n"

        elif item.type == "SEQUENCE":
            # Heuristic: A SEQUENCE is a PDU if its members were expanded from a manifest.
            # We check if the 'id' key exists in the FIRST member as a reliable indicator.
            is_pdu = item.ies and 'id' in item.ies[0]

            # 1. Generate the Struct Definition
            item_code += f"type {go_name} struct {{\n"
            for member in item.ies:
                member_go_name = self._standard_string(member['ie'])
                original_type_name = member['type']
                if not member_go_name or not original_type_name: continue
                
                base_go_type, concrete_def = self._resolve_go_type(original_type_name)
                
                final_go_type = base_go_type
                is_list = False
                if concrete_def and concrete_def.type == "LIST":
                    of_type_go_name = self._standard_string(concrete_def.of_type)
                    final_go_type = f"[]{of_type_go_name}"
                    is_list = True

                is_optional = member['presence'] in ["optional", "conditional"]
                pointer = "*" if is_optional and not is_list else ""
                
                tag_parts = []
                if concrete_def:
                    if concrete_def.min_val is not None:
                        tag_parts.append(f"lb:{self._format_go_value(concrete_def.min_val)}")
                    if concrete_def.max_val is not None:
                        tag_parts.append(f"ub:{self._format_go_value(concrete_def.max_val)}")
                
                tag_parts.append(member['presence'])
                if 'criticality' in member:
                    tag_parts.append(member['criticality'])

                if item.is_extensible:
                    tag_parts.append("ext")
                
                tag_string = f'`aper:"{",".join(tag_parts)}"`'
                item_code += f"\t{member_go_name}\t{pointer}{final_go_type}\t{tag_string}\n"
            item_code += "}\n\n"

            # 2. If it's a PDU, generate the advanced methods
            if is_pdu:
                # --- The toIes() method ---
                item_code += f"// toIes converts the PDU struct into a list of generic IEs.\n"
                item_code += f"func (s *{go_name}) toIes() ([]E1apMessageIE, error) {{\n"
                item_code += "\ties := []E1apMessageIE{}\n"
                item_code += "\tvar err error\n\n"

                for member in item.ies:
                    member_go_name = self._standard_string(member['ie'])
                    
                    # This code is now safe because we know 'id' exists for all members of a PDU.
                    ie_id_name = self._standard_string(member['id'].replace("id-",""))
                    crit_name = self._standard_string(member['criticality'])
                    _, concrete_def = self._resolve_go_type(member['type'])
                    
                    is_optional = member['presence'] == 'optional'

                    if is_optional:
                        item_code += f"\tif s.{member_go_name} != nil {{\n\t" # Indent for optional block
                    
                    item_code += "\ties = append(ies, E1apMessageIE{\n"
                    item_code += f"\t\tId:          ProtocolIEID{{Value: ProtocolIEID_{ie_id_name}}},\n"
                    item_code += f"\t\tCriticality: Criticality{{Value: Criticality_Present{crit_name}}},\n"
                    
                    value_code = "nil // TODO: Implement value encoding"
                    if concrete_def:
                        val_prefix = "s." if not is_optional else "(*s."
                        val_suffix = "" if not is_optional else ")"
                        
                        if concrete_def.type == "INTEGER":
                            value_code = f"&INTEGER{{Value: {val_prefix}{member_go_name}{val_suffix}}}"
                        elif concrete_def.type == "OCTET STRING":
                            value_code = f"&OCTETSTRING{{Value: {val_prefix}{member_go_name}{val_suffix}}}"
                        elif concrete_def.type == "BIT STRING":
                            value_code = f"&BITSTRING{{Value: {val_prefix}{member_go_name}{val_suffix}}}"
                        elif concrete_def.type in ["SEQUENCE", "CHOICE", "LIST", "ENUMERATED"]:
                            value_code = f"s.{member_go_name}"

                    item_code += f"\t\tValue: {value_code},\n"
                    item_code += "\t})\n"

                    if is_optional:
                        item_code += f"\t}}\n\n"

                item_code += "\treturn ies, err\n}\n\n"

                # --- The Encode() method ---
                message_category = self.message_to_procedure_map.get(item.name, "Unknown")
                pdu_type = f"E1apPdu{message_category}"
                
                proc_code = "0 // TODO: Find ProcedureCode"
                for proc in self.procedures.values():
                    if proc.initiating_message == item.name or proc.successful_outcome == item.name or proc.unsuccessful_outcome == item.name:
                        if proc.procedure_code:
                            proc_code = f"ProcedureCode_{self._standard_string(proc.procedure_code.replace('id-',''))}"
                        break
                
                item_code += f"// Encode implements the aper.AperMarshaller interface.\n"
                item_code += f"func (s *{go_name}) Encode(w io.Writer) error {{\n"
                item_code += "\ties, err := s.toIes()\n"
                item_code += "\tif err != nil {\n\t\treturn fmt.Errorf(\"Encode %s failed: %v\", \"{go_name}\", err)\n\t}\n"
                item_code += f"\treturn encodeMessage(w, {pdu_type}, {proc_code}, Criticality_PresentIgnore, ies)\n}}\n\n"
                
                # --- The Decode() method ---
                item_code += f"// Decode implements the aper.AperUnmarshaller interface.\n"
                item_code += f"func (s *{go_name}) Decode(r io.Reader) error {{\n"
                item_code += "\treturn fmt.Errorf(\"Decode not implemented for {go_name}\")\n}}\n\n"

        elif item.type == "CHOICE":
            item_code += "const (\n"
            item_code += f"\t{go_name}PresentNothing uint64 = iota\n"
            for member in item.ies:
                member_go_name = self._standard_string(member['ie'])
                item_code += f"\t{go_name}Present{member_go_name}\n"
            item_code += ")\n\n"
            item_code += f"type {go_name} struct {{\n"
            item_code += "\tChoice\tuint64\n"
            for member in item.ies:
                member_go_name = self._standard_string(member['ie'])
                base_go_type, concrete_def = self._resolve_go_type(member['type'])
                final_go_type = base_go_type
                is_list = False
                if concrete_def and concrete_def.type == "LIST":
                    of_type_go_name = self._standard_string(concrete_def.of_type)
                    final_go_type = f"[]{of_type_go_name}"
                    is_list = True
                pointer = "" if is_list else "*"
                item_code += f"\t{member_go_name}\t{pointer}{final_go_type}\n"
            item_code += "}\n"
            
        return item_code


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