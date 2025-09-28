import re
import logging
from typing import Dict, List, Optional, Any, Tuple
from common_types import ASN1Definition

logger = logging.getLogger(__name__)


class ASN1Parser:
    def __init__(self, lines: List[Tuple[str, str, int]]):
        self.lines = lines
        self.definitions: Dict[str, ASN1Definition] = {}

    def _parse_constraints(self, line_part: str) -> Dict[str, Any]:
        constraints = {"min_val": None, "max_val": None}

        line_part_no_braces = re.sub(r"\{.*?\}", "", line_part)

        pattern = r"\b(?![A-Z]{2,}\b)(\w+)\b"
        tokens = re.findall(pattern, line_part_no_braces)

        if not tokens:
            return constraints

        constraints["min_val"] = tokens[0]
        constraints["max_val"] = tokens[-1]

        return constraints

    def _parse_presence(self, line_part: str) -> str:
        if "OPTIONAL" in line_part:
            return "optional"
        if "CONDITIONAL" in line_part:
            return "conditional"
        return "mandatory"

    def _extract_full_definition(self, start_index: int) -> Tuple[str, int]:
        start_line = self.lines[start_index][0].strip()

        if "{" not in start_line:
            return start_line, start_index

        full_def_lines = [start_line]
        current_index = start_index + 1

        open_braces = start_line.count("{") - start_line.count("}")

        while current_index < len(self.lines) and open_braces > 0:
            next_line = self.lines[current_index][0].strip()

            if "::=" in next_line:
                break

            if not next_line.startswith("--"):
                full_def_lines.append(next_line)
                open_braces += next_line.count("{") - next_line.count("}")

            current_index += 1

        full_def_str = " ".join(full_def_lines)
        end_index = current_index - 1

        return full_def_str.replace("\n", " ").replace("\r", " "), end_index

    def _parse_simple_definition(
        self,
        name: str,
        def_part: str,
        source_file: str,
        source_line: int,
        full_text: str,
    ) -> Optional[ASN1Definition]:
        def_part_stripped = def_part.strip()
        name_parts = name.strip().split()

        if len(name_parts) >= 2 or def_part_stripped.isdigit():

            const_type = name_parts[-1] if len(name_parts) >= 2 else "INTEGER"
            const_name = " ".join(name_parts[:-1]) if len(name_parts) >= 2 else name

            item = ASN1Definition(
                const_name, const_type, source_file, source_line, full_text
            )
            item.min_val, item.max_val = def_part_stripped, def_part_stripped
            item.is_constant = True
            return item

        if def_part_stripped.startswith("INTEGER"):

            if name in ["ProcedureCode", "ProtocolIE-ID", "ProtocolExtensionID"]:
                item = ASN1Definition(
                    name, "BASE_TYPE", source_file, source_line, full_text
                )
            else:

                item = ASN1Definition(
                    name, "INTEGER", source_file, source_line, full_text
                )

            constraints = self._parse_constraints(def_part)
            item.min_val, item.max_val = constraints["min_val"], constraints["max_val"]
            return item

        elif def_part_stripped.startswith("ENUMERATED"):
            item = ASN1Definition(
                name, "ENUMERATED", source_file, source_line, full_text
            )
            match = re.search(r"ENUMERATED\s*\{(.*?)\}", def_part, re.DOTALL)
            if match:
                content = match.group(1).replace("...", "").strip()
                item.enum_values = [v.strip() for v in content.split(",") if v.strip()]
            else:
                item.enum_values = ["present"]
            return item

        elif "BIT STRING" in def_part_stripped:
            item = ASN1Definition(
                name, "BIT STRING", source_file, source_line, full_text
            )

            constraints = self._parse_constraints(def_part)
            item.min_val, item.max_val = constraints["min_val"], constraints["max_val"]
            return item

        value_parts = def_part_stripped.split()
        if len(value_parts) == 1 and value_parts[0][0].isupper():
            item = ASN1Definition(name, "ALIAS", source_file, source_line, full_text)
            item.alias_of = value_parts[0]
            return item

        return None

    def _parse_block_definition(
        self,
        name: str,
        def_part: str,
        source_file: str,
        source_line: int,
        full_text: str,
    ) -> Optional[ASN1Definition]:
        def_type = "SEQUENCE" if "SEQUENCE" in def_part else "CHOICE"
        item = ASN1Definition(name, def_type, source_file, source_line, full_text)
        item.is_extensible = "..." in def_part

        start_brace_index = def_part.find("{")
        if start_brace_index == -1:
            return item
        brace_level, end_brace_index = 1, -1
        for i in range(start_brace_index + 1, len(def_part)):
            if def_part[i] == "{":
                brace_level += 1
            elif def_part[i] == "}":
                brace_level -= 1
            if brace_level == 0:
                end_brace_index = i
                break
        if end_brace_index == -1:
            return item
        block_content = def_part[start_brace_index + 1 : end_brace_index].strip()

        member_lines = re.split(r",(?![^(){}]*\))", block_content)
        for member_line in member_lines:
            clean_line = member_line.strip()
            if (
                not clean_line
                or clean_line.startswith("...")
                or "iE-Extensions" in clean_line
            ):
                continue

            parts = clean_line.split()
            if len(parts) < 2:
                continue

            member_name, raw_type_str = parts[0], " ".join(parts[1:])
            presence = self._parse_presence(raw_type_str)

            inline_keywords = ["INTEGER", "ENUMERATED", "BIT STRING", "OCTET STRING"]
            is_inline = any(keyword in raw_type_str for keyword in inline_keywords)

            ie_name_to_add = member_name
            type_name_to_add = ""

            if is_inline:

                synthetic_type_name = member_name

                if synthetic_type_name not in self.definitions:
                    inline_def = self._parse_simple_definition(
                        synthetic_type_name,
                        raw_type_str,
                        source_file,
                        source_line,
                        raw_type_str,
                    )
                    if inline_def:
                        self.definitions[inline_def.name] = inline_def

                type_name_to_add = synthetic_type_name
            else:

                type_str_cleaned = (
                    raw_type_str.replace("OPTIONAL", "")
                    .replace("CONDITIONAL", "")
                    .strip()
                )
                type_parts = type_str_cleaned.split()
                if not type_parts:
                    continue
                type_name_to_add = type_parts[0]

            if ie_name_to_add and type_name_to_add:

                item.ies.append(
                    {
                        "ie": ie_name_to_add,
                        "type": type_name_to_add,
                        "presence": presence,
                    }
                )

        return item

    def parse(self) -> Tuple[Dict[str, ASN1Definition], List[Dict[str, Any]]]:
        failures = []
        i = 0
        in_abstract_block = False

        class_macro_pattern = re.compile(
            r"{\s*ID\s+([\w-]+)\s+CRITICALITY\s+[\w-]+\s+TYPE\s+([^}]+?)\s+PRESENCE\s+[\w-]+\s*}\s*[,|]?"
        )

        while i < len(self.lines):
            line_text, source_file, source_line = self.lines[i]
            line = line_text.strip()

            if (
                "::=" in line
                and "{" in line.split("::=")[0]
                and "}" in line.split("::=")[0]
            ):
                failures.append(
                    {
                        "name": line.split("::=")[0].strip(),
                        "text": line,
                        "file": source_file,
                        "line": source_line,
                    }
                )
                i += 1
                continue

            if "::=" in line:
                full_def_str, end_index = self._extract_full_definition(i)
                i = end_index

                name_part, def_part = [p.strip() for p in full_def_str.split("::=", 1)]
                item = None

                abstract_keywords = [
                    "E1AP-ELEMENTARY-PROCEDURE",
                    "E1AP-PROTOCOL-IES",
                    "E1AP-PROTOCOL-EXTENSION",
                    "E1AP-PRIVATE-IES",
                    "CLASS",
                ]
                if any(keyword in full_def_str for keyword in abstract_keywords):
                    if name_part not in [f["name"] for f in failures]:
                        failures.append(
                            {
                                "name": name_part,
                                "text": full_def_str,
                                "file": source_file,
                                "line": source_line,
                            }
                        )
                elif "SEQUENCE" in def_part and "OF" in def_part:
                    item = ASN1Definition(
                        name_part, "LIST", source_file, source_line, full_def_str
                    )
                    constraints = self._parse_constraints(def_part)
                    item.min_val, item.max_val = (
                        constraints["min_val"],
                        constraints["max_val"],
                    )
                elif (
                    "SEQUENCE" in def_part or "CHOICE" in def_part
                ) and "{" in def_part:
                    item = self._parse_block_definition(
                        name_part, def_part, source_file, source_line, full_def_str
                    )
                else:
                    item = self._parse_simple_definition(
                        name_part, def_part, source_file, source_line, full_def_str
                    )

                if item:
                    if item.name not in self.definitions:
                        self.definitions[item.name] = item
                elif "AUTOMATIC TAGS" not in full_def_str and not any(
                    keyword in full_def_str for keyword in abstract_keywords
                ):
                    if name_part not in [f["name"] for f in failures]:
                        failures.append(
                            {
                                "name": name_part,
                                "text": full_def_str,
                                "file": source_file,
                                "line": source_line,
                            }
                        )
            i += 1

        return self.definitions, failures