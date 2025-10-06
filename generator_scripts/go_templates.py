from typing import List
from common_types import ListDefinition, SequenceDefinition, ChoiceDefinition


def render_enum(
    go_name: str, enum_values: List[str], pascal_case_converter: callable
) -> str:
    """
    Generates the Go code for an ENUMERATED type.
    This function acts as our pure Python "template".
    """

    const_lines = []
    for i, val in enumerate(enum_values):

        member_name = pascal_case_converter(val).replace("_", "")
        const_name = f"{go_name}{member_name}"
        const_lines.append(f"\t{const_name} aper.Enumerated = {i}")

    const_block = "const (\n" + "\n".join(const_lines) + "\n)"

    struct_block = f"""\
// {go_name} is a generated ENUMERATED type.
type {go_name} struct {{
	Value aper.Enumerated
}}"""

    return f"{struct_block}\n\n{const_block}\n"


def render_sequence(
    go_name: str,
    item: SequenceDefinition,
    pascal_case_converter: callable,
    go_type_resolver: callable,
    go_value_formatter: callable,
) -> str:
    """
    Generates the Go code for a SEQUENCE type, including placeholder
    Encode/Decode methods for PDUs.
    """

    field_lines = []
    for member in item.ies:
        member_go_name = pascal_case_converter(member.ie)
        base_go_type, concrete_def = go_type_resolver(member.type)

        final_go_type = base_go_type
        is_list = False
        if concrete_def and isinstance(concrete_def, ListDefinition):
            of_type_go_name = pascal_case_converter(concrete_def.of_type)
            final_go_type = f"[]{of_type_go_name}"
            is_list = True

        is_optional = member.presence in ["optional", "conditional"]
        pointer = "*" if is_optional and not is_list else ""

        tag_parts = []
        if (
            concrete_def
            and hasattr(concrete_def, "min_val")
            and concrete_def.min_val is not None
        ):
            tag_parts.append(f"lb:{go_value_formatter(concrete_def.min_val)}")
        if (
            concrete_def
            and hasattr(concrete_def, "max_val")
            and concrete_def.max_val is not None
        ):
            tag_parts.append(f"ub:{go_value_formatter(concrete_def.max_val)}")

        tag_parts.append(member.presence)
        if member.criticality:
            tag_parts.append(member.criticality)

        if item.is_extensible:
            tag_parts.append("ext")

        tag_string = f'`aper:"{",".join(tag_parts)}"`'
        field_lines.append(
            f"\t{member_go_name}\t{pointer}{final_go_type}\t{tag_string}"
        )

    struct_body = "\n".join(field_lines)
    struct_block = f"// {go_name} is a generated SEQUENCE type.\ntype {go_name} struct {{\n{struct_body}\n}}"

    placeholder_methods = ""

    is_pdu = item.ies and item.ies[0].id is not None
    if is_pdu:
        placeholder_methods = f"""
// Encode implements the aper.AperMarshaller interface.
func (s *{go_name}) Encode(w io.Writer) error {{
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for {go_name}")
}}

// Decode implements the aper.AperUnmarshaller interface.
func (s *{go_name}) Decode(r io.Reader) error {{
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for {go_name}")
}}
"""

    return f"{struct_block}\n{placeholder_methods}"


def render_choice(
    go_name: str,
    item: ChoiceDefinition,
    pascal_case_converter: callable,
    go_type_resolver: callable,
) -> str:
    """
    Generates the Go code for a CHOICE type.
    """

    const_lines = [f"\t{go_name}PresentNothing uint64 = iota"]
    for member in item.ies:
        member_go_name = pascal_case_converter(member.ie)
        const_lines.append(f"\t{go_name}Present{member_go_name}")

    const_block = "const (\n" + "\n".join(const_lines) + "\n)"

    field_lines = ["\tChoice\tuint64"]
    for member in item.ies:
        member_go_name = pascal_case_converter(member.ie)
        base_go_type, concrete_def = go_type_resolver(member.type)

        final_go_type = base_go_type
        is_list = False
        if concrete_def and isinstance(concrete_def, ListDefinition):
            of_type_go_name = pascal_case_converter(concrete_def.of_type)
            final_go_type = f"[]{of_type_go_name}"
            is_list = True

        pointer = "" if is_list else "*"
        field_lines.append(f"\t{member_go_name}\t{pointer}{final_go_type}")

    struct_body = "\n".join(field_lines)
    struct_block = f"// {go_name} is a generated CHOICE type.\ntype {go_name} struct {{\n{struct_body}\n}}"

    return f"{struct_block}\n\n{const_block}\n"
