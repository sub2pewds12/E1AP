import re
import logging
from typing import Dict, List, Optional, Any, Tuple
from enum import Enum, auto
from common_types import (
    ASN1Definition,
    ASN1Procedure,
    InformationElement,
    IntegerDefinition,
    ConstantDefinition,
    EnumDefinition,
    StringDefinition,
    SequenceDefinition,
    ChoiceDefinition,
    ListDefinition,
    AliasDefinition,
    BuiltinDefinition,
    ASN1Class,
    ASN1ClassField,
)

logger = logging.getLogger(__name__)


class ParserState(Enum):
    """Defines the possible states for our parser."""

    SEARCHING = auto()
    IN_MULTILINE_DEF = auto()


class ASN1Parser:
    def __init__(self, lines: List[Tuple[str, str, int]]):
        self.lines = lines
        self.definitions: Dict[str, ASN1Definition] = {}
        self.ie_sets: Dict[str, List[Dict[str, str]]] = {}
        self.type_metadata: Dict[str, Dict[str, str]] = {}
        self.procedures: Dict[str, ASN1Procedure] = {}
        self.message_to_procedure_map: Dict[str, str] = {}
        self._prepopulate_builtin_types()
        self.state = ParserState.SEARCHING
        self.current_def_lines = []
        self.current_def_start_context = None
        self.open_braces = 0
        self.failures: List[Dict[str, Any]] = []
        self.classes: Dict[str, ASN1Class] = {}

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
            "OBJECT IDENTIFIER": "OBJECT IDENTIFIER",
            "BOOLEAN": "BOOLEAN",
        }
        for name in builtins:
            item = BuiltinDefinition(name=name, source_file="builtin")
            self.definitions[name] = item

    def _reset_state(self):
        """Resets the state machine variables for the next definition."""
        self.current_def_lines = []
        self.current_def_start_context = None
        self.open_braces = 0
        self.state = ParserState.SEARCHING

    def _parse_ie_set(self, name: str, def_part: str):
        """
        Parses a ...-PROTOCOL-IES block into a rich, structured list of metadata.
        It can now handle simple inline definitions for the TYPE field.
        """

        ie_pattern = re.compile(
            r"\{\s*ID\s+(?P<id>[\w-]+)\s+CRITICALITY\s+(?P<crit>[\w-]+)\s+(?:TYPE|EXTENSION)\s+(?P<type>.+?)\s+PRESENCE\s+(?P<pres>[\w-]+)\s*\}"
        )
        matches = ie_pattern.finditer(def_part)
        ie_list_for_set = []

        for match in matches:
            ie_id = match.group("id")

            raw_type_str = match.group("type").strip().rstrip(",")
            type_name_to_add = ""

            is_simple_inline = any(
                raw_type_str.strip().startswith(k)
                for k in ["INTEGER", "BIT STRING", "OCTET STRING", "ENUMERATED"]
            )

            if is_simple_inline:
                synthetic_type_name = f"{name}-{ie_id}"

                inline_def = self._parse_simple_definition(
                    synthetic_type_name,
                    raw_type_str,
                    "inline_in_ie_set",
                    -1,
                    raw_type_str,
                )
                if inline_def:
                    inline_def.is_synthetic = True
                    if inline_def.name not in self.definitions:
                        self.definitions[inline_def.name] = inline_def

                    type_name_to_add = inline_def.name
                else:
                    logger.warning(
                        f"Failed to parse inline TYPE '{raw_type_str}' in IE Set '{name}'"
                    )
                    type_name_to_add = raw_type_str
            else:

                type_name_to_add = raw_type_str

            ie_data = {
                "id": ie_id,
                "type": type_name_to_add,
                "crit": match.group("crit").lower(),
                "pres": match.group("pres").lower(),
            }
            ie_list_for_set.append(ie_data)
            self.type_metadata[type_name_to_add] = ie_data

        if ie_list_for_set:
            self.ie_sets[name] = ie_list_for_set

    def _parse_constraints(self, line_part: str) -> Dict[str, Any]:
        constraints = {"min_val": None, "max_val": None}

        match = re.search(r"\([^)]+\)", line_part)
        if not match:
            return constraints
        content = match.group(0)

        size_range_match = re.search(r"SIZE\s*\(([\w\d-]+)\.\.([\w\d-]+)\)", content)
        if size_range_match:
            constraints["min_val"] = size_range_match.group(1)
            constraints["max_val"] = size_range_match.group(2)
            return constraints

        size_fixed_match = re.search(r"SIZE\s*\(([\w\d-]+)\)", content)
        if size_fixed_match:
            val = size_fixed_match.group(1)
            constraints["min_val"] = val
            constraints["max_val"] = val
            return constraints

        range_match = re.search(r"\(([\w\d-]+)\.\.([\w\d-]+)", content)
        if range_match:
            min_val = range_match.group(1)
            max_val = range_match.group(2)
            if not min_val.isalpha() or min_val in self.definitions:
                constraints["min_val"] = min_val
            if not max_val.isalpha() or max_val in self.definitions:
                constraints["max_val"] = max_val
            return constraints

        fixed_val_match = re.search(r"\(([\w\d-]+)\)", content)
        if fixed_val_match:
            val = fixed_val_match.group(1)
            constraints["min_val"] = val
            constraints["max_val"] = val
            return constraints

        return constraints

    def _parse_presence(self, line_part: str) -> str:
        if "OPTIONAL" in line_part:
            return "optional"
        if "CONDITIONAL" in line_part:
            return "conditional"
        return "mandatory"

    def _extract_full_definition(self, start_index: int) -> Tuple[str, int]:
        full_def_lines = []
        current_index = start_index
        open_braces = 0
        has_started = False

        while current_index < len(self.lines):
            line = self.lines[current_index][0]
            cleaned_line = line.split("--", 1)[0].strip()

            if not cleaned_line:
                current_index += 1
                continue

            if has_started and "::=" in cleaned_line:
                break

            full_def_lines.append(cleaned_line)
            open_braces += cleaned_line.count("{") - cleaned_line.count("}")
            
            # --- THIS IS THE ONE-LINE FIX ---
            # We must also count parentheses to correctly find the end of the definition.
            open_braces += cleaned_line.count("(") - cleaned_line.count(")")
            # --- END OF FIX ---

            has_started = True

            # Your existing break logic now works correctly because open_braces is accurate.
            if (
                open_braces == 0
                and "{" not in "".join(full_def_lines)
                and not "".join(full_def_lines).rstrip().endswith("OF")
            ):
                current_index += 1
                break
            if open_braces <= 0 and "{" in "".join(full_def_lines):
                current_index += 1
                break

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
        if "{" in name or "}" in name:
            return None
        def_part_stripped = def_part.strip()
        name_parts = name.strip().split()

        complex_keywords = ["SEQUENCE OF", "SEQUENCE {", "CHOICE {", "SEQUENCE ("]
        if any(keyword in def_part for keyword in complex_keywords):
            return None
        
        if (
            len(name_parts) >= 2 or def_part_stripped.isdigit()
        ) and "{" not in def_part:
            const_type = name_parts[-1] if len(name_parts) >= 2 else "INTEGER"
            const_name = " ".join(name_parts[:-1]) if len(name_parts) >= 2 else name

            if "AUTOMATIC TAGS" in full_text:
                return None

            item = ConstantDefinition(
                const_type, const_name, source_file, source_line, full_text
            )
            item.min_val, item.max_val = def_part_stripped, def_part_stripped
            return item

        if def_part_stripped.startswith("INTEGER"):

            item = IntegerDefinition(name, source_file, source_line, full_text)
            constraints = self._parse_constraints(def_part)
            item.min_val, item.max_val = constraints["min_val"], constraints["max_val"]
            return item

        elif def_part_stripped.startswith("ENUMERATED"):
            item = EnumDefinition(name, source_file, source_line, full_text)
            match = re.search(r"ENUMERATED\s*\{([^}]+)\}", def_part, re.DOTALL)
            if match:
                content = match.group(1).replace("...", "").strip()
                item.enum_values = [v.strip() for v in content.split(",") if v.strip()]
            else:
                item.enum_values = ["present"]
            return item

        elif "BIT STRING" in def_part_stripped or "OCTET STRING" in def_part_stripped:
            base_type = (
                "BIT STRING" if "BIT STRING" in def_part_stripped else "OCTET STRING"
            )
            item = StringDefinition(
                base_type, name, source_file, source_line, full_text
            )
            constraints = self._parse_constraints(def_part)
            item.min_val, item.max_val = constraints["min_val"], constraints["max_val"]
            return item

        value_parts = def_part_stripped.split()
        if value_parts:

            base_type = value_parts[0].split("(")[0].strip()
            if not base_type:
                return None
            if base_type in ["SEQUENCE", "CHOICE"]:
                return None

            is_container_alias = (
                "ProtocolIE-Field" in base_type or "ProtocolExtensionField" in base_type
            )
            if is_container_alias:
                item = AliasDefinition(
                    base_type, name, source_file, source_line, full_text
                )
                return item

            is_complex_list = (
                "SEQUENCE" in def_part_stripped and " OF " in def_part_stripped
            )
            is_simple_alias = (
                base_type[0].isupper()
                and "{" not in def_part_stripped
                and "{" not in name
                and not is_complex_list
            )
            if is_simple_alias:
                item = AliasDefinition(
                    base_type, name, source_file, source_line, full_text
                )
                constraints = self._parse_constraints(def_part)
                item.min_val, item.max_val = (
                    constraints["min_val"],
                    constraints["max_val"],
                )
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
        item = (
            SequenceDefinition(name, source_file, source_line, full_text)
            if def_type == "SEQUENCE"
            else ChoiceDefinition(name, source_file, source_line, full_text)
        )
        item.is_extensible = "..." in def_part

        container_match = re.search(
            r"ProtocolIE-Container\s*\{\s*\{\s*([\w-]+)\s*\}\s*\}", def_part
        )
        if container_match:
            ie_set_name = container_match.group(1)
            if ie_set_name in self.ie_sets:
                for ie_data in self.ie_sets[ie_set_name]:
                    ie_name_to_add = ie_data["id"].replace("id-", "")
                    if not ie_name_to_add:
                        continue

                    new_ie = InformationElement(
                        id=ie_data["id"],
                        ie=ie_name_to_add,
                        type=ie_data["type"],
                        presence=ie_data["pres"],
                        criticality=ie_data["crit"],
                    )
                    item.ies.append(new_ie)
                return item
            else:
                logger.warning(
                    f"Could not find pre-parsed IE Set '{ie_set_name}' for container '{name}'."
                )

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
            if char in ["{", "("]:
                nesting_level += 1
            elif char in ["}", ")"]:
                nesting_level -= 1

            if char == "," and nesting_level == 0:
                member_lines.append(current_line)
                current_line = ""
            else:
                current_line += char
        member_lines.append(current_line)

        for member_line in member_lines:
            clean_line = member_line.strip()
            if not clean_line or clean_line.startswith("..."):
                continue

            parts = clean_line.split(None, 1)
            if len(parts) < 2:
                continue

            member_name, raw_type_str = parts[0], parts[1]
            ie_name_to_add = member_name
            type_name_to_add = None

            is_inline_list = " OF " in raw_type_str

            if is_inline_list:
                synthetic_type_name = f"{name}-{member_name}"
                inline_def = self._parse_sequence_of(
                    synthetic_type_name,
                    raw_type_str,
                    source_file,
                    source_line,
                    raw_type_str,
                )
                if inline_def:
                    inline_def.is_synthetic = True
                    if inline_def.name not in self.definitions:
                        self.definitions[inline_def.name] = inline_def
                    type_name_to_add = inline_def.name

            elif re.search(r"([\w-]+)\s*\.&([\w-]+)", raw_type_str):
                class_match = re.search(r"([\w-]+)\s*\.&([\w-]+)", raw_type_str)
                class_name = class_match.group(1)
                field_name = class_match.group(2).lower()
                class_def = self.classes.get(class_name)
                if class_def and class_def.fields.get(field_name):
                    type_name_to_add = class_def.fields.get(field_name).type
                else:
                    type_name_to_add = "ANY"

            elif any(k in raw_type_str for k in ["SEQUENCE {", "CHOICE {"]):
                synthetic_type_name = f"{name}-{member_name}"
                inline_def = self._parse_block_definition(
                    synthetic_type_name,
                    raw_type_str,
                    source_file,
                    source_line,
                    raw_type_str,
                )
                if inline_def:
                    if inline_def.name not in self.definitions:
                        self.definitions[inline_def.name] = inline_def
                    type_name_to_add = inline_def.name

            elif any(
                raw_type_str.strip().startswith(k)
                for k in ["INTEGER", "ENUMERATED", "BIT STRING", "OCTET STRING"]
            ):
                synthetic_type_name = f"{name}-{member_name}"
                inline_def = self._parse_simple_definition(
                    synthetic_type_name,
                    raw_type_str,
                    source_file,
                    source_line,
                    raw_type_str,
                )
                if inline_def:
                    inline_def.is_synthetic = True
                    if inline_def.name not in self.definitions:
                        self.definitions[inline_def.name] = inline_def
                    type_name_to_add = inline_def.name

            else:

                match = re.match(r"^([\w-]+)", raw_type_str)
                if match:
                    type_name_to_add = match.group(1).strip()
                else:
                    logger.warning(
                        f"Could not extract type name for member '{member_name}' in '{name}'"
                    )
                    continue

            if ie_name_to_add and type_name_to_add:

                member_metadata = self.type_metadata.get(type_name_to_add)
                presence = (
                    member_metadata.get("pres")
                    if member_metadata
                    else self._parse_presence(raw_type_str)
                )
                crit = member_metadata.get("crit") if member_metadata else None
                ie_id = member_metadata.get("id") if member_metadata else None
                ie = InformationElement(
                    ie=ie_name_to_add,
                    type=type_name_to_add,
                    presence=presence,
                    criticality=crit,
                    id=ie_id,
                )
                item.ies.append(ie)

        return item

    def _parse_procedure(self, name: str, def_part: str):
        """Parses an E1AP-ELEMENTARY-PROCEDURE and builds the message->procedure map."""
        proc = ASN1Procedure(name)

        init_msg_match = re.search(r"INITIATING MESSAGE\s+([\w-]+)", def_part)
        succ_out_match = re.search(r"SUCCESSFUL OUTCOME\s+([\w-]+)", def_part)
        unsucc_out_match = re.search(r"UNSUCCESSFUL OUTCOME\s+([\w-]+)", def_part)
        proc_code_match = re.search(r"PROCEDURE CODE\s+([\w-]+)", def_part)

        if init_msg_match:
            proc.initiating_message = init_msg_match.group(1)

            self.message_to_procedure_map[proc.initiating_message] = "InitiatingMessage"
        if succ_out_match:
            proc.successful_outcome = succ_out_match.group(1)
            self.message_to_procedure_map[proc.successful_outcome] = "SuccessfulOutcome"
        if unsucc_out_match:
            proc.unsuccessful_outcome = unsucc_out_match.group(1)
            self.message_to_procedure_map[proc.unsuccessful_outcome] = (
                "UnsuccessfulOutcome"
            )
        if proc_code_match:
            proc.procedure_code = proc_code_match.group(1)

        self.procedures[proc.name] = proc

    def _parse_sequence_of(
        self, name_part, def_part, source_file, source_line, full_def_str
    ):
        """
        Parses a SEQUENCE OF definition.
        """
        item = ListDefinition(name_part, source_file, source_line, full_def_str)
        constraints = self._parse_constraints(def_part)
        item.min_val, item.max_val = constraints["min_val"], constraints["max_val"]

        try:

            of_part = re.split(r"\sOF\s", def_part, 1)[1].strip()
        except IndexError:
            logger.warning(f"Could not determine the 'OF' type for LIST: {name_part}")
            return item

        container_match = re.search(r"^([\w-]+)\s*\{", of_part)
        if container_match and container_match.group(1) in [
            "ProtocolIE-SingleContainer",
            "ProtocolIE-Container",
        ]:

            item.of_type = container_match.group(1)

            synthetic_item_name = (
                name_part.replace("-List", "Item")
                .replace("ResAck", "ResAckItem")
                .replace("Res", "ResItem")
            )

            if synthetic_item_name not in self.definitions:
                self.definitions[synthetic_item_name] = SequenceDefinition(
                    synthetic_item_name, source_file, source_line
                )

            item.of_type = synthetic_item_name

        elif of_part.startswith("SEQUENCE {"):

            synthetic_name = name_part.replace("-List", "-Item")
            if synthetic_name == name_part:
                synthetic_name = name_part + "Item"

            inline_def = self._parse_block_definition(
                synthetic_name, of_part, source_file, source_line, of_part
            )
            if inline_def:
                inline_def.is_synthetic = True
                if inline_def.name not in self.definitions:
                    self.definitions[inline_def.name] = inline_def
                item.of_type = inline_def.name

        else:
            match = re.search(r"^([\w-]+)", of_part)
            if match:
                item.of_type = match.group(1)

        return item

    def _handle_searching_state(self, line: str, source_file: str, source_line: int):
        """Handles a line when we are looking for a new definition."""
        if "::=" not in line:
            return

        self.current_def_lines = [line]
        self.current_def_start_context = (source_file, source_line)
        self.open_braces = line.count("{") - line.count("}")

        if self.open_braces == 0 or "{" not in line:
            self._process_complete_definition()
        else:

            self.state = ParserState.IN_MULTILINE_DEF

    def _handle_multiline_def_state(self, line: str):
        """Handles a line when we are inside a multi-line definition block."""
        if "::=" in line:
            # This part handles cases where a new definition starts before the old one ended.
            self._process_complete_definition(is_malformed=True)
            # We then need to re-process the current line in the SEARCHING state.
            self._handle_searching_state(line, "unknown", -1)
            return

        self.current_def_lines.append(line)
        
        self.open_braces += line.count("{") - line.count("}")
        
        # --- THIS IS THE ONE-LINE FIX ---
        # We must also count parentheses to correctly find the end of the definition.
        self.open_braces += line.count("(") - line.count(")")
        # --- END OF FIX ---

        if self.open_braces <= 0:
            self._process_complete_definition()
            # This ensures the state is clean for the next definition.
            self._reset_state()


    def _process_complete_definition(self, is_malformed=False):
        """
        Takes the collected lines for a definition, joins them,
        and calls the appropriate parsing logic using a robust, ordered dispatcher.
        """
        full_def_str = " ".join(self.current_def_lines)
        source_file, source_line = self.current_def_start_context

        if "::=" not in full_def_str:
            self._reset_state()
            return

        name_part, def_part = [p.strip() for p in full_def_str.split("::=", 1)]

        if "{" in name_part and "}" in name_part:
            self.failures.append({
                "name": name_part, "text": full_def_str,
                "file": source_file, "line": source_line,
                "reason": "Skipping known abstract/parameterized definition"
            })
            self._reset_state()
            return

        item = None

        if "E1AP-PROTOCOL-IES" in name_part:
            sanitized_name = name_part.replace("E1AP-PROTOCOL-IES", "").strip()
            self._parse_ie_set(sanitized_name, def_part)
            self._reset_state()
            return

        elif "E1AP-ELEMENTARY-PROCEDURE" in name_part:
            proc_match = re.match(r"^([a-z][\w-]*)\s+E1AP-ELEMENTARY-PROCEDURE", name_part)
            if proc_match:
                self._parse_procedure(proc_match.group(1), def_part)
            self._reset_state()
            return

        elif def_part.strip().startswith("CLASS"):
            self._parse_class_definition(name_part, def_part, source_file, source_line, full_def_str)
            self._reset_state()
            return

        elif "SEQUENCE" in def_part and " OF " in def_part:
            item = self._parse_sequence_of(
                name_part, def_part, source_file, source_line, full_def_str
            )
        elif ("SEQUENCE" in def_part or "CHOICE" in def_part) and "{" in def_part:
            item = self._parse_block_definition(name_part, def_part, source_file, source_line, full_def_str)
        else:
            item = self._parse_simple_definition(name_part, def_part, source_file, source_line, full_def_str)

        if item:
            if item.name not in self.definitions:
                self.definitions[item.name] = item
        else:
            self.failures.append({"name": name_part, "text": full_def_str, "file": source_file, "line": source_line})
        self._reset_state()

    def _parse_class_definition(
        self,
        name: str,
        def_part: str,
        source_file: str,
        source_line: int,
        full_text: str,
    ):
        """Parses a CLASS definition block using a robust split-and-parse strategy."""
        class_def = ASN1Class(name, source_file, source_line, full_text)

        body_match = re.search(r"\{(.*)\}", def_part, re.DOTALL)
        if not body_match:
            logger.warning(f"Could not find body for CLASS '{name}'")
            return

        body = body_match.group(1)

        field_declarations = body.split(",")

        for decl in field_declarations:
            clean_decl = decl.strip()
            if not clean_decl:
                continue

            field_match = re.match(r"&([\w-]+)\s*(.*)", clean_decl)
            if not field_match:
                continue

            field_name = field_match.group(1).lower()
            rest_of_line = field_match.group(2).strip()

            field_type = "ANY"
            if rest_of_line and not rest_of_line.upper() == "OPTIONAL":

                field_type = rest_of_line.split()[0]

            class_def.fields[field_name] = ASN1ClassField(
                name=field_name, type=field_type
            )

        self.classes[name] = class_def
        logger.info(
            f"Successfully parsed CLASS '{name}' with {len(class_def.fields)} fields."
        )

    def parse(self) -> Tuple[Dict[str, ASN1Definition], List[Dict[str, Any]]]:
        #failures = []

        logger.info("PASS 1: Scanning for and parsing all IE Set definitions...")
        for i in range(len(self.lines)):
            line_text, _, _ = self.lines[i]
            line = line_text.strip()

            if "::=" in line and "E1AP-PROTOCOL-IES" in line:
                full_def_str, _ = self._extract_full_definition(i)
                name_part, def_part = [p.strip() for p in full_def_str.split("::=", 1)]

                if "GTPTLAs" in name_part:
                    print("\n\n" + "="*20 + " DEBUGGING GTPTLAs " + "="*20)
                    print(f"FOUND DEFINITION: {name_part}")
                    print(f"  >>> Full Text: {full_def_str}")
                    print("="*60 + "\n\n")

                if "E1AP-PROTOCOL-IES" in name_part:
                    ie_set_name = name_part.replace("E1AP-PROTOCOL-IES", "").strip()
                    self._parse_ie_set(ie_set_name, def_part)
        logger.info(f"PASS 1 COMPLETE: Found and parsed {len(self.ie_sets)} IE Sets.")

        logger.info("PASS 1.5: Scanning for and parsing all CLASS definitions...")
        for i in range(len(self.lines)):
            line_text, source_file, source_line = self.lines[i]
            line = line_text.strip()
            if "::=" in line and "CLASS" in line.split("::=", 1)[1]:
                full_def_str, _ = self._extract_full_definition(i)
                name_part, def_part = [p.strip() for p in full_def_str.split("::=", 1)]
                self._parse_class_definition(
                    name_part, def_part, source_file, source_line, full_def_str
                )
        logger.info(f"PASS 1.5 COMPLETE: Found and parsed {len(self.classes)} CLASSes.")

        logger.info("PASS 2: Parsing all other ASN.1 definitions...")
        i = 0

        patched_definitions_to_skip = {
            "UE-associatedLogicalE1-ConnectionListRes",
            "UE-associatedLogicalE1-ConnectionListResAck",
        }
        lines_to_skip = set()
        for i in range(len(self.lines)):
            line_text, _, _ = self.lines[i]
            if "::=" in line_text:
                name_part = line_text.split("::=", 1)[0].strip()
                if name_part in patched_definitions_to_skip:
                    _, end_index = self._extract_full_definition(i)
                    # Add all line numbers belonging to this definition to our skip set.
                    for skip_line_num in range(i, end_index + 1):
                        lines_to_skip.add(skip_line_num)

        # This loop now correctly drives the state machine for all definitions.
        for i, (line_text, source_file, source_line) in enumerate(self.lines):
            # Your skip logic is now here
            if i in lines_to_skip:
                continue

            cleaned_line = line_text.split("--", 1)[0].strip()
            if not cleaned_line:
                continue
            
            # Your guards against malformed lines are preserved here
            if "::=" in cleaned_line:
                name_part = cleaned_line.split("::=", 1)[0].strip()
                if "{" in name_part and "}" in name_part:
                     self.failures.append(
                        {
                            "name": name_part,
                            "text": cleaned_line,
                            "file": source_file,
                            "line": source_line,
                        }
                    )
                     continue

            # This is the correct state machine driver
            if self.state == ParserState.SEARCHING:
                self._handle_searching_state(cleaned_line, source_file, source_line)
            elif self.state == ParserState.IN_MULTILINE_DEF:
                self._handle_multiline_def_state(cleaned_line)

        # Process any remaining definition at the end of the file.
        if self.current_def_lines:
            self._process_complete_definition()

        logger.info("PASS 2 COMPLETE.")
        failures = []

        if self.current_def_lines:
            self._process_complete_definition()

        logger.info("PASS 2 COMPLETE.")

        # The function now immediately returns the correct data.
        return (
            self.definitions,
            self.failures, # This list is no longer being wiped out
            self.procedures,
            self.message_to_procedure_map,
        )
