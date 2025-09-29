import re
import logging
from typing import Dict, List, Optional, Any, Tuple
from common_types import ASN1Definition

logger = logging.getLogger(__name__)


class ASN1Parser:
    def __init__(self, lines: List[Tuple[str, str, int]]):
        self.lines = lines
        self.definitions: Dict[str, ASN1Definition] = {}
        self._prepopulate_builtin_types()

    def _prepopulate_builtin_types(self):
        """Creates ASN1Definition objects for built-in ASN.1 types."""
        builtins = {
            "INTEGER": "INTEGER",
            "ENUMERATED": "ENUMERATED",
            "BIT STRING": "BIT STRING",
            "OCTET STRING": "OCTET STRING",
            "PrintableString": "PrintableString",
            "VisibleString": "VisibleString",
            "UTF8String": "UTF8String",
            "SEQUENCE": "SEQUENCE",
            "CHOICE": "CHOICE",
            "NULL": "NULL",
            "OCTET": "OCTET STRING",
            "STRING": "OCTET STRING",
            "OBJECT IDENTIFIER": "OBJECT IDENTIFIER",
        }

        for name, def_type in builtins.items():
            item = ASN1Definition(name=name, def_type=def_type, source_file="builtin")
            item.is_builtin = True
            self.definitions[name] = item

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
        if value_parts:
            base_type = value_parts[0].split('(')[0].strip()
            is_complex_list = "SEQUENCE" in def_part_stripped and "OF" in def_part_stripped
            is_potential_alias = base_type[0].isupper() and "{" not in def_part_stripped and not is_complex_list
            
            if is_potential_alias:
                item = ASN1Definition(name, "ALIAS", source_file, source_line, full_text)
                item.alias_of = base_type
                
                constraints = self._parse_constraints(def_part)
                item.min_val, item.max_val = constraints["min_val"], constraints["max_val"]
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
                or "choice-extension" in clean_line 
            ):
                continue

            parts = clean_line.split()
            if len(parts) < 2:
                continue

            member_name, raw_type_str = parts[0], " ".join(parts[1:])
            presence = self._parse_presence(raw_type_str)

            
            
            
            inline_keywords = ["INTEGER", "ENUMERATED", "BIT STRING", "OCTET STRING", "SEQUENCE", "CHOICE"]
            is_inline = any(raw_type_str.strip().startswith(keyword) for keyword in inline_keywords)

            ie_name_to_add = member_name
            type_name_to_add = ""

            if is_inline:
                
                
                synthetic_type_name = f"{name}-{member_name}"
                
                
                inline_def = None
                if raw_type_str.strip().startswith("SEQUENCE") and "OF" in raw_type_str:
                    inline_def = self._parse_simple_definition(synthetic_type_name, raw_type_str, source_file, source_line, raw_type_str)
                else:
                    inline_def = self._parse_simple_definition(synthetic_type_name, raw_type_str, source_file, source_line, raw_type_str)
                    if not inline_def:
                        
                        inline_def = self._parse_block_definition(synthetic_type_name, raw_type_str, source_file, source_line, raw_type_str)
                
                if inline_def:
                    if inline_def.name not in self.definitions:
                        self.definitions[inline_def.name] = inline_def
                    type_name_to_add = inline_def.name

            else:
                
                
                base_type_name = raw_type_str
                if '{' in base_type_name:
                    base_type_name = base_type_name.split('{', 1)[0]
                if '(' in base_type_name:
                    base_type_name = base_type_name.split('(', 1)[0]
                
                base_type_name = base_type_name.replace("OPTIONAL", "").replace("CONDITIONAL", "").strip()
                type_parts = base_type_name.split()
                if not type_parts:
                    continue
                type_name_to_add = " ".join(type_parts)
            
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
                    "E1AP-ELEMENTARY-PROCEDURE", "E1AP-PROTOCOL-IES",
                    "E1AP-PROTOCOL-EXTENSION", "E1AP-PRIVATE-IES", "CLASS",
                    "InitiatingMessage", "SuccessfulOutcome", "UnsuccessfulOutcome",
                    "E1AP-PDU" 
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
                    continue
                is_sequence_of = "SEQUENCE OF " in def_part or def_part.startswith("SEQUENCE (") and " OF " in def_part

                if is_sequence_of:
                    
                    item = ASN1Definition(
                        name_part, "LIST", source_file, source_line, full_def_str
                    )
                    constraints = self._parse_constraints(def_part)
                    item.min_val, item.max_val = (
                        constraints["min_val"],
                        constraints["max_val"],
                    )
                    split_parts = def_part.split(" OF ", 1)
                    if len(split_parts) > 1:
                        of_part = split_parts[1].strip()

                    if of_part.startswith("SEQUENCE {"):
                        
                        
                        synthetic_name = name_part.replace("-List", "-Item")
                        if synthetic_name == name_part: 
                            synthetic_name = name_part + "Item"
                        
                        logger.info(f"Detected inline SEQUENCE for LIST '{name_part}'. Creating synthetic type '{synthetic_name}'.")

                        
                        
                        inline_def = self._parse_block_definition(
                            synthetic_name, of_part, source_file, source_line, full_def_str
                        )
                        if inline_def:
                            
                            if inline_def.name not in self.definitions:
                                self.definitions[inline_def.name] = inline_def
                            
                            
                            item.of_type = inline_def.name
                        else:
                            logger.warning(f"Failed to parse inline SEQUENCE for LIST: {name_part}")

                    else:
                        
                        
                        match = re.search(r"([\w-]+)", of_part)
                        if match:
                            item.of_type = match.group(1)
                        else:
                            logger.warning(f"Could not determine the 'OF' type for LIST: {name_part}")
                    
                    
                    match = re.search(r"SEQUENCE(?:.*?)OF\s+([\w-]+)", def_part)
                    if match:
                        item.of_type = match.group(1)
                    else:
                        logger.warning(f"Could not determine the 'OF' type for LIST: {name_part}")
                
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