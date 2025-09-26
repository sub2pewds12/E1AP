import re
import logging
from typing import Dict, List, Optional, Any


logger = logging.getLogger(__name__)
# A simple data structure is better than a raw dictionary
class ASN1Definition:
    def __init__(self, name: str, def_type: str):
        self.name = name
        self.type = def_type
        self.ies: List[Dict[str, Any]] = []
        self.list_item_type: Optional[str] = None
        self.enum_values: List[str] = []
        self.min_val: Optional[str] = None
        self.max_val: Optional[str] = None
        self.is_extensible: bool = False
        self.is_constant: bool = False

class ASN1Parser:
    def __init__(self, spec_content: str):
        self.lines = spec_content.splitlines()
        self.definitions: Dict[str, ASN1Definition] = {}
        logger.debug(f"ASN1Parser initialized with {len(self.lines)} lines of spec content.")

    def _parse_constraints(self, line_part: str) -> Dict[str, Any]:
        """Helper to parse (SIZE) or value range constraints."""
        constraints = {'min_val': None, 'max_val': None, 'is_extensible': "..." in line_part}
        match = re.search(r"\((?:SIZE\s*\()?\s*([\w-]+)(?:\s*\.\.\s*([\w-]+))?\s*\)?\)", line_part)
        if match:
            constraints['min_val'] = match.group(1)
            constraints['max_val'] = match.group(2) or match.group(1)
        return constraints

    def _parse_presence(self, line_part: str) -> str:
        """Helper to determine member presence."""
        if "OPTIONAL" in line_part: return "optional"
        if "CONDITIONAL" in line_part: return "conditional"
        return "mandatory"

    def parse(self) -> Dict[str, ASN1Definition]:
        """Parses the entire spec content into a dictionary of structured definitions."""
        i = 0
        while i < len(self.lines):
            line = self.lines[i].strip()
            if not line or line.startswith("--"):
                i += 1
                continue

            if "::=" in line:
                if line.endswith("{"):
                    item, end_index = self._parse_block_definition(i)
                    if item: self.definitions[item.name] = item
                    i = end_index
                else:
                    item, end_index = self._parse_line_definition(i)
                    if item: self.definitions[item.name] = item
                    i = end_index
            i += 1
        return self.definitions

    def _parse_block_definition(self, start_index: int):
        """
        Parses a multi-line SEQUENCE or CHOICE block.
        Crucially, it now IGNORES other block types (like IE Contaiers).
        """
        line = self.lines[start_index].strip()
        name = line.split("::=")[0].strip()
        
        # --- FIX: Be specific. Only handle SEQUENCE and CHOICE blocks. ---
        if "::= SEQUENCE" in line and line.endswith("{"):
            item = ASN1Definition(name, "SEQUENCE")
            logger.debug(f"Parsing SEQUENCE block definition for: {name}")
        elif "::= CHOICE" in line and line.endswith("{"):
            item = ASN1Definition(name, "CHOICE")
            logger.debug(f"Parsing CHOICE block definition for: {name}")
        else:
            # This is not a block type we handle. Ignore it and return.
            logger.debug(f"Skipping unhandled block definition: {name}")
            return None, start_index

        block_content, end_index = self._extract_block(start_index)
        item.is_extensible = "..." in "".join(block_content)

        for member_line in block_content:
            # --- FIX: Defensive initialization to prevent crashes ---
            raw_type_str = ""

            if not member_line or member_line.startswith("--") or "..." in member_line:
                continue
            
            parts = member_line.split()
            if len(parts) < 2: continue

            member_name = parts[0].replace(',', '')
            raw_type_str = " ".join(parts[1:]).split(',')[0].strip()
            member_type = parts[1].replace(',', '')

            inline_base_types = ["BIT STRING", "INTEGER", "ENUMERATED", "OCTET STRING"]
            for base_type in inline_base_types:
                if raw_type_str.startswith(base_type):
                    synthetic_name = f"{name}-{member_name}"
                    logger.debug(f"  Found inline '{base_type}' for member '{member_name}'. Creating synthetic type: {synthetic_name}")
                    
                    new_def = ASN1Definition(synthetic_name, base_type)
                    if base_type in ["BIT STRING", "INTEGER", "OCTET STRING"]:
                        constraints = self._parse_constraints(raw_type_str)
                        new_def.min_val, new_def.max_val = constraints['min_val'], constraints['max_val']
                    elif base_type == "ENUMERATED":
                         match = re.search(r'\{(.*)\}', raw_type_str)
                         if match:
                             new_def.enum_values = [v.strip() for v in match.group(1).replace("...", "").split(',') if v.strip()]
                    
                    self.definitions[synthetic_name] = new_def
                    member_type = synthetic_name
                    break
            
            constraints = self._parse_constraints(member_line)
            member_details = {
                "ie": member_name,
                "type": member_type,
                "presence": self._parse_presence(member_line),
                "min": constraints['min_val'],
                "max": constraints['max_val'],
                "ext": constraints['is_extensible']
            }
            item.ies.append(member_details)
        return item, end_index
    

    
    def _extract_block(self, start_index: int):
        """Safely extracts content between matching braces {}."""
        block_content = []
        nest_level = 1
        j = start_index + 1
        while j < len(self.lines):
            line_in_block = self.lines[j] # Don't strip yet, preserve indentation
            if "{" in line_in_block: nest_level += 1
            if "}" in line_in_block: nest_level -= 1
            if nest_level == 0: break
            block_content.append(line_in_block.strip())
            j += 1
        return block_content, j

    def _parse_line_definition(self, start_index: int):
        """Parses a single-line or simple multi-line definition."""
        full_line = self.lines[start_index].strip()
        if start_index + 1 < len(self.lines) and self.lines[start_index+1].strip().startswith(('(', '{')):
            full_line += " " + self.lines[start_index+1].strip()
            end_index = start_index + 1
        else:
            end_index = start_index

        if "::=" not in full_line:
            return None, end_index

        name_part, def_part = [p.strip() for p in full_line.split("::=", 1)]
        item = None

        # --- IMPROVED LOGIC: Check for the most specific types first ---

        # Handles all SEQUENCE OF definitions, even with constraints in the middle.
        if "SEQUENCE" in full_line and "OF" in full_line:
            name = name_part.split("SEQUENCE")[0].strip()
            item = ASN1Definition(name, "LIST")
            match = re.search(r'OF\s+([\w-]+)', def_part)
            if match: item.list_item_type = match.group(1).replace('}', '') # Clean up potential trailing braces
            constraints = self._parse_constraints(def_part)
            item.min_val, item.max_val = constraints['min_val'], constraints['max_val']
        
        # This handles all INTEGER definitions
        elif "INTEGER" in full_line:
            name = name_part.split("INTEGER")[0].strip()
            item = ASN1Definition(name, "INTEGER")
            if def_part.isdigit():
                item.min_val = item.max_val = def_part
            else:
                constraints = self._parse_constraints(def_part)
                item.min_val, item.max_val = constraints['min_val'], constraints['max_val']
                if item.min_val and not item.max_val:
                    item.max_val = item.min_val

        # This handles all BIT STRING definitions
        elif "BIT STRING" in full_line:
            name = name_part.split("BIT STRING")[0].strip()
            item = ASN1Definition(name, "BIT STRING")
            constraints = self._parse_constraints(def_part)
            item.min_val, item.max_val = constraints['min_val'], constraints['max_val']

        # --- NEW: Handle OCTET STRING and other common string types ---
        elif "OCTET STRING" in full_line:
            name = name_part.split("OCTET STRING")[0].strip()
            item = ASN1Definition(name, "OCTET STRING")
            constraints = self._parse_constraints(def_part)
            item.min_val, item.max_val = constraints['min_val'], constraints['max_val']
        
        elif any(s in full_line for s in ["PrintableString", "VisibleString", "UTF8String"]):
            # Find which string type it is
            str_type = "PrintableString"
            if "VisibleString" in full_line: str_type = "VisibleString"
            if "UTF8String" in full_line: str_type = "UTF8String"
            
            name = name_part.split(str_type)[0].strip()
            item = ASN1Definition(name, str_type) # Store the specific type
            constraints = self._parse_constraints(def_part)
            item.min_val, item.max_val = constraints['min_val'], constraints['max_val']
        # --- END NEW ---

        # This handles all ENUMERATED definitions
        elif "ENUMERATED" in full_line:
            name = name_part.split("ENUMERATED")[0].strip()
            item = ASN1Definition(name, "ENUMERATED")
            match = re.search(r'\{(.*)\}', def_part)
            if match:
                item.enum_values = [v.strip() for v in match.group(1).replace("...", "").split(',') if v.strip()]

        elif "NULL" in full_line:
            name = name_part.split("NULL")[0].strip()
            item = ASN1Definition(name, "NULL")
        
        # This handles ProcedureCode constants
        elif "ProcedureCode" in full_line:
             name = name_part.split("ProcedureCode")[0].strip()
             item = ASN1Definition(name, "ProcedureCode")
             value = def_part.split()[-1]
             if value.isdigit():
                 item.min_val = item.max_val = value

        # This handles ProtocolIE-ID constants
        elif "ProtocolIE-ID" in full_line:
             name = name_part.split("ProtocolIE-ID")[0].strip()
             item = ASN1Definition(name, "ProtocolIE-ID")
             value = def_part.split()[-1]
             if value.isdigit():
                 item.min_val = item.max_val = value
        
        if item:
            logger.debug(f"Parsed line definition: {item.name} (Type: {item.type})")
            if item.min_val and item.min_val == item.max_val and item.min_val.isdigit():
                item.is_constant = True
        else:
            logger.warning(f"Skipped unhandled line definition: {full_line}")
            
        return item, end_index

    def resolve_type_references(self):
        """
        Post-parsing step to link custom types within SEQUENCE and CHOICE definitions.

        This method iterates through all definitions. For SEQUENCEs and CHOICEs,
        it inspects each member's type. If that type corresponds to another
        definition we've parsed, it enriches the member's data with details
        from the referenced definition (like its base type and constraints).

        Returns:
            Tuple[int, set]: A tuple containing:
                - The total count of successfully resolved type references.
                - A set of type names that were referenced but could not be found.
        """
        resolved_count = 0
        unresolved_types = set()
        
        # A set of known, fundamental ASN.1 types that don't need a custom definition.
        # We also include types we parse directly, like BIT STRING.
        known_base_types = {
            "INTEGER", "BIT STRING", "OCTET STRING", "ENUMERATED", "BOOLEAN",
            "NULL", "SEQUENCE", "CHOICE", "UTF8String", "PrintableString",
            "VisibleString", "IA5String", "NumericString"
        }

        logger.debug("--- Starting Type Resolution ---")

        # Sort for deterministic logging order
        for def_name, def_body in sorted(self.definitions.items()):
            if def_body.type in ["SEQUENCE", "CHOICE"]:
                logger.debug(f"Analyzing members of '{def_name}' ({def_body.type})")
                for member in def_body.ies:
                    member_name = member.get("ie")
                    member_type_name = member.get("type")

                    # Check if the member's type is a custom type we have parsed
                    if member_type_name in self.definitions:
                        resolved_count += 1
                        ref_def = self.definitions[member_type_name]
                        
                        logger.debug(f"  Resolving member '{member_name}' -> type '{member_type_name}'")
                        
                        # Add the base type (e.g., INTEGER, BIT STRING) for clarity
                        member['base_type'] = ref_def.type
                        logger.debug(f"    -> Base type is '{ref_def.type}'")

                        # Inherit constraints ONLY if not defined locally on the member
                        if member.get('min') is None and ref_def.min_val is not None:
                            member['min'] = ref_def.min_val
                            logger.debug(f"    -> Inherited min: {ref_def.min_val}")
                        
                        if member.get('max') is None and ref_def.max_val is not None:
                            member['max'] = ref_def.max_val
                            logger.debug(f"    -> Inherited max: {ref_def.max_val}")

                    # If not a custom type, check if it's a known base type
                    elif member_type_name in known_base_types:
                        logger.debug(f"  Member '{member_name}' is a known base type: '{member_type_name}'")
                    
                    # Otherwise, it's an unresolved type
                    else:
                        logger.debug(f"  Member '{member_name}' has UNRESOLVED type: '{member_type_name}'")
                        unresolved_types.add(member_type_name)

        logger.debug("--- Finished Type Resolution ---")
        return resolved_count, unresolved_types