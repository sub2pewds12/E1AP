from typing import List
from common_types import ListDefinition, SequenceDefinition, ChoiceDefinition


def render_enum_struct(
    go_name: str, enum_values: List[str], pascal_case_converter: callable
) -> str:
    """
    Generates the Go code for an ENUMERATED type definition and its constants.
    """
    const_lines = []
    for i, val in enumerate(enum_values):
        member_name = pascal_case_converter(val).replace("_", "")
        const_name = f"{go_name}{member_name}"
        const_lines.append(f"\t{const_name} int64 = {i}")

    const_block = "const (\n" + "\n".join(const_lines) + "\n)"

    struct_block = f"""\
// {go_name} is a generated ENUMERATED type.
type {go_name} struct {{
	Value int64
}}"""

    return f"{struct_block}\n\n{const_block}\n"


def render_sequence_struct(
    go_name: str,
    item: SequenceDefinition,
    pascal_case_converter: callable,
    go_type_resolver: callable,
    go_value_formatter: callable,
    extension_sets: dict,
) -> str:
    """
    Generates ONLY the Go code for a SEQUENCE struct definition.
    """
    field_lines = []
    for member in item.ies:
        member_go_name = pascal_case_converter(member.ie)
        base_go_type, concrete_def = go_type_resolver(member.type)

        final_go_type = ""
        if member.type == "ProtocolExtensionContainer" and hasattr(member, 'extension_set_name'):
             # Always use the specialized type, even if the extension set is empty,
             # because we now generate specialized extension structs for all such fields.
             final_go_type = pascal_case_converter(item.name) + "Extensions"
        else:
             final_go_type = pascal_case_converter(member.type)
        if final_go_type == "ANY":
            final_go_type = "aper.IE"
        
        is_list = isinstance(concrete_def, ListDefinition)
        is_optional = member.presence in ["optional", "conditional"]
        pointer = "*" if is_optional else ""

        field_lines.append(f"\t{member_go_name}\t{pointer}{final_go_type}")
    
    struct_body = "\n".join(field_lines)
    struct_block = f"// {go_name} is a generated SEQUENCE type.\ntype {go_name} struct {{\n{struct_body}\n}}"

    return struct_block


def render_choice_struct(
    go_name: str,
    item: ChoiceDefinition,
    pascal_case_converter: callable,
    go_type_resolver: callable,
) -> str:
    """
    Generates the Go code for a CHOICE type definition and its constants.
    """
    const_lines = [f"\t{go_name}PresentNothing uint64 = iota"]
    for member in item.ies:
        member_go_name = pascal_case_converter(member.ie)
        const_lines.append(f"\t{go_name}Present{member_go_name}")

    const_block = "const (\n" + "\n".join(const_lines) + "\n)"

    field_lines = ["\tChoice\tuint64 `json:\"-\"`"] 
    for member in item.ies:
        member_go_name = pascal_case_converter(member.ie)
        base_go_type, concrete_def = go_type_resolver(member.type)
        final_go_type = base_go_type

        field_lines.append(f"\t{member_go_name}\t*{final_go_type}")

    struct_body = "\n".join(field_lines)
    struct_block = f"// {go_name} is a generated CHOICE type.\ntype {go_name} struct {{\n{struct_body}\n}}"

    return f"{struct_block}\n\n{const_block}\n"

def render_extension_struct_only(
    go_name: str,
    extension_set: list,
    pascal_case_converter: callable,
) -> str:
    """
    Generates ONLY the dedicated, type-safe Go struct for a set of extensions.
    """
    field_lines = []
    for ext in extension_set:
        ext_field_name = pascal_case_converter(ext['id'])
        ext_type_name = pascal_case_converter(ext['type'])
        field_lines.append(f"\t{ext_field_name}\t*{ext_type_name}")

    struct_body = "\n".join(field_lines)
    struct_block = f"// {go_name} is a generated type-safe wrapper for extensions.\ntype {go_name} struct {{\n{struct_body}\n}}"

    return struct_block

def render_list_struct(go_name: str, item: ListDefinition, pascal_case_converter: callable) -> str:
    """
    Generates the Go struct for a LIST (SEQUENCE OF) type.
    e.g., type MyList struct { Value []MyItem }
    """
    of_type_go_name = pascal_case_converter(item.of_type)
    return f"// {go_name} is a generated LIST type.\ntype {go_name} struct {{\n\tValue []{of_type_go_name}\n}}"
