import os
import re
import logging
import shutil
from typing import Dict, List, Tuple
from asn1_parser import ASN1Definition # Import the data structure


logger = logging.getLogger(__name__)

class GoCodeGenerator:
    def __init__(self, definitions: Dict[str, ASN1Definition], output_dir: str):
        self.definitions = definitions
        self.output_dir = output_dir

        # --- NEW: Cleanup and Recreate the Output Directory ---
        # If the directory already exists from a previous run, delete it completely.
        if os.path.exists(self.output_dir):
            logger.info(f"Cleaning up old output directory: '{self.output_dir}'")
            try:
                shutil.rmtree(self.output_dir)
            except OSError as e:
                logger.error(f"Error removing directory {self.output_dir}: {e}", exc_info=True)
                # If we can't clean the directory, it's safer to stop.
                raise

        # Create a fresh, empty directory for the new output.
        logger.info(f"Creating fresh output directory: '{self.output_dir}'")
        os.makedirs(self.output_dir)

    def _standard_string(self, input_string: str) -> str:
        """Utility to convert ASN.1 names to Go-style names."""
        if not input_string: return ""
        # More robustly handle complex names
        input_string = re.sub(r"^(id-)", "", input_string)
        return "".join([word.capitalize() for word in re.split(r"[-_.&{}\s:]", input_string.replace('INTEGER','').strip()) if word])

    def generate_files(self):
        """Generates all the Go output files."""
        self._generate_constants_file()
        self._generate_integer_types_file()
        self._generate_bitstring_types_file()
        self._generate_octetstring_types_file()
        self._generate_sequence_types_file()
        self._generate_choice_types_file()
        self._generate_enumerated_types_file()

    def _generate_constants_file(self):
        """Generates the e1ap_constants.go file."""
        integer_consts, proc_code_consts, protocol_ie_consts = [], [], []

        for item in self.definitions.values():
            if not item.is_constant:
                continue
            
            base_name = self._standard_string(item.name)
            if not base_name: continue

            if item.type == "ProcedureCode":
                proc_code_consts.append((f"ProcedureCode_{base_name}", item.min_val))
            elif item.type == "ProtocolIE-ID":
                protocol_ie_consts.append((f"ProtocolIEID_{base_name}", item.min_val))
            elif item.type == "INTEGER":
                integer_consts.append((base_name, item.min_val))
        
        # Sorter function
        sorter = lambda item: int(item[1]) if item[1] and item[1].isdigit() else float('inf')
        integer_consts.sort(key=sorter)
        proc_code_consts.sort(key=sorter)
        protocol_ie_consts.sort(key=sorter)

        # Build Go file content
        go_code = "// This file contains constant values extracted from the ASN.1 specification.\n\n"
        if integer_consts:
            go_code += "// General INTEGER Constants\nconst (\n" + self._format_const_block(integer_consts, typed=True) + ")\n\n"
        if proc_code_consts:
            go_code += "// Procedure Codes\nconst (\n" + self._format_const_block(proc_code_consts) + ")\n\n"
        if protocol_ie_consts:
            go_code += "// Protocol IE IDs\nconst (\n" + self._format_const_block(protocol_ie_consts) + ")\n\n"

        total_consts = len(integer_consts) + len(proc_code_consts) + len(protocol_ie_consts)
        if total_consts > 0:
            file_path = os.path.join(self.output_dir, "e1ap_constants.go")
            with open(file_path, 'w', encoding='utf-8') as f:
                f.write(f"package e1ap_ies\n\n{go_code}")
            logger.info(f"SUCCESS: Wrote {total_consts} constants to 'e1ap_constants.go'.")
        else:
            logger.info("No constants found to generate.")


    def _format_const_block(self, constants: List[Tuple[str, str]], typed=False) -> str:
        if not constants: return ""
        max_len = max(len(name) for name, _ in constants)
        lines = []
        for name, value in constants:
            padding = " " * (max_len - len(name))
            type_str = " int64" if typed else ""
            lines.append(f"\t{name}{padding}{type_str} = {value}")
        return "\n".join(lines) + "\n"

    def _generate_integer_types_file(self):
        """Generates the integer_types.go file."""
        go_code = ""
        count = 0
        
        sorted_items = sorted(self.definitions.items())

        for name, item in sorted_items:
            if item.type == "INTEGER" and not item.is_constant:
                go_name = self._standard_string(name)
                if not go_name: continue
                logger.debug(f"Generating INTEGER type: {go_name}")
                
                go_code += f"// {name} ::= INTEGER\n"
                go_code += f"type {go_name} int64\n"
                
                const_lines = []
                if item.min_val is not None:
                    min_val_str = self._standard_string(item.min_val) if not item.min_val.isdigit() else item.min_val
                    const_lines.append(f"\t{go_name}MinValue = {min_val_str}")
                if item.max_val is not None:
                    max_val_str = self._standard_string(item.max_val) if not item.max_val.isdigit() else item.max_val
                    const_lines.append(f"\t{go_name}MaxValue = {max_val_str}")

                if const_lines:
                    go_code += "const (\n" + "\n".join(const_lines) + "\n)\n"
                go_code += "\n"
                count += 1

        if go_code:
            file_path = os.path.join(self.output_dir, "integer_types.go")
            with open(file_path, 'w', encoding='utf-8') as f:
                f.write("package e1ap_ies\n\n" + go_code)
            logger.info(f"SUCCESS: Wrote {count} INTEGER type definitions to 'integer_types.go'.")
        else:
            logger.info("No INTEGER type aliases were found to generate.")



    def _generate_bitstring_types_file(self):
        """Generates the bitstring_types.go file."""
        go_code = ""
        count = 0
        
        sorted_items = sorted(self.definitions.items())

        for name, item in sorted_items:
            if item.type == "BIT STRING":
                go_name = self._standard_string(name)
                if not go_name: continue
                logger.debug(f"Generating BIT STRING type: {go_name}")
                
                # A common and precise Go representation for a BIT STRING uses a struct
                # to hold both the byte slice and the exact number of bits.
                go_code += f"// {name} ::= BIT STRING\n"
                go_code += f"type {go_name} struct {{\n"
                go_code += f"\tValue []byte `asn1:\"value\"` // The byte array representing the bit string\n"
                go_code += f"\tLen   uint   `asn1:\"len\"`   // The number of bits in the bit string\n"
                go_code += f"}}\n"

                const_lines = []
                # min_val and max_val from the parser refer to the number of BITS
                if item.min_val is not None:
                    min_val_str = self._standard_string(item.min_val) if not item.min_val.isdigit() else item.min_val
                    const_lines.append(f"\t{go_name}MinBits = {min_val_str}")
                if item.max_val is not None:
                    max_val_str = self._standard_string(item.max_val) if not item.max_val.isdigit() else item.max_val
                    const_lines.append(f"\t{go_name}MaxBits = {max_val_str}")

                if const_lines:
                    go_code += "const (\n" + "\n".join(const_lines) + "\n)\n"
                go_code += "\n"
                count += 1

        if go_code:
            file_path = os.path.join(self.output_dir, "bitstring_types.go")
            with open(file_path, 'w', encoding='utf-8') as f:
                f.write("package e1ap_ies\n\n" + go_code)
            logger.info(f"SUCCESS: Wrote {count} BIT STRING type definitions to 'bitstring_types.go'.")
        else:
            logger.info("No BIT STRING type aliases were found to generate.")


    def _generate_octetstring_types_file(self):
        """Generates the octetstring_types.go file for OCTET STRING and other string types."""
        go_code = ""
        count = 0
        
        # Define which ASN.1 types we want to handle in this file
        string_types_to_generate = {"OCTET STRING", "PrintableString", "VisibleString", "UTF8String"}
        
        sorted_items = sorted(self.definitions.items())

        for name, item in sorted_items:
            if item.type in string_types_to_generate:
                go_name = self._standard_string(name)
                if not go_name: continue

                # OCTET STRING is a byte slice, others are native Go strings.
                go_type = "[]byte" if item.type == "OCTET STRING" else "string"
                
                logger.debug(f"Generating {item.type} type: {go_name}")
                
                go_code += f"// {name} ::= {item.type}\n"
                go_code += f"type {go_name} {go_type}\n\n"
                count += 1

        if go_code:
            file_path = os.path.join(self.output_dir, "octetstring_types.go")
            with open(file_path, 'w', encoding='utf-8') as f:
                f.write("package e1ap_ies\n\n" + go_code)
            logger.info(f"SUCCESS: Wrote {count} OCTET STRING/String type definitions to 'octetstring_types.go'.")

    def _map_asn1_to_go_type(self, asn1_type_name: str, presence: str) -> str:
        """
        Maps an ASN.1 type name to a suitable Go type name.

        This helper handles the conversion of ASN.1 standard names to Go-style
        names, maps fundamental types to their Go equivalents (e.g., INTEGER to int64),
        and crucially, converts optional ASN.1 members into pointer types in Go
        to allow for 'nil' values.

        Args:
            asn1_type_name: The ASN.1 type of the member (e.g., "GNB-CU-UP-ID").
            presence: The presence of the member ("optional", "mandatory", etc.).

        Returns:
            The corresponding Go type as a string (e.g., "*GnbCuUpId").
        """
        # Start with the standardized Go name (e.g., "GNB-CU-UP-ID" -> "GnbCuUpId")
        go_type = self._standard_string(asn1_type_name)

        # Handle overrides for fundamental types that have direct Go equivalents
        if go_type == "INTEGER": go_type = "int64"
        elif go_type == "BOOLEAN": go_type = "bool"
        elif go_type == "OCTETSTRING": go_type = "[]byte"
        elif go_type == "NULL": go_type = "[]byte" # ASN.1 NULL is often a nil byte slice in Go libraries

        # Optional and Conditional members must be pointers in Go to be nillable.
        if presence.lower() in ["optional", "conditional"]:
            # Slices and maps are already reference types in Go, so they don't need to be pointers.
            if go_type.startswith("[]"):
                return go_type
            return f"*{go_type}"
        
        return go_type

    def _generate_sequence_types_file(self):
        """Generates the sequence_types.go file."""
        go_code, count = "", 0
        sorted_items = sorted(self.definitions.items())
        for name, item in sorted_items:
            if item.type == "SEQUENCE":
                go_name = self._standard_string(name)
                if not go_name: continue
                logger.debug(f"Generating SEQUENCE type: {go_name}")
                go_code += f"// {go_name} represents the ASN.1 SEQUENCE type.\ntype {go_name} struct {{\n"
                for member in item.ies:
                    member_name_go = self._standard_string(member['ie'])
                    member_type_go = self._map_asn1_to_go_type(member['type'], member['presence'])
                    tag = f'`json:"{member["ie"]},omitempty"`'
                    go_code += f"\t{member_name_go} {member_type_go} {tag}\n"
                go_code += "}\n\n"
                count += 1
        if go_code:
            file_path = os.path.join(self.output_dir, "sequence_types.go")
            with open(file_path, 'w', encoding='utf-8') as f: f.write("package e1ap_ies\n\n" + go_code)
            logger.info(f"SUCCESS: Wrote {count} SEQUENCE type definitions to 'sequence_types.go'.")

    def _generate_choice_types_file(self):
        """Generates the choice_types.go file."""
        go_code, count = "", 0
        sorted_items = sorted(self.definitions.items())
        for name, item in sorted_items:
            if item.type == "CHOICE":
                go_name = self._standard_string(name)
                if not go_name: continue
                logger.debug(f"Generating CHOICE type: {go_name}")
                go_code += f"// {go_name} represents the ASN.1 CHOICE type.\ntype {go_name} struct {{\n"
                for member in item.ies:
                    member_name_go = self._standard_string(member['ie'])
                    member_type_go = self._map_asn1_to_go_type(member['type'], member['presence'])
                    tag = f'`json:"{member["ie"]},omitempty"`'
                    go_code += f"\t{member_name_go} {member_type_go} {tag}\n"
                go_code += "}\n\n"
                count += 1
        if go_code:
            file_path = os.path.join(self.output_dir, "choice_types.go")
            with open(file_path, 'w', encoding='utf-8') as f: f.write("package e1ap_ies\n\n" + go_code)
            logger.info(f"SUCCESS: Wrote {count} CHOICE type definitions to 'choice_types.go'.")

    def _generate_enumerated_types_file(self):
        """Generates the enumerated_types.go file."""
        go_code, count = "", 0
        sorted_items = sorted(self.definitions.items())
        for name, item in sorted_items:
            if item.type == "ENUMERATED":
                go_name = self._standard_string(name)
                if not go_name: continue
                logger.debug(f"Generating ENUMERATED type: {go_name}")
                go_code += f"// {go_name} represents the ASN.1 ENUMERATED type.\ntype {go_name} int32\n\n"
                if item.enum_values:
                    go_code += f"// {go_name} Values\nconst (\n"
                    for i, val in enumerate(item.enum_values):
                        enum_val_name = self._standard_string(val)
                        go_code += f"\t{go_name}_{enum_val_name} {go_name} = {i}\n"
                    go_code += ")\n"
                go_code += "\n"
                count += 1
        if go_code:
            file_path = os.path.join(self.output_dir, "enumerated_types.go")
            with open(file_path, 'w', encoding='utf-8') as f: f.write("package e1ap_ies\n\n" + go_code)
            logger.info(f"SUCCESS: Wrote {count} ENUMERATED type definitions to 'enumerated_types.go'.")