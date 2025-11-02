from common_types import (
    SequenceDefinition,
    ChoiceDefinition,
    EnumDefinition,
    IntegerDefinition,
    StringDefinition,
    ListDefinition,
    BuiltinDefinition,
)
from asn1_parser import ASN1Parser
from typing import Tuple, Set


def render_pdu_methods(
    go_name: str,
    item: SequenceDefinition,
    parser: ASN1Parser,
    message_to_procedure_map: dict,
    procedures: dict,
) -> Tuple[str, Set[str]]:
    """
    Generates the full set of methods for a Top-Level PDU (SEQUENCE type).
    This includes toIes, Encode, Decode, and the Decoder helper struct/method.
    """
    required_imports = {"io", "fmt", "bytes", parser.aper_import_path}

    to_ies_code = _generate_to_ies(go_name, item, parser)
    encode_code = _generate_pdu_encode(go_name, item, message_to_procedure_map, procedures, parser)
    decode_code, decode_imports = _generate_pdu_decode(go_name, item, parser)
    decoder_helper_code, helper_imports = _generate_decoder_helper(
        go_name, item, parser
    )

    required_imports.update(decode_imports)
    required_imports.update(helper_imports)

    full_method_block = f"""
{to_ies_code}

{encode_code}

{decode_code}

{decoder_helper_code}
"""
    return full_method_block, required_imports


def _generate_to_ies(go_name: str, item: SequenceDefinition, parser: ASN1Parser) -> str:
    """Generates the toIes() method body as a string."""

    pascal_case_converter = parser.pascal_case_converter
    go_type_resolver = parser.go_type_resolver

    body_parts = []

    for ie in item.ies:
        field_name = pascal_case_converter(ie.ie)
        _, asn1_def = go_type_resolver(ie.type)

        if not asn1_def:
            print(
                f"WARNING: Could not resolve type '{ie.type}' for IE '{ie.ie}' in PDU '{go_name}'. Skipping."
            )
            continue

        is_optional = ie.presence in ["optional", "conditional"]
        is_extensible = getattr(asn1_def, "is_extensible", False)

        value_code = ""
        pre_append_code = ""

        if isinstance(asn1_def, IntegerDefinition):
            value_accessor = f"msg.{field_name}"
            if is_optional:
                value_accessor = f"(*{value_accessor})"
            value_code = f"""&INTEGER{{
                c:     aper.Constraint{{Lb: {asn1_def.min_val or 0}, Ub: {asn1_def.max_val or 0}}},
                ext:   {str(is_extensible).lower()},
                Value: aper.Integer({value_accessor}),
            }}"""

        elif isinstance(asn1_def, EnumDefinition):
            value_accessor = f"msg.{field_name}"
            if is_optional:
                value_accessor = f"(*{value_accessor})"
            num_enums = len(asn1_def.enum_values)
            upper_bound = num_enums - 1 if num_enums > 0 else 0
            value_code = f"""&ENUMERATED{{
                c:     aper.Constraint{{Lb: 0, Ub: {upper_bound}}},
                ext:   {str(is_extensible).lower()},
                Value: {value_accessor}.Value,
            }}"""

        elif isinstance(asn1_def, StringDefinition) or (
            isinstance(asn1_def, BuiltinDefinition)
            and "STRING" in asn1_def.name.upper()
        ):
            value_accessor = f"msg.{field_name}"
            if is_optional:
                value_accessor = f"(*{value_accessor})"
            min_val = getattr(asn1_def, "min_val", 0) or 0
            max_val = getattr(asn1_def, "max_val", 0) or 0

            if "BIT STRING" in asn1_def.name.upper():
                value_code = f"""&BITSTRING{{
                    c:     aper.Constraint{{Lb: {min_val}, Ub: {max_val}}},
                    ext:   {str(is_extensible).lower()},
                    Value: aper.BitString{{
                        Bytes:   []byte({value_accessor}),
                        NumBits: uint64(len({value_accessor}) * 8),
                    }},
                }}"""
            else:
                value_code = f"""&OCTETSTRING{{
                    c:     aper.Constraint{{Lb: {min_val}, Ub: {max_val}}},
                    ext:   {str(is_extensible).lower()},
                    Value: aper.OctetString({value_accessor}),
                }}"""

        elif isinstance(asn1_def, ListDefinition):
            item_type_go_name = pascal_case_converter(asn1_def.of_type)
            tmp_var_name = f"tmp_{field_name}"

            min_val_str = asn1_def.min_val if asn1_def.min_val is not None else "0"
            if not min_val_str.isdigit():
                min_val_str = pascal_case_converter(min_val_str)

            max_val_str = asn1_def.max_val if asn1_def.max_val is not None else "0"
            if not max_val_str.isdigit():
                max_val_str = pascal_case_converter(max_val_str)

            value_code = f"&{tmp_var_name}"

            pre_append_code = f"""
            {tmp_var_name} := Sequence[aper.IE]{{
			    c:   aper.Constraint{{Lb: {min_val_str}, Ub: {max_val_str}}},
			    ext: {str(is_extensible).lower()},
		    }}
		    for i := 0; i < len(msg.{field_name}); i++ {{
			    {tmp_var_name}.Value = append({tmp_var_name}.Value, &msg.{field_name}[i])
		    }}"""

        else:
            value_code = f"msg.{field_name}"
            if not is_optional:
                value_code = f"&{value_code}"

        ie_id_const = pascal_case_converter(ie.id)
        if ie_id_const.startswith("Id"):
            ie_id_const = ie_id_const[2:]
        ie_id_const = f"ProtocolIEID{ie_id_const}"

        crit_const = f"Criticality{pascal_case_converter(ie.criticality)}"

        full_ie_append_code = f"""
        {{
            {pre_append_code}
            ies = append(ies, E1APMessageIE{{
			    Id:          {ie_id_const}, 
			    Criticality: Criticality{{Value: {crit_const}}},
			    Value:       {value_code},
		    }})
        }}"""

        final_block_for_this_ie = full_ie_append_code
        if is_optional:
            check = f"msg.{field_name} != nil"
            if isinstance(asn1_def, ListDefinition):
                check = f"len(msg.{field_name}) > 0"
            final_block_for_this_ie = f"if {check} {{\n{full_ie_append_code}\n}}"

        body_parts.append(final_block_for_this_ie)
    full_body = "\n".join(body_parts)

    final_func = f"""
func (msg *{go_name}) toIes() ([]E1APMessageIE, error) {{
	ies := make([]E1APMessageIE, 0)
    {full_body}
	return ies, nil
}}"""

    return final_func


def _generate_pdu_encode(
    go_name: str,
    item: SequenceDefinition,
    message_to_procedure_map: dict,
    procedures: dict,
) -> str:
    """Generates the Encode() dispatcher method for a PDU."""

    print(f"// TODO: Generating PDU Encode for {go_name}")
    return f"// Encode function for {go_name} to be generated here.\n"


def _generate_pdu_decode(
    go_name: str, item: SequenceDefinition
) -> Tuple[str, Set[str]]:
    """Generates the Decode() setup and validation method for a PDU."""

    print(f"// TODO: Generating PDU Decode for {go_name}")
    imports = set()
    return f"// Decode function for {go_name} to be generated here.\n", imports


def _generate_decoder_helper(
    go_name: str, item: SequenceDefinition, parser: ASN1Parser
) -> Tuple[str, Set[str]]:
    """Generates the Decoder helper struct and its decodeIE() method with the switch statement."""

    print(f"// TODO: Generating Decoder helper for {go_name}")
    imports = set()
    return f"// Decoder helper for {go_name} to be generated here.\n", imports


def render_internal_struct_methods(
    go_name: str, item: SequenceDefinition | ChoiceDefinition, parser: ASN1Parser
) -> Tuple[str, Set[str]]:
    """
    Generates Encode and Decode methods for an internal SEQUENCE or CHOICE struct.
    """
    required_imports = {"fmt", parser.aper_import_path}
    if any(ie.presence in ["optional", "conditional"] for ie in item.ies):
        required_imports.add("io")

    encode_body = ""
    decode_body = ""

    if isinstance(item, SequenceDefinition):
        encode_body = _generate_sequence_encode_body(item, parser)
        decode_body = _generate_sequence_decode_body(go_name, item, parser)
    elif isinstance(item, ChoiceDefinition):
        encode_body = _generate_choice_encode_body(go_name, item, parser)
        decode_body = _generate_choice_decode_body(go_name, item, parser)

    encode_func = f"""
// Encode implements the aper.AperMarshaller interface.
func (s *{go_name}) Encode(w *aper.AperWriter) (err error) {{
    {encode_body}
    return nil
}}
"""
    decode_func = f"""
// Decode implements the aper.AperUnmarshaller interface.
func (s *{go_name}) Decode(r *aper.AperReader) (err error) {{
    {decode_body}
    return nil
}}
"""
    return f"{encode_func}\n{decode_func}", required_imports


def render_enum_methods(go_name: str, item: EnumDefinition) -> Tuple[str, Set[str]]:
    """
    Generates Encode and Decode methods for an ENUMERATED type.
    """
    required_imports = {"fmt", "github.com/lvdund/ngap/aper"} # Use your aper import path

    # Determine the constraints for the enumeration
    num_enums = len(item.enum_values)
    upper_bound = num_enums - 1 if num_enums > 0 else 0
    is_extensible = str(item.is_extensible).lower()

    # --- Generate the Encode method ---
    encode_body = f'return w.WriteEnumerate(uint64(e.Value), aper.Constraint{{Lb: 0, Ub: {upper_bound}}}, {is_extensible})'

    encode_func = f"""
// Encode implements the aper.AperMarshaller interface.
func (e *{go_name}) Encode(w *aper.AperWriter) error {{
    {encode_body}
}}"""

    # --- Generate the Decode method ---
    decode_body = f"""
    val, err := r.ReadEnumerate(aper.Constraint{{Lb: 0, Ub: {upper_bound}}}, {is_extensible})
	if err != nil {{
		return err
	}}
	e.Value = aper.Enumerated(val)
	return nil"""

    decode_func = f"""
// Decode implements the aper.AperUnmarshaller interface.
func (e *{go_name}) Decode(r *aper.AperReader) error {{
    {decode_body}
}}"""

    return f"{encode_func}\n\n{decode_func}", required_imports


def render_extension_methods(go_name: str) -> Tuple[str, Set[str]]:
    """
    Generates Encode and Decode methods for a type-safe extension struct.
    This is a complex pattern involving bitmaps and open types.
    """
    required_imports = {"io", "fmt", "github.com/lvdund/ngap/aper"}

    print(f"// TODO: Generating extension methods for {go_name}")

    encode_code = f"""
func (s *{go_name}) Encode(w *aper.AperWriter) error {{
    // Extension container encoding logic for {go_name} to be generated here.
    return fmt.Errorf("Encode not implemented for extension struct {go_name}")
}}
"""
    decode_code = f"""
func (s *{go_name}) Decode(r *aper.AperReader) error {{
    // Extension container decoding logic for {go_name} to be generated here.
    return fmt.Errorf("Decode not implemented for extension struct {go_name}")
}}
"""
    return f"{encode_code}\n{decode_code}", required_imports


def render_list_methods(
    go_name: str, item: ListDefinition, parser: ASN1Parser
) -> Tuple[str, Set[str]]:
    """
    Generates Encode and Decode methods for a LIST (SEQUENCE OF) type.
    """
    required_imports = {"fmt", parser.aper_import_path}

    of_type_go_name = parser.pascal_case_converter(item.of_type)

    pascal_case_converter = parser.pascal_case_converter

    min_val_str = item.min_val if item.min_val is not None else "0"
    if not min_val_str.isdigit():
        min_val_str = pascal_case_converter(min_val_str)

    max_val_str = item.max_val if item.max_val is not None else "0"
    if not max_val_str.isdigit():
        max_val_str = pascal_case_converter(max_val_str)
    is_extensible = str(getattr(item, "is_extensible", False)).lower()

    encode_body = f"""
    // 1. Create a temporary slice of the INTERFACE type.
    itemPointers := make([]aper.AperMarshaller, len(s.Value)) // <-- FIX: Changed from []*aper.AperMarshaller
    for i := 0; i < len(s.Value); i++ {{
        // Assigning a pointer-to-struct to an interface value is correct.
        itemPointers[i] = &s.Value[i] 
    }}

    // 2. Call the generic WriteSequenceOf helper with the slice of interfaces.
    if err = aper.WriteSequenceOf(itemPointers, w, &aper.Constraint{{Lb: {min_val_str}, Ub: {max_val_str}}}, {is_extensible}); err != nil {{
        return fmt.Errorf("WriteSequenceOf for {go_name} failed: %w", err)
    }}"""

    decode_body = f"""
    // 1. Create a decoder function for the item type.
    decoder := func(r *aper.AperReader) (*{of_type_go_name}, error) {{
        item := new({of_type_go_name})
        if err := item.Decode(r); err != nil {{
            return nil, err
        }}
        return item, nil
    }}

    // 2. Call the generic ReadSequenceOf helper.
    //    The variable type `[]AlternativeQoSParaSetItem` now matches the function's return type.
    var decodedItems []{of_type_go_name} // <--- FIX: Removed the '*'
    if decodedItems, err = aper.ReadSequenceOf(decoder, r, &aper.Constraint{{Lb: {min_val_str}, Ub: {max_val_str}}}, {is_extensible}); err != nil {{
        return fmt.Errorf("ReadSequenceOf for {go_name} failed: %w", err)
    }}

    // 3. Assign the decoded slice of values directly.
    s.Value = decodedItems"""

    encode_func = f"""
func (s *{go_name}) Encode(w *aper.AperWriter) (err error) {{
    {encode_body}
    return nil
}}
"""
    decode_func = f"""
func (s *{go_name}) Decode(r *aper.AperReader) (err error) {{
    {decode_body}
    return nil
}}
"""
    return f"{encode_func}\n{decode_func}", required_imports


def _generate_direct_encode_call(
    field_name, ie, asn1_def, parser, is_struct_extensible
):
    """Generates the direct APER call for encoding a field inside a SEQUENCE."""
    pascal_case_converter = parser.pascal_case_converter
    is_optional = ie.presence in ["optional", "conditional"]
    is_extensible = getattr(asn1_def, "is_extensible", False)

    value_accessor = f"s.{field_name}"

    if isinstance(asn1_def, IntegerDefinition):
        if is_optional:
            value_accessor = f"(*{value_accessor})"

        return f'if err = w.WriteInteger(int64({value_accessor}), &aper.Constraint{{Lb: {asn1_def.min_val or 0}, Ub: {asn1_def.max_val or 0}}}, {str(is_extensible).lower()}); err != nil {{ return fmt.Errorf("Encode {field_name} failed: %w", err) }}'

    elif isinstance(asn1_def, EnumDefinition):
        if is_optional:
            value_accessor = f"(*{value_accessor})"

        num_enums = len(asn1_def.enum_values)
        upper_bound = num_enums - 1 if num_enums > 0 else 0

        return f'if err = w.WriteEnumerate(uint64({value_accessor}.Value), aper.Constraint{{Lb: 0, Ub: {upper_bound}}}, {str(is_extensible).lower()}); err != nil {{ return fmt.Errorf("Encode {field_name} failed: %w", err) }}'
    elif isinstance(asn1_def, StringDefinition):
        if is_optional:
            value_accessor = f"(*{value_accessor})"
        if asn1_def.string_type == "BIT STRING":
            return f"if err = w.WriteBitString({value_accessor}.Bytes, uint({value_accessor}.NumBits), &aper.Constraint{{...}}); err != nil {{...}}"
        else:
            return f"if err = w.WriteOctetString([]byte({value_accessor}), &aper.Constraint{{...}}); err != nil {{...}}"

    else:
        return f'if err = s.{field_name}.Encode(w); err != nil {{ return fmt.Errorf("Encode {field_name} failed: %w", err) }}'


def _generate_direct_decode_call(
    field_name, go_type, ie, asn1_def, parser, is_struct_extensible
):
    """Generates the direct APER call for decoding a field inside a SEQUENCE."""
    is_optional = ie.presence in ["optional", "conditional"]

    if isinstance(asn1_def, (IntegerDefinition, EnumDefinition, StringDefinition)):

        return f"// TODO: Decode logic for primitive field {field_name}"
    else:
        alloc_and_decode = f'if err = s.{field_name}.Decode(r); err != nil {{ return fmt.Errorf("Decode {field_name} failed: %w", err) }}'
        if is_optional:
            alloc_and_decode = (
                f"s.{field_name} = new({go_type})\n\t\t{alloc_and_decode}"
            )
        return alloc_and_decode


def _generate_sequence_encode_body(item: SequenceDefinition, parser: ASN1Parser) -> str:
    """Generates the Go code for the body of a SEQUENCE Encode method."""
    pascal_case_converter = parser.pascal_case_converter
    encode_parts = []

    is_struct_extensible = item.is_extensible
    encode_parts.append(
        f'if err = w.WriteBool({str(is_struct_extensible).lower()}); err != nil {{ return fmt.Errorf("Encode extensibility bool failed: %w", err) }}'
    )

    optional_fields = [
        ie for ie in item.ies if ie.presence in ["optional", "conditional"]
    ]
    num_optional = len(optional_fields)
    if num_optional > 0:
        bitmap_len = (num_optional + 7) // 8
        encode_parts.append(f"var optionalityBitmap [{bitmap_len}]byte")
        for i, opt_ie in enumerate(optional_fields):
            opt_field_name = pascal_case_converter(opt_ie.ie)
            byte_index = i // 8
            bit_index = 7 - (i % 8)
            encode_parts.append(
                f"if s.{opt_field_name} != nil {{ optionalityBitmap[{byte_index}] |= 1 << {bit_index} }}"
            )
        encode_parts.append(
            f'if err = w.WriteBitString(optionalityBitmap[:], uint({num_optional}), &aper.Constraint{{Lb: {num_optional}, Ub: {num_optional}}}, false); err != nil {{ return fmt.Errorf("Encode optionality bitmap failed: %w", err) }}'
        )

    for ie in item.ies:
        field_name = pascal_case_converter(ie.ie)
        _, asn1_def = parser.go_type_resolver(ie.type)
        is_optional = ie.presence in ["optional", "conditional"]

        encode_call = _generate_direct_encode_call(field_name, ie, asn1_def, parser)

        if is_optional:
            encode_parts.append(f"if s.{field_name} != nil {{ {encode_call} }}")
        else:
            encode_parts.append(encode_call)

    return "\n\t".join(encode_parts)


def _generate_sequence_decode_body(
    go_name: str, item: SequenceDefinition, parser: ASN1Parser
) -> str:
    """Generates the Go code for the body of a SEQUENCE Decode method."""
    pascal_case_converter = parser.pascal_case_converter
    decode_parts = []

    decode_parts.append(
        'var isExtensible bool\n\tif isExtensible, err = r.ReadBool(); err != nil { return fmt.Errorf("Read extensibility bool failed: %w", err) }'
    )

    optional_fields = [
        ie for ie in item.ies if ie.presence in ["optional", "conditional"]
    ]
    num_optional = len(optional_fields)
    if num_optional > 0:
        decode_parts.append("var optionalityBitmap []byte")
        decode_parts.append(
            f'if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{{Lb: {num_optional}, Ub: {num_optional}}}, false); err != nil {{ return fmt.Errorf("Read optionality bitmap failed: %w", err) }}'
        )

    opt_idx = 0
    for ie in item.ies:
        field_name = pascal_case_converter(ie.ie)
        go_type, asn1_def = parser.go_type_resolver(ie.type)
        is_optional = ie.presence in ["optional", "conditional"]

        decode_call = _generate_direct_decode_call(
            field_name, go_type, ie, asn1_def, parser
        )

        if go_type == "ProtocolExtensionContainer" and hasattr(
            ie, "extension_set_name"
        ):

            if parser.extension_sets.get(ie.extension_set_name):

                extension_go_type = f"{go_name}Extensions"

                decode_statement = f'if err = s.{field_name}.Decode(r); err != nil {{ return fmt.Errorf("Decode {field_name} failed: %w", err) }}'
                decode_call = (
                    f"s.{field_name} = new({extension_go_type})\n\t\t{decode_statement}"
                )

        if is_optional:
            byte_index = opt_idx // 8
            bit_index = 7 - (opt_idx % 8)
            check = f"len(optionalityBitmap) > {byte_index} && optionalityBitmap[{byte_index}] & (1 << {bit_index}) > 0"

            complete_if_block = f"""if {check} {{
                {decode_call}
            }}"""
            decode_parts.append(complete_if_block)
            opt_idx += 1
        else:
            decode_parts.append(decode_call)

    decode_parts.append(
        f'if isExtensible {{ return fmt.Errorf("Extensions not yet implemented for {go_name}") }}'
    )
    return "\n\t".join(decode_parts)


def _generate_choice_encode_body(
    go_name: str, item: ChoiceDefinition, parser: ASN1Parser
) -> str:
    """Generates the Go code for the body of a CHOICE Encode method."""
    pascal_case_converter = parser.pascal_case_converter

    num_choices = len(item.ies)
    ub_choices = num_choices - 1 if num_choices > 0 else 0
    is_extensible = str(item.is_extensible).lower()

    encode_switch_cases = []
    for choice_ie in item.ies:
        choice_name = pascal_case_converter(choice_ie.ie)
        const_name = f"{go_name}Present{choice_name}"
        
        _, concrete_def = parser.go_type_resolver(choice_ie.type)
        
        encode_call = ""
        # --- THIS IS THE FIX ---
        if isinstance(concrete_def, ListDefinition):
            # If it's a list, create the temporary Sequence wrapper and encode it
            of_type_go_name = pascal_case_converter(concrete_def.of_type)
            min_val_str = concrete_def.min_val or "0" # Add safety checks
            max_val_str = concrete_def.max_val or "0"
            if not max_val_str.isdigit(): max_val_str = pascal_case_converter(max_val_str)
            
            encode_call = f"""
        tmp_wrapper := Sequence[*{of_type_go_name}]{{
            c:   aper.Constraint{{Lb: {min_val_str}, Ub: {max_val_str}}},
            ext: {str(getattr(concrete_def, 'is_extensible', False)).lower()},
        }}
        for _, i := range *s.{choice_name} {{
            tmp_wrapper.Value = append(tmp_wrapper.Value, &i)
        }}
        if err = tmp_wrapper.Encode(w); err != nil {{
            return fmt.Errorf("Encode {choice_name} failed: %w", err)
        }}"""
        else:
            # Otherwise, it's a normal struct, so call its Encode method
            encode_call = f"""
        if err = s.{choice_name}.Encode(w); err != nil {{
            return fmt.Errorf("Encode {choice_name} failed: %w", err)
        }}"""
        
        encode_switch_cases.append(f"case {const_name}:{encode_call}")

    encode_switch_body = "\n".join(encode_switch_cases)

    return f"""
    // 1. Write the choice index.
    if err = w.WriteChoice(uint64(s.Choice-1), {ub_choices}, {is_extensible}); err != nil {{
        return fmt.Errorf("Encode choice index failed: %w", err)
    }}
    
    // 2. Encode the selected member.
    switch s.Choice {{
    {encode_switch_body}
    default:
        return fmt.Errorf("Encode choice of {go_name} with unknown choice value %d", s.Choice)
    }}"""


def _generate_choice_decode_body(
    go_name: str, item: ChoiceDefinition, parser: ASN1Parser
) -> str:
    """Generates the Go code for the body of a CHOICE Decode method."""
    pascal_case_converter = parser.pascal_case_converter

    num_choices = len(item.ies)
    ub_choices = num_choices - 1 if num_choices > 0 else 0
    is_extensible = str(item.is_extensible).lower()

    decode_switch_cases = []
    for i, choice_ie in enumerate(item.ies):
        choice_name = pascal_case_converter(choice_ie.ie)
        const_name = f"{go_name}Present{choice_name}"
        
        base_go_type, concrete_def = parser.go_type_resolver(choice_ie.type)

        alloc_and_decode = ""
        if isinstance(concrete_def, ListDefinition):
            of_type_go_name = pascal_case_converter(concrete_def.of_type)
            min_val_str = concrete_def.min_val or "0"
            max_val_str = concrete_def.max_val or "0"
            if not max_val_str.isdigit(): max_val_str = pascal_case_converter(max_val_str)
            
            alloc_and_decode = f"""
        // 1. Create a DECODER function for the list item, as required by ReadSequenceOf.
        itemDecoder := func(r *aper.AperReader) (*{of_type_go_name}, error) {{
            item := new({of_type_go_name})
            if err := item.Decode(r); err != nil {{
                return nil, err
            }}
            return item, nil
        }}

        // 2. Decode the list using the aper library's generic function.
        var decodedItems []{of_type_go_name}
        if decodedItems, err = aper.ReadSequenceOf(itemDecoder, r, &aper.Constraint{{Lb: {min_val_str}, Ub: {max_val_str}}}, {str(getattr(concrete_def, 'is_extensible', False)).lower()}); err != nil {{
            return fmt.Errorf("Decode {choice_name} failed: %w", err)
        }}
        
        // 3. Assign the decoded slice of values.
        s.{choice_name} = &decodedItems"""
        else:
            final_go_type = base_go_type
            alloc_and_decode = f"""
        s.{choice_name} = new({final_go_type})
        if err = s.{choice_name}.Decode(r); err != nil {{
            return fmt.Errorf("Decode {choice_name} failed: %w", err)
        }}"""
        
        decode_switch_cases.append(f"case {i}:{alloc_and_decode}")
        
    decode_switch_body = "\n".join(decode_switch_cases)

    return f"""
    // 1. Read the choice index.
    var choice uint64
    if choice, err = r.ReadChoice({ub_choices}, {is_extensible}); err != nil {{
        return fmt.Errorf("Read choice index failed: %w", err)
    }}

    // 2. Decode the selected member.
    switch choice {{
    {decode_switch_body}
    default:
        return fmt.Errorf("Decode choice of {go_name} with unknown choice index %d", choice)
    }}"""


def _generate_direct_encode_call(field_name, ie, asn1_def, parser):
    """Helper to generate the Go code line for encoding a single field."""
    is_optional = ie.presence in ["optional", "conditional"]
    is_extensible = getattr(asn1_def, "is_extensible", False)

    value_accessor = f"s.{field_name}"

    if isinstance(asn1_def, IntegerDefinition):
        if is_optional:
            value_accessor = f"(*{value_accessor})"
        return f'if err = w.WriteInteger(int64({value_accessor}), &aper.Constraint{{Lb: {asn1_def.min_val or 0}, Ub: {asn1_def.max_val or 0}}}, {str(is_extensible).lower()}); err != nil {{ return fmt.Errorf("Encode {field_name} failed: %w", err) }}'

    elif isinstance(asn1_def, EnumDefinition):
        if is_optional:
            value_accessor = f"(*{value_accessor})"
        num_enums = len(asn1_def.enum_values)
        upper_bound = num_enums - 1 if num_enums > 0 else 0
        return f'if err = w.WriteEnumerate(uint64({value_accessor}.Value), aper.Constraint{{Lb: 0, Ub: {upper_bound}}}, {str(is_extensible).lower()}); err != nil {{ return fmt.Errorf("Encode {field_name} failed: %w", err) }}'

    elif isinstance(asn1_def, StringDefinition):
        if is_optional:
            value_accessor = f"(*{value_accessor})"
        if asn1_def.string_type == "BIT STRING":
            return f'if err = w.WriteBitString({value_accessor}.Bytes, uint({value_accessor}.NumBits), &aper.Constraint{{Lb: {asn1_def.min_val or 0}, Ub: {asn1_def.max_val or 0}}}, {str(is_extensible).lower()}); err != nil {{ return fmt.Errorf("Encode {field_name} failed: %w", err) }}'
        else:
            return f'if err = w.WriteOctetString([]byte({value_accessor}), &aper.Constraint{{Lb: {asn1_def.min_val or 0}, Ub: {asn1_def.max_val or 0}}}, {str(is_extensible).lower()}); err != nil {{ return fmt.Errorf("Encode {field_name} failed: %w", err) }}'

    else:
        return f'if err = s.{field_name}.Encode(w); err != nil {{ return fmt.Errorf("Encode {field_name} failed: %w", err) }}'


def _generate_direct_decode_call(field_name, go_type, ie, asn1_def, parser):
    """
    Generates the Go code for decoding a single field inside a SEQUENCE,
    correctly handling mandatory (value) vs. optional (pointer) assignments.
    """
    is_optional = ie.presence in ["optional", "conditional"]
    is_extensible = getattr(asn1_def, 'is_extensible', False)

    # ==============================================================================
    # == INTEGER DEFINITION
    # ==============================================================================
    if isinstance(asn1_def, IntegerDefinition):
        assignment = f"s.{field_name} = {go_type}(val)"
        if is_optional:
            assignment = f"tmp := {go_type}(val)\n\t\t\ts.{field_name} = &tmp"

        return f"""
        {{
            var val int64
            if val, err = r.ReadInteger(&aper.Constraint{{Lb: {asn1_def.min_val or 0}, Ub: {asn1_def.max_val or 0}}}, {str(is_extensible).lower()}); err != nil {{
                 return fmt.Errorf("Decode {field_name} failed: %w", err)
            }}
            {assignment}
        }}
        """

    # ==============================================================================
    # == ENUM DEFINITION
    # ==============================================================================
    elif isinstance(asn1_def, EnumDefinition):
        num_enums = len(asn1_def.enum_values)
        upper_bound = num_enums - 1 if num_enums > 0 else 0
        
        decode_block = f"""
        {{
            var val uint64
            if val, err = r.ReadEnumerate(aper.Constraint{{Lb: 0, Ub: {upper_bound}}}, {str(is_extensible).lower()}); err != nil {{
                return fmt.Errorf("Decode {field_name} failed: %w", err)
            }}
            s.{field_name}.Value = aper.Enumerated(val)
        }}"""

        if is_optional:
            return f's.{field_name} = new({go_type})\n\t\t{decode_block}'
        return decode_block

    # ==============================================================================
    # == STRING DEFINITION (and Built-ins)
    # ==============================================================================
    elif isinstance(asn1_def, StringDefinition) or (isinstance(asn1_def, BuiltinDefinition) and "STRING" in asn1_def.name.upper()):
        min_val = getattr(asn1_def, 'min_val', 0) or 0
        max_val = getattr(asn1_def, 'max_val', 0) or 0

        if "BIT STRING" in asn1_def.name.upper():
            assignment = f"s.{field_name} = val"
            if is_optional:
                assignment = f"s.{field_name} = &val"
            
            return f"""
            {{
                var val aper.BitString
                if val.Bytes, val.NumBits, err = r.ReadBitString(&aper.Constraint{{Lb: {min_val}, Ub: {max_val}}}, {str(is_extensible).lower()}); err != nil {{
                    return fmt.Errorf("Decode {field_name} failed: %w", err)
                }}
                {assignment}
            }}
            """
        else: # OCTET STRING
            assignment = f"s.{field_name} = {go_type}(val)"
            if is_optional:
                assignment = f"tmp := {go_type}(val)\n\t\t\ts.{field_name} = &tmp"

            return f"""
            {{
                var val []byte
                if val, err = r.ReadOctetString(&aper.Constraint{{Lb: {min_val}, Ub: {max_val}}}, {str(is_extensible).lower()}); err != nil {{
                    return fmt.Errorf("Decode {field_name} failed: %w", err)
                }}
                {assignment}
            }}
            """
            
    # ==============================================================================
    # == COMPLEX TYPES (SEQUENCE, CHOICE, LIST)
    # ==============================================================================
    else:
        decode_statement = f'if err = s.{field_name}.Decode(r); err != nil {{ return fmt.Errorf("Decode {field_name} failed: %w", err) }}'
        if is_optional:
             return f's.{field_name} = new({go_type})\n\t\t{decode_statement}'
        return decode_statement

def _generate_pdu_encode(go_name: str, item: SequenceDefinition, message_to_procedure_map: dict, procedures: dict, parser: ASN1Parser) -> str:
    """Generates the Encode() dispatcher method for a PDU."""
    
    # Find the procedure for this message to get the ProcedureCode
    proc_name = next((proc.name for proc in procedures.values() if go_name in [proc.initiating_message, proc.successful_outcome, proc.unsuccessful_outcome]), None)
    if not proc_name or proc_name not in procedures:
        return f'// Encode for {go_name}: Could not find associated procedure.'

    procedure = procedures[proc_name]
    proc_code_const_base = parser.pascal_case_converter(procedure.procedure_code)
    if proc_code_const_base.startswith("Id"):
        proc_code_const_base = proc_code_const_base[2:]
    proc_code_const = f"ProcedureCode{proc_code_const_base}"

    # Determine the message type and the master encoder function to call
    message_type = message_to_procedure_map.get(go_name)
    master_encoder_func = ""
    if message_type == "InitiatingMessage":
        master_encoder_func = "EncodeInitiatingMessage"
    elif message_type == "SuccessfulOutcome":
        master_encoder_func = "EncodeSuccessfulOutcome"
    elif message_type == "UnsuccessfulOutcome":
        master_encoder_func = "EncodeUnsuccessfulOutcome"
    else:
        return f'// Encode for {go_name}: Could not determine message type for master encoder.'

    pdu_crit = "CriticalityReject" if message_type == "InitiatingMessage" else "CriticalityIgnore"

    return f"""
// Encode implements the aper.AperMarshaller interface for {go_name}.
func (msg *{go_name}) Encode(w io.Writer) error {{
	ies, err := msg.toIes()
	if err != nil {{
		return fmt.Errorf("could not convert {go_name} to IEs: %w", err)
	}}

	return {master_encoder_func}(w, {proc_code_const}, Criticality{{Value: {pdu_crit}}}, ies)
}}"""


def _generate_pdu_decode(go_name: str, item: SequenceDefinition, parser: ASN1Parser) -> Tuple[str, Set[str]]:
    """Generates the Decode() setup and validation method for a PDU."""
    pascal_case_converter = parser.pascal_case_converter
    
    validation_parts = []
    for ie in item.ies:
        if ie.presence == "mandatory":
            ie_id_const_base = pascal_case_converter(ie.id)
            if ie_id_const_base.startswith("Id"):
                ie_id_const_base = ie_id_const_base[2:]
            ie_id_const = f"ProtocolIEID{ie_id_const_base}"
            field_name = pascal_case_converter(ie.ie)
            validation_parts.append(f"""
    if _, ok := decoder.list[{ie_id_const}]; !ok {{
		err = fmt.Errorf("mandatory field {field_name} is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{{
			IECriticality: Criticality{{Value: CriticalityReject}}, // Or from IE spec
			IEID:          ProtocolIEID{{Value: {ie_id_const}}},
			TypeOfError:   TypeOfError{{Value: TypeOfErrorMissing}},
		}})
	}}""")
    
    # Add a final return if any mandatory field was missing.
    if validation_parts:
        validation_parts.append("if err != nil { return }")

    validation_block = "\n".join(validation_parts)
    decoder_name = f"{go_name}Decoder"
    
    return f"""
// Decode implements the aper.AperUnmarshaller interface for {go_name}.
func (msg *{go_name}) Decode(buf []byte) (err error, diagList []CriticalityDiagnosticsIEItem) {{
	defer func() {{
		if err != nil {{
			err = fmt.Errorf("{go_name}: %w", err)
		}}
	}}()

	r := aper.NewReader(bytes.NewReader(buf))
	
	decoder := {decoder_name}{{
		msg:  msg,
		list: make(map[aper.Integer]*E1APMessageIE),
	}}
	
	// aper.ReadSequenceOf will decode the IEs and call the callback for each one.
	if _, err = aper.ReadSequenceOf[E1APMessageIE](decoder.decodeIE, r, &aper.Constraint{{Lb: 0, Ub: 65535}}, false); err != nil {{
		return
	}}

    // After decoding all present IEs, validate that mandatory ones were found.
    {validation_block}

	return
}}""", set()

def _generate_decoder_helper(go_name: str, item: SequenceDefinition, parser: ASN1Parser) -> Tuple[str, Set[str]]:
    """Generates the Decoder helper struct and its decodeIE() method with the switch statement."""
    
    decoder_name = f"{go_name}Decoder"
    pascal_case_converter = parser.pascal_case_converter
    
    case_blocks = []
    for ie in item.ies:
        ie_id_const_base = pascal_case_converter(ie.id)
        if ie_id_const_base.startswith("Id"):
            ie_id_const_base = ie_id_const_base[2:]
        ie_id_const = f"ProtocolIEID{ie_id_const_base}"
        
        field_name = pascal_case_converter(ie.ie)
        go_type, asn1_def = parser.go_type_resolver(ie.type)
        
        decode_logic = _generate_direct_decode_call(field_name, go_type, ie, asn1_def, parser)

        case_blocks.append(f"""
    case {ie_id_const}:
        {decode_logic}""")

    switch_body = "\n".join(case_blocks)

    return f"""
type {decoder_name} struct {{
	msg      *{go_name}
	diagList []CriticalityDiagnosticsIEItem
	list     map[aper.Integer]*E1APMessageIE
}}

func (decoder *{decoder_name}) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {{
	var id int64
	var c uint64
	var buf []byte
	if id, err = r.ReadInteger(&aper.Constraint{{Lb: 0, Ub: 65535}}, false); err != nil {{
		return
	}}
	msgIe = new(E1APMessageIE)
	msgIe.Id.Value = aper.Integer(id)

	if c, err = r.ReadEnumerate(aper.Constraint{{Lb: 0, Ub: 2}}, false); err != nil {{
		return
	}}
	msgIe.Criticality.Value = aper.Enumerated(c)

	if buf, err = r.ReadOpenType(); err != nil {{
		return
	}}

	ieId := msgIe.Id.Value
	if _, ok := decoder.list[ieId]; ok {{
		err = fmt.Errorf("duplicated protocol IE ID %%d", ieId)
		return
	}}
	decoder.list[ieId] = msgIe

	ieR := aper.NewReader(bytes.NewReader(buf))
	msg := decoder.msg

	switch msgIe.Id.Value {{
    {switch_body}
	default:
		// Handle unknown IEs based on criticality here, if needed.
	}}
	return
}}""", {"bytes"}

# ==============================================================================
# == CATEGORY 4: EXTENSION STRUCT METHODS (Special Case)
# ==============================================================================

def render_extension_methods(go_name: str) -> Tuple[str, Set[str]]:
    """
    Generates placeholder Encode and Decode methods for a type-safe extension struct.
    """
    required_imports = {"fmt", "io"} # Aper import not strictly needed for placeholders

    encode_code = f"""
// Encode implements the aper.AperMarshaller interface.
func (s *{go_name}) Encode(w *aper.AperWriter) error {{
	// TODO: Implement the complex APER extension container encoding logic.
	// This involves creating a bitmap of present extensions and encoding each one as an open type.
	return fmt.Errorf("Encode not yet implemented for {go_name}")
}}
"""
    decode_code = f"""
// Decode implements the aper.AperUnmarshaller interface.
func (s *{go_name}) Decode(r *aper.AperReader) error {{
	// TODO: Implement the complex APER extension container decoding logic.
	// This involves reading a bitmap and then decoding each present extension as an open type.
	return fmt.Errorf("Decode not yet implemented for {go_name}")
}}
"""
    return f"{encode_code}\n{decode_code}", required_imports