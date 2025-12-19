from common_types import (
    SequenceDefinition,
    ChoiceDefinition,
    EnumDefinition,
    IntegerDefinition,
    StringDefinition,
    ListDefinition,
    BuiltinDefinition,
    ASN1Definition,
    InformationElement,
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
    """
    Generates the toIes() method body with a robust and clean structure.
    """
    pascal_case_converter = parser.pascal_case_converter
    go_type_resolver = parser.go_type_resolver

    body_parts = []
    
    for ie in item.ies:
        field_name = pascal_case_converter(ie.ie)
        _, asn1_def = go_type_resolver(ie.type)
        if not asn1_def:
            print(f"WARNING: Could not resolve type '{ie.type}' for IE '{ie.ie}' in PDU '{go_name}'. Skipping.")
            continue
        if not ie.id:
            continue

        is_optional = ie.presence in ["optional", "conditional"]
        is_extensible = getattr(asn1_def, 'is_extensible', False)
        
        value_code = ""
        pre_append_code = ""
        
        if isinstance(asn1_def, IntegerDefinition) and not isinstance(asn1_def, StringDefinition):
            
            value_code = f"""&INTEGER{{
                c:     aper.Constraint{{Lb: {asn1_def.min_val or 0}, Ub: {asn1_def.max_val or 0}}},
                ext:   {str(is_extensible).lower()},
                Value: msg.{field_name}.Value,
            }}"""
        
        elif isinstance(asn1_def, EnumDefinition):
            
            num_enums = len(asn1_def.enum_values)
            upper_bound = num_enums - 1 if num_enums > 0 else 0
            value_code = f"""&ENUMERATED{{
                c:     aper.Constraint{{Lb: 0, Ub: {upper_bound}}},
                ext:   {str(is_extensible).lower()},
                Value: msg.{field_name}.Value,
            }}"""

        elif isinstance(asn1_def, StringDefinition) or (isinstance(asn1_def, BuiltinDefinition) and "STRING" in asn1_def.name.upper()):
            
            min_val = getattr(asn1_def, "min_val", 0) or 0
            max_val = getattr(asn1_def, "max_val", 0) or 0
            
            
            if (isinstance(asn1_def, StringDefinition) and asn1_def.string_type == "BIT STRING") or \
            (isinstance(asn1_def, BuiltinDefinition) and "BIT STRING" in asn1_def.name.upper()):
                value_code = f"""&BITSTRING{{
                    c:     aper.Constraint{{Lb: {min_val}, Ub: {max_val}}},
                    ext:   {str(is_extensible).lower()},
                    Value: msg.{field_name}.Value,
                }}"""
            else: 
                value_code = f"""&OCTETSTRING{{
                    c:     aper.Constraint{{Lb: {min_val}, Ub: {max_val}}},
                    ext:   {str(is_extensible).lower()},
                    Value: msg.{field_name}.Value,
                }}"""

        elif isinstance(asn1_def, ListDefinition):
            tmp_var_name = f"tmp_{field_name}"
            min_val_str = asn1_def.min_val if asn1_def.min_val is not None else "0"
            if not min_val_str.isdigit(): min_val_str = pascal_case_converter(min_val_str)
            max_val_str = asn1_def.max_val if asn1_def.max_val is not None else "0"
            if not max_val_str.isdigit(): max_val_str = pascal_case_converter(max_val_str)

            _, of_type_def = go_type_resolver(asn1_def.of_type)
            loop_body = ""
            wrapper_type = None

            
            if isinstance(of_type_def, IntegerDefinition) and not isinstance(of_type_def, StringDefinition):
                wrapper_type = "INTEGER"
                value_constructor = "item.Value"
            elif isinstance(of_type_def, EnumDefinition):
                wrapper_type = "ENUMERATED"
                value_constructor = "item.Value"
            elif isinstance(of_type_def, StringDefinition) or (isinstance(of_type_def, BuiltinDefinition) and "STRING" in of_type_def.name.upper()):
                if "BIT STRING" in of_type_def.name.upper() or (isinstance(of_type_def, StringDefinition) and of_type_def.string_type == "BIT STRING"):
                    wrapper_type = "BITSTRING"
                else:
                    wrapper_type = "OCTETSTRING"
                value_constructor = "item.Value"

            
            list_accessor = f"msg.{field_name}"

            
            
            field_type_def = parser.definitions.get(ie.type)
            is_wrapper_struct = isinstance(field_type_def, ListDefinition)

            if is_wrapper_struct:
                
                
                list_accessor += ".Value"
            elif is_optional:
                
                
                list_accessor = f"(*{list_accessor})"

            if wrapper_type:
                
                value_constructor = "item.Value"
                loop_body = f"""
                for _, item := range {list_accessor} {{
                    wrapped_item := &{wrapper_type}{{
                        c:     aper.Constraint{{Lb: {getattr(of_type_def, 'min_val', 0) or 0}, Ub: {getattr(of_type_def, 'max_val', 0) or 0}}},
                        ext:   {str(getattr(of_type_def, 'is_extensible', False)).lower()},
                        Value: {value_constructor},
                    }}
                    {tmp_var_name}.Value = append({tmp_var_name}.Value, wrapped_item)
                }}"""
            else:
                
                loop_body = f"""
                for i := 0; i < len({list_accessor}); i++ {{
                    {tmp_var_name}.Value = append({tmp_var_name}.Value, &{list_accessor}[i])
                }}"""

            value_code = f"&{tmp_var_name}"
            pre_append_code = f"""
            {tmp_var_name} := Sequence[aper.IE]{{
                c:   aper.Constraint{{Lb: {min_val_str}, Ub: {max_val_str}}},
                ext: {str(is_extensible).lower()},
            }}
            {loop_body}
            """
        
        else: 
            value_code = f"msg.{field_name}"
            if not is_optional:
                value_code = f"&{value_code}"
        
        
        
        ie_id_const = pascal_case_converter(ie.id)
        if ie_id_const.startswith("Id"): ie_id_const = ie_id_const[2:]
        ie_id_const = f"ProtocolIEID{ie_id_const}"
        
        crit_const = f"Criticality{pascal_case_converter(ie.criticality)}"

        
        full_ie_append_code = f"""
            ies = append(ies, E1APMessageIE{{
			    Id:          ProtocolIEID{{Value: {ie_id_const}}},
			    Criticality: Criticality{{Value: {crit_const}}},
			    Value:       {value_code},
		    }})"""
        
        
        full_logic_block = f"{pre_append_code}\n{full_ie_append_code}"

        
        if is_optional:
            
            check = f"msg.{field_name} != nil"
            final_block = f"if {check} {{\n{full_logic_block}\n}}"
        elif pre_append_code.strip():
            # If there is pre-logic (like loops), keep the braces for scope safety (e.g. 'i' variable)
            final_block = f"{{\n{full_logic_block}\n}}" 
        else:
            # Simple mandatory IE, no braces needed
            final_block = full_logic_block
            
        body_parts.append(final_block)

    full_body = "\n\t".join(body_parts)
    
    
    final_func = f"""
// toIes transforms the {go_name} struct into a slice of E1APMessageIEs.
func (msg *{go_name}) toIes() ([]E1APMessageIE, error) {{
	ies := make([]E1APMessageIE, 0)
    {full_body}
	var err error
	return ies, err
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
    required_imports = {"fmt", "math", parser.aper_import_path}
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
    required_imports = {"fmt", "github.com/lvdund/ngap/aper"} 

    
    num_enums = len(item.enum_values)
    upper_bound = num_enums - 1 if num_enums > 0 else 0
    is_extensible = str(item.is_extensible).lower()

    
    encode_body = f'return w.WriteEnumerate(uint64(e.Value), aper.Constraint{{Lb: 0, Ub: {upper_bound}}}, {is_extensible})'

    encode_func = f"""
// Encode implements the aper.AperMarshaller interface.
func (e *{go_name}) Encode(w *aper.AperWriter) error {{
    {encode_body}
}}"""

    
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


def render_extension_methods(go_name: str, extension_set: list, parser: ASN1Parser) -> Tuple[str, Set[str]]:
    """
    Generates Encode and Decode methods for a type-safe extension struct.
    """
    required_imports = {"io", "fmt", "bytes", parser.aper_import_path}
    pascal_case_converter = parser.pascal_case_converter

    # --- ENCODE GENERATION ---
    encode_checks = []
    for ext in extension_set:
        field_name = pascal_case_converter(ext['id'])
        type_name = pascal_case_converter(ext['type'])
        
        id_const = f"ProtocolIEID{field_name}"
        crit_const = f"Criticality{pascal_case_converter(ext['crit'])}"
        
        encode_checks.append(f"""
    if s.{field_name} != nil {{
        extensions = append(extensions, &ProtocolExtensionField{{
            Id:          ProtocolIEID{{Value: {id_const}}},
            Criticality: Criticality{{Value: {crit_const}}},
            ExtensionValue: s.{field_name},
        }})
    }}""")

    encode_body = "\n".join(encode_checks)
    
    encode_func = f"""
// Encode implements the aper.AperMarshaller interface.
func (s *{go_name}) Encode(w *aper.AperWriter) error {{
    var extensions []*ProtocolExtensionField
    {encode_body}

    if len(extensions) > 0 {{
        itemPointers := make([]aper.AperMarshaller, len(extensions))
        for i := 0; i < len(extensions); i++ {{
            itemPointers[i] = extensions[i]
        }}
        if err := aper.WriteSequenceOf(itemPointers, w, &aper.Constraint{{Lb: 1, Ub: MaxProtocolExtensions}}, false); err != nil {{
            return fmt.Errorf("Encode extension container failed: %w", err)
        }}
    }} else {{
         if err := aper.WriteSequenceOf([]aper.AperMarshaller(nil), w, &aper.Constraint{{Lb: 1, Ub: MaxProtocolExtensions}}, false); err != nil {{
            return fmt.Errorf("Encode empty extension container failed: %w", err)
        }}
    }}
    return nil
}}"""

    # --- DECODE GENERATION ---
    decode_cases = []
    for ext in extension_set:
        field_name = pascal_case_converter(ext['id'])
        type_name = pascal_case_converter(ext['type'])
        id_const = f"ProtocolIEID{field_name}"
        
        decode_cases.append(f"""
        case {id_const}:
            s.{field_name} = new({type_name})
            if err := s.{field_name}.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {{
                return fmt.Errorf("Decode extension {field_name} failed: %w", err)
            }}""")

    decode_switch = "\n".join(decode_cases)

    decode_func = f"""
// Decode implements the aper.AperUnmarshaller interface.
func (s *{go_name}) Decode(r *aper.AperReader) error {{
    var decoder func(*aper.AperReader) (**ProtocolExtensionField, error)
    decoder = func(r *aper.AperReader) (**ProtocolExtensionField, error) {{
        var item ProtocolExtensionField
        if err := item.Decode(r); err != nil {{
            return nil, err
        }}
        ptr := &item
        return &ptr, nil
    }}

    var extensions []*ProtocolExtensionField
    var err error
    if extensions, err = aper.ReadSequenceOf[*ProtocolExtensionField](decoder, r, &aper.Constraint{{Lb: 1, Ub: MaxProtocolExtensions}}, false); err != nil {{
        return fmt.Errorf("Decode extension container failed: %w", err)
    }}

    for _, ext := range extensions {{
        switch ext.Id.Value {{
        {decode_switch}
        default:
            // Unknown extension, ignore
        }}
    }}
    return nil
}}"""

    return f"{encode_func}\n{decode_func}", required_imports


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
        
        min_val = asn1_def.min_val or 0
        if not str(min_val).isdigit(): 
            min_val = pascal_case_converter(str(min_val))
            if min_val and min_val[0].islower(): min_val = min_val[0].upper() + min_val[1:]
        
        max_val = asn1_def.max_val or 0
        max_val_str = str(max_val).strip()
        print(f"DEBUG: field={field_name}, max_val_raw={max_val}, max_val_str='{max_val_str}'")
        if max_val_str == "18446744073709551615":
             max_val = "math.MaxInt64" # Cap at MaxInt64 for Go int64 compatibility
        elif not max_val_str.isdigit(): 
            max_val = pascal_case_converter(max_val_str)
            if max_val and max_val[0].islower(): max_val = max_val[0].upper() + max_val[1:]
        print(f"DEBUG: field={field_name}, final max_val='{max_val}'")

        return f'if err = w.WriteInteger(int64({value_accessor}.Value), &aper.Constraint{{Lb: {min_val}, Ub: {max_val}}}, {str(is_extensible).lower()}); err != nil {{ return fmt.Errorf("Encode {field_name} failed: %w", err) }}'
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
             return f'if err = w.WriteBitString({value_accessor}.Value.Bytes, uint({value_accessor}.Value.NumBits), &aper.Constraint{{Lb: {asn1_def.min_val or 0}, Ub: {asn1_def.max_val or 0}}}, {str(is_extensible).lower()}); err != nil {{ return fmt.Errorf("Encode {field_name} failed: %w", err) }}'
        else:
             return f'if err = w.WriteOctetString([]byte({value_accessor}.Value), &aper.Constraint{{Lb: {asn1_def.min_val or 0}, Ub: {asn1_def.max_val or 0}}}, {str(is_extensible).lower()}); err != nil {{ return fmt.Errorf("Encode {field_name} failed: %w", err) }}'

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
        
        go_type, asn1_def = parser.go_type_resolver(ie.type)
        
        # FIX: Allow ProtocolExtensionContainer even if asn1_def is None
        if not asn1_def and ie.type != "ProtocolExtensionContainer":
            print(f"WARNING: Could not resolve type '{ie.type}' for field '{ie.ie}' in SEQUENCE '{item.name}'. Skipping field.")
            continue
            
        is_optional = ie.presence in ["optional", "conditional"]
        
        if ie.type == "ProtocolExtensionContainer":
             encode_call = f'if err = s.{field_name}.Encode(w); err != nil {{ return fmt.Errorf("Encode {field_name} failed: %w", err) }}'
        else:
             encode_call = _generate_direct_encode_call(field_name, ie, asn1_def, parser, is_struct_extensible)
        
        if is_optional:
            encode_parts.append(f'if s.{field_name} != nil {{ {encode_call} }}')
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
        _, asn1_def = parser.go_type_resolver(ie.type)
        
        
        semantic_go_type = pascal_case_converter(ie.type)
        
        
        decode_call = _generate_internal_decode_call(field_name, semantic_go_type, ie, asn1_def, parser)

        is_optional = ie.presence in ["optional", "conditional"]

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
        f'if isExtensible {{ /* TODO: Implement extension skipping for {go_name} */ }}'
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

    # --- DEBUGGING STEP 1: Log when this generator is run ---
    print(f"[GENERATOR DEBUG] Generating CHOICE encoder for '{go_name}' with {num_choices} members (ub={ub_choices}). Extensible: {is_extensible}")
    if num_choices <= 1 and not item.is_extensible:
        print(f"[GENERATOR WARNING] '{go_name}' is a non-extensible CHOICE with {num_choices} members. This may be problematic.")


    encode_switch_cases = []
    for choice_ie in item.ies:
        choice_name = pascal_case_converter(choice_ie.ie)
        const_name = f"{go_name}Present{choice_name}"
        
        base_go_type, concrete_def = parser.go_type_resolver(choice_ie.type)
        
        encode_call = ""
        
        if base_go_type == "string":
            encode_call = f"""
        if err = w.WriteOctetString([]byte(*s.{choice_name}), nil, false); err != nil {{
            return fmt.Errorf("Encode {choice_name} failed: %w", err)
        }}"""
        elif isinstance(concrete_def, IntegerDefinition):
            min_val = parser.pascal_case_converter(str(getattr(concrete_def, 'min_val', '0')))
            max_val = parser.pascal_case_converter(str(getattr(concrete_def, 'max_val', '0')))
            choice_is_extensible = str(getattr(concrete_def, 'is_extensible', False)).lower()
            encode_call = f"""
        if err = w.WriteInteger(int64(s.{choice_name}.Value), &aper.Constraint{{Lb: {min_val}, Ub: {max_val}}}, {choice_is_extensible}); err != nil {{
            return fmt.Errorf("Encode {choice_name} failed: %w", err)
        }}"""
        else:
            encode_call = f"""
        if err = s.{choice_name}.Encode(w); err != nil {{
            return fmt.Errorf("Encode {choice_name} failed: %w", err)
        }}"""
        
        encode_switch_cases.append(f"case {const_name}:{encode_call}")

    encode_switch_body = "\n".join(encode_switch_cases)

    # --- DEBUGGING STEP 2: Add a Printf to the generated Go code ---
    return f"""
    // 1. Write the choice index.
    // fmt.Printf("--- GO DEBUG: Encoding CHOICE {go_name} | Choice: %d, UpperBound: {ub_choices}, Extensible: {is_extensible}\\n", s.Choice-1) // UNCOMMENT FOR DEEP DEBUGGING
    if err = w.WriteChoice(uint64(s.Choice), {ub_choices}, {is_extensible}); err != nil {{
        return fmt.Errorf("Encode choice index failed for {go_name}: %w", err)
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
        
        
        base_go_type, concrete_def = parser.go_type_resolver(choice_ie.type)
        
        alloc_and_decode = ""
        
        if base_go_type == "string":
            
            alloc_and_decode = f"""
        var val []byte
        if val, err = r.ReadOctetString(nil, false); err != nil {{
            return fmt.Errorf("Decode {choice_name} failed: %w", err)
        }}
        tmpStr := string(val)
        s.{choice_name} = &tmpStr"""
        elif isinstance(concrete_def, IntegerDefinition):
            
            min_val = parser.pascal_case_converter(str(getattr(concrete_def, 'min_val', '0')))
            max_val = parser.pascal_case_converter(str(getattr(concrete_def, 'max_val', '0')))
            is_extensible = str(getattr(concrete_def, 'is_extensible', False)).lower()
            alloc_and_decode = f"""
        s.{choice_name} = new({base_go_type})
        var val int64
        if val, err = r.ReadInteger(&aper.Constraint{{Lb: {min_val}, Ub: {max_val}}}, {is_extensible}); err != nil {{
            return fmt.Errorf("Decode {choice_name} failed: %w", err)
        }}
        s.{choice_name}.Value = aper.Integer(val)"""
        else:
            
            alloc_and_decode = f"""
        s.{choice_name} = new({base_go_type})
        if err = s.{choice_name}.Decode(r); err != nil {{
            return fmt.Errorf("Decode {choice_name} failed: %w", err)
        }}"""
        
        
        decode_switch_cases.append(f"case {i+1}:{alloc_and_decode}")
        
    decode_switch_body = "\n".join(decode_switch_cases)

    return f"""
    // 1. Read the choice index (0-based) and assign it to the struct's Choice field.
    var choice uint64
    if choice, err = r.ReadChoice({ub_choices}, {is_extensible}); err != nil {{
        return fmt.Errorf("Read choice index failed: %w", err)
    }}
    s.Choice = choice // Choice is 1-based from ReadChoice

    // 2. Decode the selected member.
    switch choice {{
    {decode_switch_body}
    default:
        return fmt.Errorf("Decode choice of {go_name} with unknown choice index %d", choice)
    }}"""





def _generate_pdu_decode_call(field_name, go_type, ie, asn1_def, parser):
    is_optional = ie.presence in ["optional", "conditional"]
    is_extensible = getattr(asn1_def, 'is_extensible', False)

    

    if isinstance(asn1_def, IntegerDefinition):
        
        assignment = f"msg.{field_name}.Value = aper.Integer(val)"
        if is_optional:
            assignment = f"msg.{field_name} = new({go_type}); msg.{field_name}.Value = aper.Integer(val)"
        return f"""
        {{
            var val int64
            if val, err = ieR.ReadInteger(&aper.Constraint{{Lb: {asn1_def.min_val or 0}, Ub: {asn1_def.max_val or 0}}}, {str(is_extensible).lower()}); err != nil {{
                 return nil, fmt.Errorf("Decode {field_name} failed: %w", err)
            }}
            {assignment}
        }}"""

    elif isinstance(asn1_def, EnumDefinition):
        num_enums = len(asn1_def.enum_values)
        upper_bound = num_enums - 1 if num_enums > 0 else 0
        decode_block = f"""
        {{
            var val uint64
            if val, err = ieR.ReadEnumerate(aper.Constraint{{Lb: 0, Ub: {upper_bound}}}, {str(is_extensible).lower()}); err != nil {{
                return nil, fmt.Errorf("Decode {field_name} failed: %w", err)
            }}
            msg.{field_name}.Value = aper.Enumerated(val)
        }}"""
        if is_optional:
            return f'msg.{field_name} = new({go_type});\n\t\t{decode_block}'
        return decode_block

    elif isinstance(asn1_def, (StringDefinition, BuiltinDefinition)) and "STRING" in asn1_def.name.upper() and (go_type.startswith("aper.") or go_type == "string"):
        
        min_val = getattr(asn1_def, 'min_val', 0) or 0
        max_val = getattr(asn1_def, 'max_val', 0) or 0
        ext = str(is_extensible).lower()

        if "BIT STRING" in asn1_def.name.upper():
            assignment = f"msg.{field_name}.Value = val"
            if is_optional:
                assignment = f"msg.{field_name} = new({go_type}); msg.{field_name}.Value = val"
            return f"""
            {{
                var val aper.BitString
                if val.Bytes, val.NumBits, err = ieR.ReadBitString(&aper.Constraint{{Lb: {min_val}, Ub: {max_val}}}, {ext}); err != nil {{
                    return nil, fmt.Errorf("Decode {field_name} failed: %w", err)
                }}
                {assignment}
            }}"""
        else: 
            assignment = f"msg.{field_name}.Value = aper.OctetString(val)"
            if is_optional:
                assignment = f"msg.{field_name} = new({go_type}); msg.{field_name}.Value = aper.OctetString(val)"
            return f"""
            {{
                var val []byte
                if val, err = ieR.ReadOctetString(&aper.Constraint{{Lb: {min_val}, Ub: {max_val}}}, {ext}); err != nil {{
                    return nil, fmt.Errorf("Decode {field_name} failed: %w", err)
                }}
                {assignment}
            }}"""
        
    elif isinstance(asn1_def, ListDefinition):
        of_type_go_name = parser.pascal_case_converter(asn1_def.of_type)
        min_val_str = asn1_def.min_val or "0"
        if not min_val_str.isdigit(): min_val_str = parser.pascal_case_converter(min_val_str)
        
        max_val_str = asn1_def.max_val or "0"
        if not max_val_str.isdigit(): max_val_str = parser.pascal_case_converter(max_val_str)
        
        is_extensible = str(getattr(asn1_def, 'is_extensible', False)).lower()

        

        
        _, of_type_def = parser.go_type_resolver(asn1_def.of_type)

        
        item_decoder_body = ""
        if isinstance(of_type_def, IntegerDefinition):
            item_decoder_body = f"""
                var val int64
                if val, err = r.ReadInteger(&aper.Constraint{{Lb: {of_type_def.min_val or 0}, Ub: {of_type_def.max_val or 0}}}, {str(getattr(of_type_def, 'is_extensible', False)).lower()}); err != nil {{
                    return nil, err
                }}
                item := {of_type_go_name}(val)
                return &item, nil"""
        
        elif isinstance(of_type_def, EnumDefinition):
            
            num_enums = len(of_type_def.enum_values)
            upper_bound = num_enums - 1 if num_enums > 0 else 0
            item_decoder_body = f"""
                item := new({of_type_go_name})
                var val uint64
                if val, err = r.ReadEnumerate(aper.Constraint{{Lb: 0, Ub: {upper_bound}}}, {str(getattr(of_type_def, 'is_extensible', False)).lower()}); err != nil {{
                    return nil, err
                }}
                item.Value = aper.Enumerated(val)
                return item, nil"""

        elif isinstance(of_type_def, (StringDefinition, BuiltinDefinition)) and "STRING" in of_type_def.name.upper():
            min_val = getattr(of_type_def, 'min_val', 0) or 0
            max_val = getattr(of_type_def, 'max_val', 0) or 0
            ext = str(getattr(of_type_def, 'is_extensible', False)).lower()

            if "BIT STRING" in of_type_def.name.upper():
                 item_decoder_body = f"""
                 var val aper.BitString
                 if val.Bytes, val.NumBits, err = r.ReadBitString(&aper.Constraint{{Lb: {min_val}, Ub: {max_val}}}, {ext}); err != nil {{
                     return nil, err
                 }}
                 item := {of_type_go_name}(val)
                 return &item, nil"""
            else: 
                item_decoder_body = f"""
                var val []byte
                if val, err = r.ReadOctetString(&aper.Constraint{{Lb: {min_val}, Ub: {max_val}}}, {ext}); err != nil {{
                    return nil, err
                }}
                item := {of_type_go_name}(val)
                return &item, nil"""
        else:
            
            item_decoder_body = f"""
            item := new({of_type_go_name})
            if err := item.Decode(r); err != nil {{
                return nil, err
            }}
            return item, nil"""

        assignment = ""
        if is_optional:
            
            assignment = f"""
            msg.{field_name} = new({go_type})
            msg.{field_name}.Value = decodedItems"""
        else:
            
            assignment = f"msg.{field_name}.Value = decodedItems"

        return f"""
        {{
            itemDecoder := func(r *aper.AperReader) (*{of_type_go_name}, error) {{
                {item_decoder_body}
            }}
            var decodedItems []{of_type_go_name}
            if decodedItems, err = aper.ReadSequenceOf(itemDecoder, ieR, &aper.Constraint{{Lb: {min_val_str}, Ub: {max_val_str}}}, {is_extensible}); err != nil {{
                return nil, fmt.Errorf("Decode {field_name} failed: %w", err)
            }}
            {assignment}
        }}"""

        
    else: 
        decode_statement = f'if err = msg.{field_name}.Decode(ieR); err != nil {{ return nil, fmt.Errorf("Decode {field_name} failed: %w", err) }}'
        if is_optional:
             return f'msg.{field_name} = new({go_type});\n\t\t{decode_statement}'
        return decode_statement

def _generate_pdu_encode(
    go_name: str,
    item: SequenceDefinition,
    message_to_procedure_map: dict,
    procedures: dict,
    parser: ASN1Parser,
) -> str:
    """Generates the Encode() dispatcher method for a PDU."""
    
    pascal_case_converter = parser.pascal_case_converter
    
    asn1_msg_name = item.name

    found_proc = None
    for proc in procedures.values():
        if proc.initiating_message == asn1_msg_name or \
           proc.successful_outcome == asn1_msg_name or \
           proc.unsuccessful_outcome == asn1_msg_name:
            found_proc = proc
            break

    if not found_proc:
        return f'// Encode for {go_name}: Could not find associated procedure for ASN.1 name "{asn1_msg_name}".'

    procedure = found_proc
    proc_code_const_base = pascal_case_converter(procedure.procedure_code)
    if proc_code_const_base.startswith("Id"):
        proc_code_const_base = proc_code_const_base[2:]
    proc_code_const = f"ProcedureCode{proc_code_const_base}"

    message_type = message_to_procedure_map.get(asn1_msg_name)
    pdu_choice_const = ""
    
    # --- THIS IS THE CRITICAL FIX ---
    # We will be extremely explicit to override any faulty logic in the map.
    if go_name.endswith("Request"):
        message_type = "InitiatingMessage"
    elif go_name.endswith("Response") or go_name.endswith("Acknowledge") or go_name.endswith("Confirm") or go_name.endswith("Complete"):
         message_type = "SuccessfulOutcome"
    # --- END OF FIX ---
    
    if message_type == "InitiatingMessage":
        pdu_choice_const = "E1apPduInitiatingMessage"
    elif message_type == "SuccessfulOutcome":
        pdu_choice_const = "E1apPduSuccessfulOutcome"
    elif message_type == "UnsuccessfulOutcome":
        pdu_choice_const = "E1apPduUnsuccessfulOutcome"
    else:
        return f'// Could not determine PDU choice for {go_name}'

    pdu_crit_const = "CriticalityReject" if message_type == "InitiatingMessage" else "CriticalityIgnore"
    
    return f"""
// Encode implements the aper.AperMarshaller interface for {go_name}.
func (msg *{go_name}) Encode(w io.Writer) error {{
    ies, err := msg.toIes()
    if err != nil {{
        return fmt.Errorf("could not convert {go_name} to IEs: %w", err)
    }}

    return encodeMessage(w, {pdu_choice_const}, ProcedureCode{{Value: {proc_code_const}}}, Criticality{{Value: {pdu_crit_const}}}, ies)
}}"""


def _generate_pdu_decode(go_name: str, item: SequenceDefinition, parser: ASN1Parser) -> Tuple[str, Set[str]]:
    """Generates the Decode() setup and validation method for a PDU."""
    pascal_case_converter = parser.pascal_case_converter
    
    validation_parts = []
    for ie in item.ies:
        if ie.presence == "mandatory" and ie.id:
            ie_id_const_base = pascal_case_converter(ie.id)
            if ie_id_const_base.startswith("Id"):
                ie_id_const_base = ie_id_const_base[2:]
            ie_id_const = f"ProtocolIEID{ie_id_const_base}"
            field_name = pascal_case_converter(ie.ie)
            validation_parts.append(f"""
    if _, ok := decoder.list[ProtocolIEID{{Value: {ie_id_const}}}]; !ok {{
		err = fmt.Errorf("mandatory field {field_name} is missing")
		diagList = append(diagList, CriticalityDiagnosticsIEItem{{
			IECriticality: Criticality{{Value: CriticalityReject}},
			IEID:          ProtocolIEID{{Value: {ie_id_const}}},
			TypeOfError:   TypeOfError{{Value: TypeOfErrorMissing}},
		}})
	}}""")
    
    
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
		list: make(map[ProtocolIEID]*E1APMessageIE),
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
        if not ie.id:
            continue
        
        ie_id_const_base = pascal_case_converter(ie.id)
        if ie_id_const_base.startswith("Id"):
            ie_id_const_base = ie_id_const_base[2:]
        ie_id_const = f"ProtocolIEID{ie_id_const_base}"
        
        field_name = pascal_case_converter(ie.ie)
        
        semantic_go_type = pascal_case_converter(ie.type)
        
        _, asn1_def = parser.go_type_resolver(ie.type)

        
        decode_logic = _generate_pdu_decode_call(field_name, semantic_go_type, ie, asn1_def, parser)

        case_blocks.append(f"case {ie_id_const}:\n{decode_logic}")

    default_case = f"""default:
        switch msgIe.Criticality.Value {{
        case CriticalityReject:
            // If an unknown IE is critical, the PDU cannot be processed.
            err = fmt.Errorf("not comprehended IE ID %d (criticality: reject)", msgIe.Id.Value)
            decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{{
                IECriticality: msgIe.Criticality,
                IEID:          msgIe.Id,
                TypeOfError:   TypeOfError{{Value: TypeOfErrorNotUnderstood}},
            }})
        case CriticalityNotify:
            // Per 3GPP TS 38.463 Section 10.3, report and proceed.
            decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{{
                IECriticality: msgIe.Criticality,
                IEID:          msgIe.Id,
                TypeOfError:   TypeOfError{{Value: TypeOfErrorNotUnderstood}},
            }})
        case CriticalityIgnore:
            // Ignore and proceed.
        }}"""
    case_blocks.append(default_case)

    switch_body = "\n".join(case_blocks)
    unused_var_logic = ""
    if len(case_blocks) <= 1: 
        unused_var_logic = "\n\t_ = ieR\n\t_ = msg"

    return f"""
type {decoder_name} struct {{
	msg      *{go_name}
	diagList []CriticalityDiagnosticsIEItem
	list     map[ProtocolIEID]*E1APMessageIE
}}

func (decoder *{decoder_name}) decodeIE(r *aper.AperReader) (msgIe *E1APMessageIE, err error) {{
	var id int64
	var c uint64
	var buf []byte
	if id, err = r.ReadInteger(&aper.Constraint{{Lb: 0, Ub: 65535}}, false); err != nil {{
		return nil, err
	}}
	msgIe = new(E1APMessageIE)
	msgIe.Id = ProtocolIEID{{Value: aper.Integer(id)}}
	if c, err = r.ReadEnumerate(aper.Constraint{{Lb: 0, Ub: 2}}, false); err != nil {{
		return nil, err
	}}
	msgIe.Criticality = Criticality{{Value: aper.Enumerated(c)}}

	if buf, err = r.ReadOpenType(); err != nil {{
		return nil, err
	}}

	ieId := msgIe.Id
	if _, ok := decoder.list[ieId]; ok {{
		return nil, fmt.Errorf("duplicated protocol IE ID %d", ieId.Value)
	}}
	decoder.list[ieId] = msgIe

	ieR := aper.NewReader(bytes.NewReader(buf))
    msg := decoder.msg
    {unused_var_logic}
    switch msgIe.Id.Value {{
	{switch_body}
	}}
	return msgIe, nil // Return the populated msgIe and a nil error
}}""", {"bytes"}


    return f"{encode_code}\n{decode_code}", required_imports

def _generate_internal_decode_call(field_name, go_type, ie, asn1_def, parser):
    """
    Generates the Go code for decoding a single field inside a SEQUENCE by
    delegating the call to the field's own Decode method.
    """
    is_optional = ie.presence in ["optional", "conditional"]

    
    decode_statement = f'if err = s.{field_name}.Decode(r); err != nil {{ return fmt.Errorf("Decode {field_name} failed: %w", err) }}'
    if is_optional:
         return f's.{field_name} = new({go_type});\n\t\t{decode_statement}'
    return decode_statement