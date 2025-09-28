import os
import re
import logging
import shutil
from typing import Dict, List, Any, Tuple

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

    def _standard_string(self, input_string: str) -> str:
        if not input_string:
            return ""

        overrides = {
            "ProcedureCode": "ProcedureCode",
            "ProtocolIE-ID": "ProtocolIeId",
            "ProtocolExtensionID": "ProtocolExtensionId",
            "ProtocolIE-ID": "ProtocolIeId",
        }
        if input_string in overrides:
            return overrides[input_string]

        s = re.sub(r"^(id-)", "", input_string)

        parts = re.split(r"[-_\s]", s)

        processed_parts = []
        for part in parts:
            if not part:
                continue

            if part.isupper():
                processed_parts.append(part.capitalize())
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

    def generate_files(self):
        if os.path.exists(self.output_dir):
            shutil.rmtree(self.output_dir)
        os.makedirs(self.output_dir)

        self._generate_constants_file()
        self._generate_integer_types_file()
        self._generate_enumerated_types_file()
        self._generate_bit_string_types_file()

    def _generate_constants_file(self):
        integer_consts, proc_code_consts, protocol_ie_consts = [], [], []

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

                go_code += f"\t{name}{padding} ProcedureCode = {value}\n"
            go_code += ")\n\n"

        if protocol_ie_consts:
            go_code += "// Protocol IE IDs\nconst (\n"
            max_len = max(len(name) for name, _ in protocol_ie_consts)
            for name, value in protocol_ie_consts:
                padding = " " * (max_len - len(name))

                go_code += f"\t{name}{padding} ProtocolIeId = {value}\n"
            go_code += ")\n\n"

        total_consts = (
            len(integer_consts) + len(proc_code_consts) + len(protocol_ie_consts)
        )
        if total_consts > 0:
            file_path = os.path.join(self.output_dir, "e1ap_constants.go")
            with open(file_path, "w", encoding="utf-8") as f:
                f.write(f"package e1ap_ies\n\n{go_code}")
            logger.info(
                f"SUCCESS: Wrote {total_consts} constants to 'e1ap_constants.go'."
            )
        else:
            logger.info("No constants found to generate.")

    def _generate_integer_types_file(self):
        go_code, count = "", 0
        MAX_INT64 = 9223372036854775807

        base_types = [
            item for item in self.definitions.values() if item.type == "BASE_TYPE"
        ]
        integers = [
            item
            for item in self.definitions.values()
            if item.type == "INTEGER" and not item.is_constant
        ]

        all_integer_defs = sorted(
            base_types + integers, key=lambda x: self._standard_string(x.name)
        )

        for item in all_integer_defs:
            go_name = self._standard_string(item.name)
            go_type = "int64"

            try:
                if item.max_val and int(item.max_val) > MAX_INT64:
                    go_type = "uint64"
            except (ValueError, TypeError):
                pass

            go_code += f"type {go_name} {go_type}\n"

            if item.min_val is not None:
                go_code += "const (\n"

                min_val_str = self._format_go_value(item.min_val)
                max_val_str = self._format_go_value(item.max_val)

                go_code += f"\t{go_name}MinValue {go_type} = {min_val_str}\n"
                go_code += f"\t{go_name}MaxValue {go_type} = {max_val_str}\n"
                go_code += ")\n"
            go_code += "\n"
            count += 1

        if go_code:
            file_path = os.path.join(self.output_dir, "integer_types.go")
            with open(file_path, "w", encoding="utf-8") as f:
                f.write(f"package e1ap_ies\n\n{go_code}")
            logger.info(
                f"SUCCESS: Wrote {count} INTEGER type definitions to 'integer_types.go'."
            )

    def _generate_enumerated_types_file(self):
        go_code, count = "", 0
        enums = sorted(
            [item for item in self.definitions.values() if item.type == "ENUMERATED"],
            key=lambda x: x.name,
        )

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
            with open(file_path, "w", encoding="utf-8") as f:
                f.write(f"package e1ap_ies\n\n{go_code}")
            logger.info(
                f"SUCCESS: Wrote {count} ENUMERATED type definitions to 'enumerated_types.go'."
            )

    def _generate_bit_string_types_file(self):
        count = 0

        go_file_content = "package e1ap_ies\n\n"

        bit_strings = sorted(
            [item for item in self.definitions.values() if item.type == "BIT STRING"],
            key=lambda x: self._standard_string(x.name),
        )

        for item in bit_strings:
            go_name = self._standard_string(item.name)

            go_code_block = ""
            go_code_block += f"// {go_name} represents the ASN.1 BIT STRING type.\n"
            go_code_block += f"type {go_name} []byte\n\n"

            if item.min_val is not None:
                go_code_block += "const (\n"
                min_val_str = self._format_go_value(item.min_val)
                max_val_str = self._format_go_value(item.max_val)
                go_code_block += f"\t{go_name}MinSize int = {min_val_str}\n"
                go_code_block += f"\t{go_name}MaxSize int = {max_val_str}\n"
                go_code_block += ")\n"

            go_file_content += go_code_block + "\n\n"
            count += 1

        if count > 0:
            file_path = os.path.join(self.output_dir, "bit_string_types.go")
            with open(file_path, "w", encoding="utf-8") as f:
                f.write(go_file_content)
            logger.info(
                f"SUCCESS: Wrote {count} BIT STRING type definitions to 'bit_string_types.go'."
            )

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