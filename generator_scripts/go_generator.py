import os
import re
import logging
import shutil
from typing import Dict, List, Any

logger = logging.getLogger(__name__)

class GoCodeGenerator:
    def __init__(self, definitions: Dict[str, Any], failures: List[Dict[str, Any]], output_dir: str):
        self.definitions = definitions
        self.failures = failures
        self.output_dir = output_dir

    def _standard_string(self, input_string: str) -> str:
        """
        Final, robust version of the string standardizer.
        It correctly handles mixed-case words and all-caps acronyms to produce
        idiomatic Go names in PascalCase.
        """
        if not input_string:
            return ""

        # Special overrides for the most important base types with tricky casing.
        # This is the safest and most guaranteed way to get the exact output you want.
        overrides = {
            "ProcedureCode": "ProcedureCode",
            "ProtocolIE-ID": "ProtocolIeId",
            "ProtocolExtensionID": "ProtocolExtensionId",
            "ProtocolIE-ID": "ProtocolIeId" # Redundant but safe
        }
        if input_string in overrides:
            return overrides[input_string]

        # General logic for all other names
        s = re.sub(r"^(id-)", "", input_string)
        
        # Split by delimiters. This correctly separates words and acronyms.
        parts = re.split(r"[-_\s]", s)
        
        processed_parts = []
        for part in parts:
            if not part: continue
            
            # This is the key logic:
            # If a part is all uppercase (like "ID", "GNB"), it's an acronym.
            # We use .capitalize() which turns it into "Id", "Gnb" - idiomatic for Go.
            if part.isupper():
                processed_parts.append(part.capitalize())
            else:
                # If it's mixed case (like "ProcedureCode"), we must NOT use .capitalize().
                # We just ensure the first letter is capitalized and leave the rest alone.
                processed_parts.append(part[0].upper() + part[1:])

        return "".join(processed_parts)
    

    def _format_go_value(self, value: str) -> str:
        """
        Sanitizes a parsed value string and formats it for Go code.
        This version includes a MANDATORY DEBUG PRINT to prove it is running.
        """
        if not value:
            sanitized_value = "0"
        else:
            try:
                # Convert to integer to handle different bases (like '0255')
                # and then convert back to a base-10 string.
                sanitized_value = str(int(value))
            except (ValueError, TypeError):
                # It's a text constant like 'maxProtocolExtensions'
                sanitized_value = self._standard_string(value)
        return sanitized_value

    def generate_files(self):
        """Generates all the Go output files."""
        if os.path.exists(self.output_dir):
            shutil.rmtree(self.output_dir)
        os.makedirs(self.output_dir)
        
        self._generate_constants_file()
        self._generate_integer_types_file()
        self._generate_enumerated_types_file()

    def _generate_constants_file(self):
        """
        Generates the e1ap_constants.go file, ensuring all values are sanitized.
        """
        go_code, count = "", 0
        int_consts = sorted([item for item in self.definitions.values() if item.is_constant and item.type == "INTEGER"], key=lambda x: x.name)
        proc_consts = sorted([item for item in self.definitions.values() if item.is_constant and item.type == "ProcedureCode"], key=lambda x: x.name)
        ie_consts = sorted([item for item in self.definitions.values() if item.is_constant and item.type == "ProtocolIE-ID"], key=lambda x: x.name)
        
        if int_consts:
            max_len = max((len(self._standard_string(item.name)) for item in int_consts), default=0)
            go_code += "const (\n"
            for item in int_consts:
                name = self._standard_string(item.name)
                value = self._format_go_value(item.min_val)
                go_code += f"\t{name.ljust(max_len)} int64 = {value}\n"
            go_code += ")\n\n"
        
        if proc_consts:
            max_len = max((len(f"ProcedureCode_{self._standard_string(item.name)}") for item in proc_consts), default=0)
            go_code += "const (\n"
            for item in proc_consts:
                name = f"ProcedureCode_{self._standard_string(item.name)}"
                value = self._format_go_value(item.min_val)
                go_code += f"\t{name.ljust(max_len)} ProcedureCode = {value}\n"
            go_code += ")\n\n"

        if ie_consts:
            max_len = max((len(f"ProtocolIEID_{self._standard_string(item.name)}") for item in ie_consts), default=0)
            go_code += "const (\n"
            for item in ie_consts:
                name = f"ProtocolIEID_{self._standard_string(item.name)}"
                value = self._format_go_value(item.min_val)
                go_code += f"\t{name.ljust(max_len)} ProtocolIeId = {value}\n"
            go_code += ")\n\n"
            
        count = len(int_consts) + len(proc_consts) + len(ie_consts)
        if go_code:
            file_path = os.path.join(self.output_dir, "e1ap_constants.go")
            with open(file_path, 'w', encoding='utf-8') as f:
                f.write(f"package e1ap_ies\n\n{go_code}")
            logger.info(f"SUCCESS: Wrote {count} constants to 'e1ap_constants.go'.")

    def _generate_integer_types_file(self):
        """
        Generates the integer_types.go file. This FINAL version uses a unified
        loop to correctly generate 'type' and 'const' blocks for both regular
        INTEGERs and the special BASE_TYPEs.
        """
        go_code, count = "", 0
        MAX_INT64 = 9223372036854775807

        # --- 1. Get all integer-like definitions ---
        base_types = [item for item in self.definitions.values() if item.type == "BASE_TYPE"]
        integers = [item for item in self.definitions.values() if item.type == "INTEGER" and not item.is_constant]
        
        # --- 2. Combine them into a single list for unified processing ---
        all_integer_defs = sorted(base_types + integers, key=lambda x: x.name)

        # --- 3. Process every definition in the unified loop ---
        for item in all_integer_defs:
            go_name = self._standard_string(item.name)
            go_type = "int64" # Default to int64

            try:
                if item.max_val and int(item.max_val) > MAX_INT64:
                    go_type = "uint64"
            except (ValueError, TypeError):
                pass
            
            # --- Step A: Generate the 'type' declaration for EVERY item ---
            go_code += f"type {go_name} {go_type}\n"
            
            # --- Step B: Generate the 'const' block for ANY item that has constraints ---
            if item.min_val is not None:
                go_code += "const (\n"
                # Use the safe sanitizer that is already proven to work
                min_val_str = self._format_go_value(item.min_val)
                max_val_str = self._format_go_value(item.max_val)
                
                go_code += f"\t{go_name}MinValue {go_type} = {min_val_str}\n"
                go_code += f"\t{go_name}MaxValue {go_type} = {max_val_str}\n"
                go_code += ")\n"
            go_code += "\n" # Add a space after each full definition
            count += 1
            
        # --- 4. Write the final, correct file ---
        if go_code:
            file_path = os.path.join(self.output_dir, "integer_types.go")
            with open(file_path, 'w', encoding='utf-8') as f:
                f.write(f"package e1ap_ies\n\n{go_code}")
            logger.info(f"SUCCESS: Wrote {count} INTEGER type definitions to 'integer_types.go'.")

    def _generate_enumerated_types_file(self):
        """
        Generates the enumerated_types.go file.
        """
        go_code, count = "", 0
        enums = sorted([item for item in self.definitions.values() if item.type == "ENUMERATED"], key=lambda x: x.name)

        for item in enums:
            go_name = self._standard_string(item.name)
            go_code += f"type {go_name} int32\n"
            
            if item.enum_values:
                go_code += "const (\n"
                for i, val in enumerate(item.enum_values):
                    enum_name = self._standard_string(val)
                    go_code += f"\t{go_name}_{enum_name} {go_name} = {i}\n"
                go_code += ")\n"
            go_code += "\n"
            count += 1
            
        if go_code:
            file_path = os.path.join(self.output_dir, "enumerated_types.go")
            with open(file_path, 'w', encoding='utf-8') as f:
                f.write(f"package e1ap_ies\n\n{go_code}")
            logger.info(f"SUCCESS: Wrote {count} ENUMERATED type definitions to 'enumerated_types.go'.")

    def run_full_diagnostic(self):
        """
        Prints a comprehensive and structured report of parsing successes and failures.
        This function is the primary tool for verifying the parser's correctness.
        """
        print("\n" + "="*80)
        print(" " * 25 + "ASN.1 PARSER DIAGNOSTIC REPORT")
        print("="*80)

        
        successes = {
            "CONSTANT": [], "INTEGER": [], "ENUMERATED": [], "SEQUENCE": [],
            "CHOICE": [], "LIST": [], "BIT STRING": [], "OCTET STRING": [],
            "PrintableString": [], "VisibleString": [], "UTF8String": [], "NULL": []
        }
        
        for item in self.definitions.values():
            if item.is_constant:
                successes["CONSTANT"].append(item)
            elif item.type in successes:
                successes[item.type].append(item)

        
        total_success_count = 0
        
        report_format = "  {context:<50} {name}"

        
        for def_type in ["CONSTANT", "INTEGER", "ENUMERATED", "BIT STRING", "OCTET STRING", "PrintableString", "VisibleString", "UTF8String", "NULL", "SEQUENCE", "CHOICE", "LIST"]:
            items = successes.get(def_type)
            if items:
                sorted_items = sorted(items, key=lambda x: x.name)
                print(f"\n--- [ SUCCESS ] PARSED {def_type} DEFINITIONS ({len(sorted_items)}) ---")
                for item in sorted_items:
                    context = f"({item.source_file}:{item.source_line})"
                    print(report_format.format(context=context, name=item.name))
                total_success_count += len(sorted_items)
        
        print("\n" + "-"*40)
        print(f"TOTAL SUCCESSFUL DEFINITIONS: {total_success_count} (out of {len(self.definitions)} total)")
        print("-" * 40)

        
        print(f"\n--- [ FAILURE ] SKIPPED ABSTRACT/UNHANDLED DEFINITIONS ({len(self.failures)}) ---")
        sorted_failures = sorted(self.failures, key=lambda x: x['name'])
        for failure in sorted_failures:
            context = f"({failure['file']}:{failure['line']})"
            print(report_format.format(context=context, name=failure['name']))
        
        print("\n" + "="*80)
        print("ACTION REQUIRED: Verify the [ SUCCESS ] counts against your manual analysis.")
        print("The [ FAILURE ] list should only contain abstract schemas like E1AP-PROTOCOL-IES.")
        print("="*80)