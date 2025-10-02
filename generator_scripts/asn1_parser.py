import re
import logging
from typing import Dict, List, Optional, Any, Tuple
from common_types import ASN1Definition

logger = logging.getLogger(__name__)


class ASN1Parser:
    def __init__(self, lines: List[Tuple[str, str, int]]):
        self.lines = lines
        self.definitions: Dict[str, ASN1Definition] = {}
        self.ie_sets: Dict[str, List[Dict[str, str]]] = {}
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

    def _parse_ie_set(self, name: str, def_part: str):
        """Parses an E1AP-PROTOCOL-IES block into a structured list."""
        # --- START: THE FIX - Capture criticality and presence ---
        ie_pattern = re.compile(
            r"\{\s*ID\s+(?P<id>[\w-]+)\s+CRITICALITY\s+(?P<crit>[\w-]+)\s+TYPE\s+(?P<type>[\w-]+)\s+PRESENCE\s+(?P<pres>[\w-]+)\s*\}"
        )
        # --- END: THE FIX ---
        
        matches = ie_pattern.finditer(def_part)
        
        ie_list = []
        for match in matches:
            # --- START: THE FIX - Save the new data ---
            ie_list.append({
                "id": match.group("id"),
                "type": match.group("type"),
                "crit": match.group("crit").lower(),
                "pres": match.group("pres").lower()
            })
            # --- END: THE FIX ---
        
        if ie_list:
            self.ie_sets[name] = ie_list
            logger.info(f"Successfully parsed IE Set '{name}' with {len(ie_list)} items.")

    def _parse_constraints(self, line_part: str) -> Dict[str, Any]:
        constraints = {"min_val": None, "max_val": None}
        
        # Search for content within any kind of parentheses or braces
        match = re.search(r"\(.+\)|(\{.+\})", line_part)
        if not match:
            return constraints
        content = match.group(0)

        # Priority 1: Look for SIZE(min..max)
        size_range_match = re.search(r"SIZE\s*\(([\w\d-]+)\.\.([\w\d-]+)\)", content)
        if size_range_match:
            constraints["min_val"] = size_range_match.group(1)
            constraints["max_val"] = size_range_match.group(2)
            return constraints

        # Priority 2: Look for SIZE(fixed_value)
        size_fixed_match = re.search(r"SIZE\s*\(([\w\d-]+)\)", content)
        if size_fixed_match:
            val = size_fixed_match.group(1)
            constraints["min_val"] = val
            constraints["max_val"] = val
            return constraints

        # Priority 3: Look for a simple range (min..max)
        range_match = re.search(r"\(([\w\d-]+)\.\.([\w\d-]+)\)", content)
        if range_match:
            min_val = range_match.group(1)
            max_val = range_match.group(2)
            # Filter out junk keywords; only accept numbers or valid named constants
            if not min_val.isalpha() or min_val in self.definitions:
                constraints["min_val"] = min_val
            if not max_val.isalpha() or max_val in self.definitions:
                constraints["max_val"] = max_val
            return constraints
            
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
        
        known_keywords = ["INTEGER", "ENUMERATED", "BIT STRING", "OCTET STRING", "SEQUENCE", "CHOICE", "NULL", "OBJECT IDENTIFIER"]
        is_known_keyword_def = any(def_part_stripped.startswith(k) for k in known_keywords)

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
            item = ASN1Definition(name, "ENUMERATED", source_file, source_line, full_text)
            match = re.search(r"ENUMERATED\s*\{([^}]+)\}", def_part, re.DOTALL)
            if match:
                content = match.group(1).replace("...", "").strip()
                item.enum_values = [v.strip() for v in content.split(",") if v.strip()]
            else:
                item.enum_values = ["present"] # Fallback for abstract types
            return item
        
        elif "BIT STRING" in def_part_stripped or "OCTET STRING" in def_part_stripped:
            # Determine the base type for the ASN1Definition object
            base_type = "BIT STRING" if "BIT STRING" in def_part_stripped else "OCTET STRING"
            item = ASN1Definition(name, base_type, source_file, source_line, full_text)
            
            # This was the missing call. Now we parse the constraints.
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

        container_match = re.search(r"ProtocolIE-Container\s*\{\s*\{\s*([\w-]+)\s*\}\s*\}", def_part)
        if container_match:
            ie_set_name = container_match.group(1)
            
            if ie_set_name in self.ie_sets:
                logger.info(f"Expanding IE container for '{name}' using IE Set '{ie_set_name}'.")
                # This is a PDU definition. Expand the IE set into members.
                for ie in self.ie_sets[ie_set_name]:
                    ie_name_to_add = ie["id"].replace("id-", "")
                    type_name_to_add = ie["type"]
                    if not ie_name_to_add:
                        continue
                    
                    # --- START: THE FIX - Use the parsed presence and criticality ---
                    item.ies.append({
                        "ie": ie_name_to_add,
                        "type": type_name_to_add,
                        "presence": ie["pres"], # Use parsed value
                        "criticality": ie["crit"] # Add new criticality field
                    })
                return item # We are done with this definition
            else:
                logger.warning(f"Could not find pre-parsed IE Set '{ie_set_name}' for container '{name}'.")

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

        member_lines = []
        current_line = ""
        nesting_level = 0
        for char in block_content:
            if char in ['{', '(']:
                nesting_level += 1
            elif char in ['}', ')']:
                nesting_level -= 1
            
            # Only split on a comma if we are at the top level (not nested).
            if char == ',' and nesting_level == 0:
                member_lines.append(current_line)
                current_line = ""
            else:
                current_line += char
        member_lines.append(current_line)
        
        for member_line in member_lines:
            clean_line = member_line.strip()
            if (
                not clean_line
                or clean_line.startswith("...")
                or "iE-Extensions" in clean_line
                or "choice-extension" in clean_line 
            ):
                continue

            parts = clean_line.split(None, 1)
            if len(parts) < 2:
                continue

            member_name, raw_type_str = parts[0], parts[1]
            presence = self._parse_presence(raw_type_str)
            inline_keywords = ["INTEGER", "ENUMERATED", "BIT STRING", "OCTET STRING", "SEQUENCE", "CHOICE"]
            is_inline = any(raw_type_str.strip().startswith(keyword) for keyword in inline_keywords)

            is_inline_enum = raw_type_str.strip().startswith("ENUMERATED")
            is_inline_struct = any(k in raw_type_str for k in ["SEQUENCE {", "CHOICE {"])
            is_simple_inline = any(raw_type_str.strip().startswith(k) for k in ["INTEGER", "BIT STRING", "OCTET STRING"])
            ie_name_to_add = member_name
            type_name_to_add = ""

            if is_inline_enum:
                synthetic_type_name = f"{name}-{member_name}"
                inline_def = ASN1Definition(synthetic_type_name, "ENUMERATED", source_file, source_line, raw_type_str)

                # Use our robust regex directly on the raw type string
                enum_match = re.search(r"ENUMERATED\s*\{([^}]+)\}", raw_type_str, re.DOTALL)
                if enum_match:
                    content = enum_match.group(1).replace("...", "").strip()
                    values = [v.strip() for v in content.split(",") if v.strip()]
                    inline_def.enum_values = values if values else ["inline-parse-failed"]
                else:
                    inline_def.enum_values = ["present"] # Fallback for abstract enums

                if inline_def.name not in self.definitions:
                    self.definitions[inline_def.name] = inline_def
                type_name_to_add = inline_def.name
            # --- END: NEW INLINE ENUM LOGIC ---
            
            elif is_inline_struct:
                synthetic_type_name = f"{name}-{member_name}"
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
        
        # =====================================================================
        # PASS 1: PARSE ONLY THE IE SETS (the "manifests")
        # This pass populates self.ie_sets so that PDU containers can be expanded correctly in Pass 2,
        # regardless of the definition order in the source files.
        # =====================================================================
        logger.info("PASS 1: Scanning for and parsing all IE Set definitions...")
        for i in range(len(self.lines)):
            line_text, _, _ = self.lines[i]
            line = line_text.strip()
            # Quickly identify potential IE Set definitions
            if "::=" in line and "E1AP-PROTOCOL-IES" in line:
                full_def_str, _ = self._extract_full_definition(i)
                name_part, def_part = [p.strip() for p in full_def_str.split("::=", 1)]
                
                # Check again to be sure, then parse it
                if "E1AP-PROTOCOL-IES" in name_part:
                    ie_set_name = name_part.replace("E1AP-PROTOCOL-IES", "").strip()
                    self._parse_ie_set(ie_set_name, def_part)
        logger.info(f"PASS 1 COMPLETE: Found and parsed {len(self.ie_sets)} IE Sets.")

        # =====================================================================
        # PASS 2: PARSE ALL OTHER DEFINITIONS
        # This pass parses all concrete types. When it encounters a PDU container,
        # it will now be able to find its corresponding IE Set in self.ie_sets.
        # =====================================================================
        logger.info("PASS 2: Parsing all other ASN.1 definitions...")
        i = 0
        while i < len(self.lines):
            line_text, source_file, source_line = self.lines[i]
            line = line_text.strip()

            # Handle malformed definitions that the original code checked for
            if ("::=" in line and "{" in line.split("::=")[0] and "}" in line.split("::=")[0]):
                failures.append({
                    "name": line.split("::=")[0].strip(), "text": line,
                    "file": source_file, "line": source_line,
                })
                i += 1
                continue

            if "::=" in line:
                full_def_str, end_index = self._extract_full_definition(i)
                
                name_part, def_part = [p.strip() for p in full_def_str.split("::=", 1)]
                item = None

                # In Pass 2, we SKIP the IE Set definitions as they are already processed.
                if "E1AP-PROTOCOL-IES" in name_part:
                    i = end_index + 1
                    continue

                abstract_keywords = [
                    "E1AP-ELEMENTARY-PROCEDURE", "E1AP-PROTOCOL-EXTENSION",
                    "E1AP-PRIVATE-IES", "CLASS", "InitiatingMessage", 
                    "SuccessfulOutcome", "UnsuccessfulOutcome", "E1AP-PDU"
                ]
                if any(keyword in full_def_str for keyword in abstract_keywords):
                    if name_part not in [f["name"] for f in failures]:
                        failures.append({
                            "name": name_part, "text": full_def_str,
                            "file": source_file, "line": source_line,
                        })
                    i = end_index + 1
                    continue
                
                # --- The rest of the parsing logic is the same as your original code ---
                is_sequence_of = "SEQUENCE OF " in def_part or def_part.startswith("SEQUENCE (") and " OF " in def_part
                if is_sequence_of:
                    item = ASN1Definition(name_part, "LIST", source_file, source_line, full_def_str)
                    constraints = self._parse_constraints(def_part)
                    item.min_val, item.max_val = constraints["min_val"], constraints["max_val"]
                    split_parts = def_part.split(" OF ", 1)
                    if len(split_parts) > 1:
                        of_part = split_parts[1].strip()
                    if of_part.startswith("SEQUENCE {"):
                        synthetic_name = name_part.replace("-List", "-Item")
                        if synthetic_name == name_part:
                            synthetic_name = name_part + "Item"
                        logger.info(f"Detected inline SEQUENCE for LIST '{name_part}'. Creating synthetic type '{synthetic_name}'.")
                        inline_def = self._parse_block_definition(synthetic_name, of_part, source_file, source_line, full_def_str)
                        if inline_def:
                            inline_def.is_synthetic = True
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
                elif ("SEQUENCE" in def_part or "CHOICE" in def_part) and "{" in def_part:
                    # This call will now succeed for PDU containers because self.ie_sets is populated.
                    item = self._parse_block_definition(name_part, def_part, source_file, source_line, full_def_str)
                else:
                    item = self._parse_simple_definition(name_part, def_part, source_file, source_line, full_def_str)

                if item:
                    if item.name not in self.definitions:
                        self.definitions[item.name] = item
                elif "AUTOMATIC TAGS" not in full_def_str and not any(keyword in full_def_str for keyword in abstract_keywords):
                    if name_part not in [f["name"] for f in failures]:
                        failures.append({
                            "name": name_part, "text": full_def_str,
                            "file": source_file, "line": source_line,
                        })
                
                # Set index to the end of the processed block
                i = end_index
            i += 1
        logger.info("PASS 2 COMPLETE.")
        return self.definitions, failures