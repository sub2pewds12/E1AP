import os
import re
import logging
import shutil
from typing import Dict, List, Any, Tuple, Optional

logger = logging.getLogger(__name__)


class GoCodeGenerator:
    def __init__(
        self,
        definitions: Dict[str, Any],
        failures: List[Dict[str, Any]],
        output_dir: str,
    ):
        self.definitions = definitions
        self.failures = failures
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
            sanitized_value = "0"
        else:
            try:

                sanitized_value = str(int(value))
            except (ValueError, TypeError):

                sanitized_value = self._standard_string(value)
        return sanitized_value
    
    def _go_name_to_snake_case(self, go_name: str) -> str:
        """Converts a PascalCase string to snake_case for filenames."""
        s1 = re.sub('(.)([A-Z][a-z]+)', r'\1_\2', go_name)
        return re.sub('([a-z0-9])([A-Z])', r'\1_\2', s1).lower()
    
    def _classify_definition(self, item: Any) -> str:
        """Classifies a definition as 'constant', 'simple', or 'complex'."""
        if item.is_constant:
            return 'constant'
        
        
        if item.type in ["SEQUENCE", "CHOICE", "LIST", "ENUMERATED"]:
            return 'complex'
            
        
        if item.type in ["INTEGER", "BIT STRING", "OCTET STRING", "ALIAS", "BASE_TYPE", "PrintableString", "VisibleString", "UTF8String", "NULL"]:
            return 'simple'
            
        return 'unknown'
    
    def _resolve_go_type(self, type_name: str) -> Tuple[str, Optional[Any]]:
        """
        Resolves an ASN.1 type name to its final, concrete Go type and returns
        the concrete ASN1Definition object.

        This function traverses the alias chain (e.g., TypeC ::= TypeB, TypeB ::= TypeA,
        TypeA ::= INTEGER) to find the underlying base type.

        Args:
            type_name: The starting ASN.1 type name string.

        Returns:
            A tuple containing:
            1. The corresponding base Go type as a string (e.g., "int64", "[]byte",
               or the standardized name of a complex type).
            2. The final, concrete ASN1Definition object, or None if not found.
        """
        definition = self.definitions.get(type_name)

        
        
        if not definition:
            return self._standard_string(type_name), None

        
        
        visited = {type_name}
        while definition and definition.alias_of:
            next_alias = definition.alias_of
            if next_alias in visited:
                logger.warning(f"Circular alias dependency detected for '{type_name}' at '{next_alias}'.")
                
                break
            
            visited.add(next_alias)
            next_def = self.definitions.get(next_alias)

            if not next_def:
                logger.warning(f"Alias '{next_alias}' for type '{definition.name}' not found in definitions.")
                
                break
            
            definition = next_def

        
        if definition.type in ["INTEGER", "BASE_TYPE"]:
            return "int64", definition
        elif definition.type in ["BIT STRING", "OCTET STRING", "PrintableString", "VisibleString", "UTF8String"]:
            return "[]byte", definition
        elif definition.type == "OBJECT IDENTIFIER": 
            return "string", definition
        elif definition.type == "NULL":
            return "bool", definition 
        elif definition.type in ["ENUMERATED", "SEQUENCE", "CHOICE", "LIST"]:
            
            return self._standard_string(definition.name), definition
        
        
        logger.warning(f"Could not resolve a specific Go type for ASN.1 type '{definition.type}'. Defaulting to its name.")
        return self._standard_string(definition.name), definition
    
    def generate_files(self):
        
        os.makedirs(self.output_dir, exist_ok=True)

        
        
        

        
        classified_defs = { 'constant': [], 'simple': [], 'complex': [] }
        
        for item in self.definitions.values():
            classification = self._classify_definition(item)
            if classification != 'unknown':
                classified_defs[classification].append(item)

        
        
        self._generate_constants_file(classified_defs['constant'])
        self._generate_common_types_file(classified_defs['simple'])
        self._generate_complex_type_files(classified_defs['complex'])
        
        
        

        stale_files_found = 0
        for filename_on_disk in os.listdir(self.output_dir):
            
            if not filename_on_disk.endswith(".go"):
                continue

            if filename_on_disk not in self.generated_files:
                
                
                
                if filename_on_disk not in ["e1ap_manual_types.go"]: 
                    file_to_delete = os.path.join(self.output_dir, filename_on_disk)
                    os.remove(file_to_delete)
                    logger.info(f"Cleaned up stale generated file: {filename_on_disk}")
                    stale_files_found += 1
        
        if stale_files_found == 0:
            logger.info("No stale generated files to clean up.")

        
        self._format_generated_code()

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

        go_code = "// This file contains constant values extracted from the ASN.1 specification.\n\n"

        if integer_consts:
            go_code += "// General INTEGER Constants\nconst (\n"
            max_len = max(len(name) for name, _ in integer_consts)
            for name, value in integer_consts:
                padding = " " * (max_len - len(name))
                go_code += f"\t{name}{padding} int64 = {value}\n"
            go_code += ")\n\n"

        if proc_code_consts:
            go_code += "// Procedure Codes\nconst (\n"
            max_len = max(len(name) for name, _ in proc_code_consts)
            for name, value in proc_code_consts:
                padding = " " * (max_len - len(name))
                
                go_code += f"\t{name}{padding} {proc_code_type_name} = {value}\n"
            go_code += ")\n\n"

        if protocol_ie_consts:
            go_code += "// Protocol IE IDs\nconst (\n"
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
        """Generates a single file for all simple, reusable type definitions."""
        if not simple_items:
            logger.info("No simple types found to generate in common file.")
            return

        go_code = "package e1ap_ies\n\n"
        count = 0

        sorted_items = sorted(simple_items, key=lambda x: self._standard_string(x.name))
        
        for item in sorted_items:
            go_name = self._standard_string(item.name)
            
            
            
            go_base_type, _ = self._resolve_go_type(item.name)
            
            
            
            if go_name.lower() != go_base_type.lower():
                go_code += f"type {go_name} {go_base_type} // From: {item.source_file}:{item.source_line}\n\n"
                count += 1
            

        file_path = os.path.join(self.output_dir, "e1ap_common_types.go")
        with open(file_path, "w", encoding="utf-8") as f:
            f.write(go_code)
        self.generated_files.add("e1ap_common_types.go")
        
        if count > 0:
            logger.info(
                f"SUCCESS: Wrote {count} simple type definitions to 'e1ap_common_types.go'."
            )

    def _generate_complex_type_files(self, complex_items: List[Any]):
        """Generates a dedicated file for each complex or meaningful type."""
        if not complex_items:
            logger.info("No complex types found to generate individual files for.")
            return

        count = 0
        for item in complex_items:
            go_name = self._standard_string(item.name)
            filename = self._go_name_to_snake_case(go_name) + ".go"
            file_path = os.path.join(self.output_dir, filename)

            item_code = ""

            if item.type == "ENUMERATED":
                
                item_code += f"type {go_name} int32\n\n"
                item_code += "const (\n"
                for i, val in enumerate(item.enum_values):
                    enum_name = self._standard_string(val)
                    item_code += f"\t{go_name}_{enum_name} {go_name} = {i}\n"
                item_code += ")\n"
            
            elif item.type in ["SEQUENCE", "CHOICE"]:
                if item.type == "CHOICE":
                    item_code += f"// {go_name} represents a CHOICE type.\n"
                
                item_code += f"type {go_name} struct {{\n"
                for member in item.ies:
                    member_go_name = self._standard_string(member['ie'])
                    original_type_name = member['type']

                    if not member_go_name or not original_type_name:
                        continue

                    

                    
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
                    if item.is_extensible:
                        tag_parts.append("ext")
                    
                    tag_string = f'`asn1:"{",".join(tag_parts)}"`'

                    item_code += f"\t{member_go_name}\t{pointer}{final_go_type}\t{tag_string}\n"
                    
                item_code += "}\n"

            
            elif item.type == "LIST":
                if not item.of_type:
                    logger.warning(f"Skipping LIST {item.name} because its element type is unknown.")
                    continue
                
                of_type_go_name = self._standard_string(item.of_type)
                item_code += f"// {go_name} is a list of {of_type_go_name}\n"
                item_code += f"type {go_name} []{of_type_go_name}\n"
            

            if item_code:
                
                file_content = "package e1ap_ies\n\n"
                file_content += f"// {go_name} represents the ASN.1 definition from {item.source_file}:{item.source_line}\n"
                file_content += item_code
                with open(file_path, "w", encoding="utf-8") as f:
                    f.write(file_content)
                self.generated_files.add(filename)
                count += 1
        
        if count > 0:
            logger.info(
                f"SUCCESS: Wrote {count} complex type definitions to individual files."
            )

    def _format_generated_code(self):
        """Runs gofmt and goimports on the output directory."""
        logger.info("Formatting generated Go code...")
        try:
            os.system(f"goimports -w {self.output_dir}")
            os.system(f"gofmt -w {self.output_dir}")
            logger.info("Formatting complete.")
        except Exception as e:
            logger.warning(f"Could not format Go code. Please run 'goimports -w' and 'gofmt -w' on the '{self.output_dir}' directory manually. Error: {e}")

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