import os
import re
import logging
import shutil
from typing import Dict, List, Any, Tuple, Optional
from go_templates import render_enum, render_sequence, render_choice
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
            "ID",
            "IE",
            "PDU",
            "E1AP",
            "GNB",
            "CU",
            "CP",
            "UP",
            "QOS",
            "DRB",
            "IP",
            "ASN",
            "TEID",
            "URI",
            "SCTP",
            "GTP",
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

        visited = {type_name}

        while isinstance(definition, AliasDefinition):
            next_alias = definition.alias_of
            if next_alias in visited:
                logger.warning(
                    f"Circular alias dependency detected for '{type_name}' at '{next_alias}'."
                )
                break
            visited.add(next_alias)

            next_def = self.definitions.get(next_alias)
            if not next_def:
                logger.warning(
                    f"Alias '{next_alias}' for type '{definition.name}' not found in definitions."
                )

                break
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

            base_name = self._standard_string(item.name)
            value = self._format_go_value(item.min_val)

            if item.const_type == "ProcedureCode":
                proc_code_consts.append((f"ProcedureCode_{base_name}", value))
            elif item.const_type == "ProtocolIE-ID":
                protocol_ie_consts.append((f"ProtocolIEID_{base_name}", value))
            else:
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
        """Generates a single file for all simple, reusable type definitions with validation."""
        if not simple_items:
            logger.info("No simple types found to generate in common file.")
            return

        go_code = 'package e1ap_ies\n\nimport (\n\t"fmt"\n\t"github.com/lvdund/ngap/aper"\n)\n\n'

        from collections import defaultdict

        type_counts = defaultdict(int)
        total_count = 0

        processed_concrete_defs = set()

        sorted_items = sorted(simple_items, key=lambda x: self._standard_string(x.name))

        for item in sorted_items:
            go_name = self._standard_string(item.name)

            if isinstance(item, AliasDefinition):
                go_base_type, _ = self._resolve_go_type(item.alias_of)
            elif isinstance(item, IntegerDefinition):
                go_base_type = "aper.Integer"
            elif isinstance(item, StringDefinition):
                go_base_type = "aper.OctetString"
            else:

                continue

            if go_name.lower() != go_base_type.lower().split(".")[-1]:

                go_code += self._generate_header_comment(item, go_name)
                go_code += f"type {go_name} {go_base_type}\n\n"

                has_constraints = hasattr(item, "min_val") and item.min_val is not None

                is_validatable_integer = (
                    isinstance(item, IntegerDefinition)
                    and item.max_val != "18446744073709551615"
                )

                is_validatable_string = isinstance(item, StringDefinition)

                if has_constraints and (
                    is_validatable_integer or is_validatable_string
                ):

                    min_val_unformatted = item.min_val
                    max_val_unformatted = (
                        item.max_val if item.max_val is not None else item.min_val
                    )

                    min_val_str = self._format_go_value(min_val_unformatted)
                    max_val_str = self._format_go_value(max_val_unformatted)

                    const_type = "int" if is_validatable_string else go_name

                    if min_val_str != max_val_str:

                        min_val_assign = min_val_str
                        max_val_assign = max_val_str
                        try:
                            int(min_val_unformatted)
                        except (ValueError, TypeError):
                            min_val_assign = f"{const_type}({min_val_str})"
                        try:
                            int(max_val_unformatted)
                        except (ValueError, TypeError):
                            max_val_assign = f"{const_type}({max_val_str})"

                        go_code += "const (\n"
                        go_code += f"\tMin{go_name} {const_type} = {min_val_assign}\n"
                        go_code += f"\tMax{go_name} {const_type} = {max_val_assign}\n"
                        go_code += ")\n\n"

                    go_code += f"// Validate checks if the value is within the specified range.\n"
                    go_code += f"func (v {go_name}) Validate() error {{\n"

                    min_check = (
                        f"Min{go_name}"
                        if min_val_str != max_val_str
                        else self._format_go_value(item.min_val)
                    )
                    max_check = (
                        f"Max{go_name}"
                        if min_val_str != max_val_str
                        else self._format_go_value(item.max_val)
                    )

                    check_logic = ""
                    error_range_str = ""

                    if is_validatable_string:
                        check_logic = f"len(v) < {min_check} || len(v) > {max_check}"
                        error_range_str = f"length ({min_check}..{max_check})"
                    else:
                        check_logic = f"v < {min_check} || v > {max_check}"
                        error_range_str = f"range ({min_check}..{max_check})"

                    go_code += f"\tif {check_logic} {{\n"
                    go_code += f'\t\treturn fmt.Errorf("value for {go_name} is outside the valid {error_range_str}")\n'
                    go_code += "\t}\n"
                    go_code += "\treturn nil\n"
                    go_code += "}\n\n"

                total_count += 1

                item_type_str = type(item).__name__.replace("Definition", "").upper()
                type_counts[item_type_str] += 1

            if total_count > 0:
                file_path = os.path.join(self.output_dir, "e1ap_common_types.go")
                with open(file_path, "w", encoding="utf-8") as f:
                    f.write(go_code)
                self.generated_files.add("e1ap_common_types.go")

        for def_type, count in sorted(type_counts.items()):
            if count > 0:
                plural = "s" if count > 1 else ""
                logger.info(
                    f"SUCCESS: Wrote {count} {def_type} type definition{plural} with validation to 'e1ap_common_types.go'."
                )
        else:
            logger.info("No simple type aliases needed to be generated.")

    def _generate_complex_type_files(self, complex_items: List[Any]):
        """Generates a dedicated file for each complex type using a robust two-pass strategy."""
        if not complex_items:
            logger.info("No complex types found to generate individual files for.")
            return

        type_counts = {"SEQUENCE": 0, "CHOICE": 0, "ENUMERATED": 0, "LIST": 0}

        logger.info(
            "Complex types pass 1: Generating SEQUENCE, CHOICE, and ENUMERATED files..."
        )
        for item in complex_items:

            if isinstance(item, (SequenceDefinition, ChoiceDefinition, EnumDefinition)):
                go_name = self._standard_string(item.name)
                filename = self._go_name_to_snake_case(go_name) + ".go"
                file_path = os.path.join(self.output_dir, filename)

                item_code = self._generate_code_for_item(item)

                file_content = "package e1ap_ies\n\n"

                imports = set()
                if isinstance(item, (EnumDefinition, ChoiceDefinition)):
                    imports.add('"github.com/lvdund/ngap/aper"')

                if (
                    isinstance(item, SequenceDefinition)
                    and item.ies
                    and item.ies[0].id is not None
                ):
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

                item_type_str = type(item).__name__.replace("Definition", "").upper()
                if item_type_str in type_counts:
                    type_counts[item_type_str] += 1

        logger.info("Complex types pass 2: Generating LIST files...")
        for item in complex_items:
            if isinstance(item, ListDefinition):
                if not item.of_type:
                    logger.warning(
                        f"Skipping LIST {item.name} because its element type is unknown."
                    )
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
                type_counts["LIST"] += 1

            total_complex_count = sum(type_counts.values())
        if total_complex_count > 0:
            for def_type, count in type_counts.items():
                if count > 0:
                    logger.info(
                        f"SUCCESS: Wrote {count} {def_type} definitions to individual files."
                    )

    def _generate_code_for_item(self, item: Any) -> str:
        """Generates the Go code string for a single complex ASN.1 definition,
        with minimalist placeholder Encode/Decode methods for PDUs."""
        go_name = self._standard_string(item.name)
        item_code = ""

        if isinstance(item, EnumDefinition):

            item_code = render_enum(
                go_name=go_name,
                enum_values=item.enum_values,
                pascal_case_converter=self._standard_string,
            )
        elif isinstance(item, SequenceDefinition):
            item_code = render_sequence(
                go_name=go_name,
                item=item,
                pascal_case_converter=self._standard_string,
                go_type_resolver=self._resolve_go_type,
                go_value_formatter=self._format_go_value,
            )
        elif isinstance(item, ChoiceDefinition):
            item_code = render_choice(
                go_name=go_name,
                item=item,
                pascal_case_converter=self._standard_string,
                go_type_resolver=self._resolve_go_type,
            )
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
