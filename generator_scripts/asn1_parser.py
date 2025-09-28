# In asn1_parser.py
import re
import logging
from typing import Dict, List, Optional, Any, Tuple

logger = logging.getLogger(__name__)

class ASN1Definition:
    def __init__(self, name: str, def_type: str, source_file: str = "", source_line: int = -1, full_text: str = ""):
        self.name, self.type, self.source_file, self.source_line, self.full_text = name, def_type, source_file, source_line, full_text
        self.min_val: Optional[str] = None
        self.max_val: Optional[str] = None
        self.enum_values: List[str] = []
        self.ies: List[Dict[str, Any]] = []
        self.list_item_type: Optional[str] = None
        self.is_constant: bool = False
        self.is_extensible: bool = False

class ASN1Parser:
    def __init__(self, lines: List[Tuple[str, str, int]]):
        self.lines = lines
        self.definitions: Dict[str, ASN1Definition] = {}

    def _parse_constraints(self, line_part: str) -> Dict[str, Any]:
        """
        Final, correct version of the constraint parser.
        This regex does not greedily consume '..' as part of a token.
        """
        constraints = {'min_val': None, 'max_val': None}
        
        # A token is a word, possibly with hyphens.
        # This regex looks for a token, then an optional range '..' and another token.
        token_pattern = r"[\w-]+"
        pattern = rf"\((?:SIZE\s*\()?\s*({token_pattern})(?:\s*\.\.\s*({token_pattern}))?.*?\)"
        
        match = re.search(pattern, line_part)
        if match:
            lower_bound = match.group(1)
            upper_bound = match.group(2)
            
            constraints['min_val'] = lower_bound
            constraints['max_val'] = upper_bound if upper_bound is not None else lower_bound
            
        return constraints
    

    def _parse_presence(self, line_part: str) -> str:
        """Helper to determine member presence."""
        if "OPTIONAL" in line_part: return "optional"
        if "CONDITIONAL" in line_part: return "conditional"
        return "mandatory"

    def _extract_full_definition(self, start_index: int) -> Tuple[str, int]:
        """Consumes lines to get a complete single definition, handling continuations."""
        full_def_str = self.lines[start_index][0].strip()
        i = start_index
        if "::=" in full_def_str:
            open_braces = full_def_str.count('{') - full_def_str.count('}')
            while open_braces > 0 and i + 1 < len(self.lines):
                i += 1
                line = self.lines[i][0].strip()
                if not line or line.startswith("--"): continue
                open_braces += line.count('{') - line.count('}')
                full_def_str += " " + line
        return full_def_str.replace('\n', ' ').replace('\r', ' '), i

    def _parse_simple_definition(self, name: str, def_part: str, source_file: str, source_line: int, full_text: str) -> Optional[ASN1Definition]:
        """
        Parses a simple, non-structured type. This version correctly distinguishes
        between a TYPE definition, a CONSTANT definition, and a BASE TYPE definition.
        """
        def_part_stripped = def_part.strip()
        
        # --- PATH 1: This is a TYPE definition ---
        if def_part_stripped.startswith("INTEGER"):
            # Check if this is a special BASE TYPE that other constants rely on.
            if name in ["ProcedureCode", "ProtocolIE-ID", "ProtocolExtensionID"]:
                # Correctly classify it as a BASE_TYPE, not a regular INTEGER.
                item = ASN1Definition(name, "BASE_TYPE", source_file, source_line, full_text)
                # We can still parse its constraints if needed, but the type is what matters.
                constraints = self._parse_constraints(def_part)
                item.min_val, item.max_val = constraints['min_val'], constraints['max_val']
                return item

            # This is a legitimate new integer type.
            item = ASN1Definition(name, "INTEGER", source_file, source_line, full_text)
            constraints = self._parse_constraints(def_part)
            item.min_val, item.max_val = constraints['min_val'], constraints['max_val']
            return item

        elif def_part_stripped.startswith("ENUMERATED"):
            item = ASN1Definition(name, "ENUMERATED", source_file, source_line, full_text)
            
            # Use a robust regex to find the content inside the braces
            match = re.search(r'ENUMERATED\s*\{(.*?)\}', def_part, re.DOTALL)
            
            if match:
                # This is the standard case with an explicit value list {...}
                content = match.group(1).replace("...", "").strip()
                values = [v.strip() for v in content.split(',') if v.strip()]
                item.enum_values = values
            else:
                # --- THIS IS THE FIX ---
                # This handles the case where the definition is just "ENUMERATED"
                # without a value list, which is common for boolean flags.
                # We will create a default value, which is a standard practice.
                item.enum_values = ["present"]

            return item
        
        # (Add back BIT STRING, OCTET STRING etc. here using the same .startswith() logic)

        # --- PATH 2: This must be a CONSTANT definition ---
        else:
            name_parts = name.strip().split()
            if len(name_parts) < 2: return None
            const_type = name_parts[-1]
            const_name = " ".join(name_parts[:-1])
            const_value = def_part_stripped
            
            item = ASN1Definition(const_name, const_type, source_file, source_line, full_text)
            item.min_val = item.max_val = const_value
            item.is_constant = True
            return item
    

    def _parse_block_definition(self, name: str, def_part: str, source_file: str, source_line: int, full_text: str) -> Optional[ASN1Definition]:
        """
        Parses a SEQUENCE or CHOICE block. This version correctly handles inline
        definitions and is robust against malformed member lines that could
        otherwise cause an IndexError.
        """
        def_type = "SEQUENCE" if "SEQUENCE" in def_part else "CHOICE"
        item = ASN1Definition(name, def_type, source_file, source_line, full_text)
        item.is_extensible = "..." in def_part
        
        block_match = re.search(r'\{(.*)\}', def_part, re.DOTALL)
        if not block_match: return item
        
        block_content = block_match.group(1).strip()
        member_lines = re.split(r',(?![^(){}]*\))', block_content)
        
        for member_line in member_lines:
            clean_line = member_line.strip()
            if not clean_line or clean_line.startswith("...") or "iE-Extensions" in clean_line:
                continue
            
            parts = clean_line.split()
            if len(parts) < 2: continue
            
            member_name = parts[0]
            raw_type_str = " ".join(parts[1:])
            presence = self._parse_presence(raw_type_str)
            
            inline_keywords = ["INTEGER", "ENUMERATED", "BIT STRING", "OCTET STRING"]
            is_inline = any(keyword in raw_type_str for keyword in inline_keywords)
            
            member_type_name = ""
            if is_inline:
                synthetic_name = member_name
                inline_def = self._parse_simple_definition(synthetic_name, raw_type_str, source_file, source_line, raw_type_str)
                
                if inline_def:
                    if inline_def.name not in self.definitions:
                        self.definitions[inline_def.name] = inline_def
                    member_type_name = inline_def.name
            else:
                # It's a reference to an existing type.
                type_str_cleaned = raw_type_str.replace("OPTIONAL", "").replace("CONDITIONAL", "").strip()
                
                # --- THIS IS THE CRASH-PROOF FIX ---
                type_parts = type_str_cleaned.split()
                if not type_parts:
                    # This happens if the line was something like "memberName OPTIONAL" with no type.
                    # It's malformed or a syntax we don't handle. We log it and skip it safely.
                    logger.warning(f"Could not determine type for member '{member_name}' in '{name}'. Skipping line: '{clean_line}'")
                    continue
                
                member_type_name = type_parts[0]
            
            if member_type_name:
                item.ies.append({"ie": member_name, "type": member_type_name, "presence": presence})
                
        return item

    def parse(self) -> Tuple[Dict[str, ASN1Definition], List[Dict[str, Any]]]:
        """
        The main parsing dispatcher. This FINAL version includes a check to
        explicitly identify and skip advanced parameterized type definitions,
        preventing them from being mis-parsed.
        """
        failures = []
        i = 0
        in_abstract_block = False

        class_macro_pattern = re.compile(
            r"{\s*ID\s+([\w-]+)\s+CRITICALITY\s+\w+\s+TYPE\s+([^}]+?)\s+PRESENCE\s+\w+\s*}"
        )

        while i < len(self.lines):
            line_text, source_file, source_line = self.lines[i]
            line = line_text.strip()
            
            # --- NEW: Check for and skip parameterized definitions ---
            # A parameterized definition contains '{' and '}' on the LEFT side of '::='
            if "::=" in line and "{" in line.split("::=")[0] and "}" in line.split("::=")[0]:
                param_name = line.split("::=")[0].strip()
                failures.append({"name": param_name, "text": line, "file": source_file, "line": source_line})
                i += 1
                continue # Skip this line entirely

            line = line.replace("...", "")
            if i + 1 < len(self.lines) and not self.lines[i+1][0].strip().startswith('{') and not '::=' in self.lines[i+1][0]:
                line += " " + self.lines[i+1][0].strip()

            abstract_containers = ["E1AP-PROTOCOL-IES", "E1AP-ELEMENTARY-PROCEDURE"]
            if "::= {" in line and any(keyword in line for keyword in abstract_containers):
                in_abstract_block = True
                failures.append({"name": line.split("::=")[0].strip(), "text": line, "file": source_file, "line": source_line})
                i += 1
                continue

            if line.startswith("}") and in_abstract_block:
                in_abstract_block = False
                i += 1
                continue

            target_line_text = line
            is_synthetic = False

            if in_abstract_block:
                match = class_macro_pattern.search(target_line_text)
                if match:
                    id_name = match.group(1).strip()
                    type_def = match.group(2).strip()
                    target_line_text = f"{id_name} ::= {type_def}"
                    is_synthetic = True
                else:
                    i += 1
                    continue

            if "::=" in target_line_text:
                if not is_synthetic:
                    full_def_str, end_index = self._extract_full_definition(i)
                    i = end_index
                else:
                    full_def_str = target_line_text

                name_part, def_part = [p.strip() for p in full_def_str.split("::=", 1)]
                item = None

                keywords_to_always_skip = ["E1AP-PROTOCOL-EXTENSION", "E1AP-PRIVATE-IES", "CLASS"]
                if any(keyword in full_def_str for keyword in keywords_to_always_skip):
                    if name_part not in [f['name'] for f in failures]:
                        failures.append({"name": name_part, "text": full_def_str, "file": source_file, "line": source_line})
                elif "SEQUENCE" in def_part and "OF" in def_part:
                    item = ASN1Definition(name_part, "LIST", source_file, source_line, full_def_str)
                    constraints = self._parse_constraints(def_part)
                    item.min_val, item.max_val = constraints['min_val'], constraints['max_val']
                elif ("SEQUENCE" in def_part or "CHOICE" in def_part) and "{" in def_part:
                    item = self._parse_block_definition(name_part, def_part, source_file, source_line, full_def_str)
                else:
                    item = self._parse_simple_definition(name_part, def_part, source_file, source_line, full_def_str)

                if item:
                    if item.name not in self.definitions:
                        self.definitions[item.name] = item
                elif "AUTOMATIC TAGS" not in full_def_str and not any(keyword in full_def_str for keyword in keywords_to_always_skip):
                    if name_part not in [f['name'] for f in failures]:
                        failures.append({"name": name_part, "text": full_def_str, "file": source_file, "line": source_line})
            i += 1

        return self.definitions, failures