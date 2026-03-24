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
    required_imports = {"io", "fmt", "bytes", "asn1go/per"}

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
        
        is_named = ie.type not in ["INTEGER", "OCTET STRING", "BIT STRING", "ENUMERATED", "OBJECT IDENTIFIER", "PrintableString", "VisibleString"]
        is_simple = isinstance(asn1_def, (IntegerDefinition, EnumDefinition, StringDefinition)) or \
                    (isinstance(asn1_def, BuiltinDefinition) and "STRING" in asn1_def.name.upper())

        if is_named and is_simple:
            value_code = f"msg.{field_name}"
            if not is_optional:
                value_code = f"&{value_code}"
        
        elif isinstance(asn1_def, IntegerDefinition) and not isinstance(asn1_def, StringDefinition):
            
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

        elif isinstance(asn1_def, ListDefinition) and False: # Force use of generated Encode method
            tmp_var_name = f"tmp{field_name}"
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
                        c:     per.IntegerConstraints{{Min: int64Ptr({getattr(of_type_def, 'min_val', 0) or 0}), Max: int64Ptr({getattr(of_type_def, 'max_val', 0) or 0})}},
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
            {tmp_var_name} := NewSequence[IE](nil, per.SizeConstraints{{
                Extensible: {str(is_extensible).lower()},
                Min: int64Ptr({min_val_str}),
                Max: int64Ptr({max_val_str}),
            }})
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
			    ID:          ProtocolIEID{{Value: {ie_id_const}}},
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
	return ies, nil
}}"""
    return final_func


def _generate_pdu_encode(
    go_name: str,
    item: SequenceDefinition,
    message_to_procedure_map: dict,
    procedures: dict,
    parser: ASN1Parser = None
) -> str:
    """Generates the Encode() dispatcher method for a PDU."""
    
    if not message_to_procedure_map:
        return f"// Encode function for {go_name} to be generated here.\\n"

    asn1_msg_name = item.name
    proc_code_const = "ProcedureCode{Value: ProcedureCodeUNKNOWN}"
    
    if asn1_msg_name in message_to_procedure_map:
        proc_name = message_to_procedure_map[asn1_msg_name]
        proc = procedures.get(proc_name)
        if proc:
            proc_code_name = getattr(proc, 'code', 'ProcedureCodeUNKNOWN')
            if not isinstance(proc_code_name, str):
                 proc_code_name = str(proc_code_name)
            if proc_code_name.isdigit():
                 proc_code_name = str(proc_code_name)
            else:
                 proc_code_name = parser.pascal_case_converter(proc_code_name)
            if proc_code_name.startswith("Id"): proc_code_name = proc_code_name[2:]
            proc_code_const = f"ProcedureCode{proc_code_name}"
        else:
            proc_code_const = "ProcedureCodeUNKNOWN"
    else:
        proc_code_const = "ProcedureCodeUNKNOWN"

    message_type = message_to_procedure_map.get(asn1_msg_name, "")
    pdu_choice_const = ""
    
    if go_name.endswith("Request") or go_name.endswith("Indication") or go_name.endswith("Notification") or go_name.startswith("Error"):
        message_type = "InitiatingMessage"
    elif go_name.endswith("Response") or go_name.endswith("Acknowledge") or go_name.endswith("Confirm") or go_name.endswith("Complete"):
         message_type = "SuccessfulOutcome"
    elif go_name.endswith("Failure"):
         message_type = "UnsuccessfulOutcome"
    
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
// Encode implements the MessageEncoder interface for {go_name}.
func (msg *{go_name}) Encode(w *per.Encoder) error {{
    ies, err := msg.toIes()
    if err != nil {{
        return fmt.Errorf("could not convert {go_name} to IEs: %w", err)
    }}

    return encodeMessage(w, {pdu_choice_const}, ProcedureCode{{Value: {proc_code_const}}}, Criticality{{Value: {pdu_crit_const}}}, ies)
}}"""



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
    required_imports = {"fmt", "math", "asn1go/per"}
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
func (s *{go_name}) Encode(w *per.Encoder) (err error) {{
    {encode_body}
    return nil
}}
"""
    decode_func = f"""
// Decode implements the aper.AperUnmarshaller interface.
func (s *{go_name}) Decode(r *per.Decoder) (err error) {{
    {decode_body}
    return nil
}}
"""
    return f"{encode_func}\n{decode_func}", required_imports


def render_enum_methods(go_name: str, item: EnumDefinition) -> Tuple[str, Set[str]]:
    """
    Generates Encode and Decode methods for an ENUMERATED type.
    """
    required_imports = {"fmt", "asn1go/per"} 

    
    num_enums = len(item.enum_values)
    upper_bound = num_enums - 1 if num_enums > 0 else 0
    is_extensible = str(item.is_extensible).lower()

    
    encode_body = f"""
    c := per.EnumeratedConstraints{{ Extensible: {is_extensible}, RootValues: make([]int64, {num_enums}), ExtValues: nil }}
    return w.EncodeEnumerated(int64(e.Value), c)"""

    encode_func = f"""
// Encode implements the MessageEncoder interface for {go_name}.
func (e *{go_name}) Encode(w *per.Encoder) error {{
    {encode_body}
}}"""

    
    decode_body = f"""
    c := per.EnumeratedConstraints{{ Extensible: {is_extensible}, RootValues: make([]int64, {num_enums}), ExtValues: nil }}
    val, err := r.DecodeEnumerated(c)
	if err != nil {{
		return err
	}}
	e.Value = val
	return nil"""

    decode_func = f"""
// Decode implements the MessageDecoder interface for {go_name}.
func (e *{go_name}) Decode(r *per.Decoder) error {{
    {decode_body}
}}"""

    return f"{encode_func}\n\n{decode_func}", required_imports



def render_extension_methods(go_name: str, extension_set: list, parser: ASN1Parser) -> Tuple[str, Set[str]]:
    required_imports = {"io", "fmt", "bytes", "asn1go/per"}
    pascal_case_converter = parser.pascal_case_converter

    encode_checks = []
    for ext in extension_set:
        field_name = pascal_case_converter(ext['id'])
        type_name = pascal_case_converter(ext['type'])
        
        id_const = f"ProtocolIEID{field_name}"
        crit_const = f"Criticality{pascal_case_converter(ext['crit'])}"
        
        encode_checks.append(f"""
    if s.{field_name} != nil {{
        extensions = append(extensions, &ProtocolExtensionField{{
            ID:          ProtocolIEID{{Value: {id_const}}},
            Criticality: Criticality{{Value: {crit_const}}},
            ExtensionValue: s.{field_name},
        }})
    }}""")

    encode_body = "\n".join(encode_checks)
    
    encode_func = f"""
func (s *{go_name}) Encode(w *per.Encoder) error {{
    var extensions []*ProtocolExtensionField
    {encode_body}

    if len(extensions) > 0 {{
        c := per.SizeConstraints{{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(MaxProtocolExtensions)}}
        if err := w.EncodeLengthDeterminant(int64(len(extensions)), c); err != nil {{
            return fmt.Errorf("encode extension container length failed: %w", err)
        }}
        for _, ext := range extensions {{
            if err := ext.Encode(w); err != nil {{
                return fmt.Errorf("encode extension failed: %w", err)
            }}
        }}
    }} else {{
        // empty extension container
        c := per.SizeConstraints{{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(MaxProtocolExtensions)}}
        if err := w.EncodeLengthDeterminant(0, c); err != nil {{
            return err
        }}
    }}
    return nil
}}"""

    decode_cases = []
    for ext in extension_set:
        field_name = pascal_case_converter(ext['id'])
        type_name = pascal_case_converter(ext['type'])
        id_const = f"ProtocolIEID{field_name}"
        
        decode_cases.append(f"""
        case {id_const}:
            s.{field_name} = new({type_name})
            if err := s.{field_name}.Decode(per.NewDecoder(ext.ValueBytes, per.APER)); err != nil {{
                return fmt.Errorf("decode extension {field_name} failed: %w", err)
            }}""")

    decode_switch = "\n".join(decode_cases)

    decode_code = f"""
func (s *{go_name}) Decode(r *per.Decoder) error {{
    c := per.SizeConstraints{{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(MaxProtocolExtensions)}}
    length, err := r.DecodeLengthDeterminant(c)
    if err != nil {{
        return fmt.Errorf("decode extension container length failed: %w", err)
    }}

    extensions := make([]*ProtocolExtensionField, length)
    for i := int64(0); i < length; i++ {{
        ext := new(ProtocolExtensionField)
        if err := ext.Decode(r); err != nil {{
            return fmt.Errorf("decode extension failed: %w", err)
        }}
        extensions[i] = ext
    }}

    for _, ext := range extensions {{
        switch ext.ID.Value {{
{decode_switch}
        default:
            // Unknown extension, ignore
        }}
    }}
    return nil
}}"""

    return f"{encode_func}\n{decode_code}", required_imports


def render_list_methods(
    go_name: str, item: ListDefinition, parser: ASN1Parser
) -> Tuple[str, Set[str]]:
    required_imports = {"fmt", "asn1go/per"}

    of_type_go_name = parser.pascal_case_converter(item.of_type)
    pascal_case_converter = parser.pascal_case_converter

    min_val_str = item.min_val if item.min_val is not None else "0"
    if not str(min_val_str).isdigit():
        min_val_str = pascal_case_converter(min_val_str)

    max_val_str = item.max_val if item.max_val is not None else "0"
    max_val_str = str(max_val_str).strip()
    
    if max_val_str == "18446744073709551615" or max_val_str == "4294967295":
        max_str = "nil"
        c_code = f"per.SizeConstraints{{Extensible: {str(item.is_extensible).lower()}, Min: int64Ptr({min_val_str}), Max: nil}}"
    elif not str(max_val_str).isdigit():
        max_val_str = pascal_case_converter(max_val_str)
        c_code = f"per.SizeConstraints{{Extensible: {str(item.is_extensible).lower()}, Min: int64Ptr({min_val_str}), Max: int64Ptr({max_val_str})}}"
    else:
        c_code = f"per.SizeConstraints{{Extensible: {str(item.is_extensible).lower()}, Min: int64Ptr({min_val_str}), Max: int64Ptr({max_val_str})}}"

    encode_body = f"""
    c := {c_code}
    if err = w.EncodeLengthDeterminant(int64(len(s.Value)), c); err != nil {{
        return fmt.Errorf("encode length determinant failed: %w", err)
    }}
    for i := 0; i < len(s.Value); i++ {{
        if err = s.Value[i].Encode(w); err != nil {{
            return fmt.Errorf("encode list item %d failed: %w", i, err)
        }}
    }}"""

    decode_body = f"""
    c := {c_code}
    length, err := r.DecodeLengthDeterminant(c)
    if err != nil {{
        return fmt.Errorf("decode length determinant failed: %w", err)
    }}
    s.Value = make([]{of_type_go_name}, length)
    for i := int64(0); i < length; i++ {{
        item := new({of_type_go_name})
        if err := item.Decode(r); err != nil {{
            return fmt.Errorf("decode list item %d failed: %w", i, err)
        }}
        s.Value[i] = *item
    }}"""

    encode_func = f"""
func (s *{go_name}) Encode(w *per.Encoder) (err error) {{
    {encode_body}
    return nil
}}
"""
    decode_func = f"""
func (s *{go_name}) Decode(r *per.Decoder) (err error) {{
    {decode_body}
    return nil
}}
"""
    return f"{encode_func}\n{decode_func}", required_imports


def _generate_direct_encode_call(
    field_name, ie, asn1_def, parser, is_struct_extensible
):
    pascal_case_converter = parser.pascal_case_converter
    is_optional = ie.presence in ["optional", "conditional"]
    is_extensible = getattr(asn1_def, "is_extensible", False) if asn1_def else False

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
        if max_val_str == "18446744073709551615" or max_val_str == "4294967295":
             constraint = "per.Unconstrained()"
        elif not max_val_str.isdigit(): 
            max_val = pascal_case_converter(max_val_str)
            if max_val and max_val[0].islower(): max_val = max_val[0].upper() + max_val[1:]
            constraint = f"per.Constrained({min_val}, {max_val})"
        else:
            constraint = f"per.Constrained({min_val}, {max_val})"

        if is_extensible:
            constraint = constraint.replace("Constrained", "ConstrainedExtensible")

        return f'if err = w.EncodeInteger(int64({value_accessor}.Value), {constraint}); err != nil {{ return fmt.Errorf("encode {field_name} failed: %w", err) }}'
    elif isinstance(asn1_def, EnumDefinition):
        if is_optional:
            value_accessor = f"(*{value_accessor})"

        num_enums = len(asn1_def.enum_values)
        
        return f"""
        {{
            enumC := per.EnumeratedConstraints{{ Extensible: {str(is_extensible).lower()}, RootValues: make([]int64, {num_enums}), ExtValues: nil }}
            if err = w.EncodeEnumerated(int64({value_accessor}.Value), enumC); err != nil {{ return fmt.Errorf("encode {field_name} failed: %w", err) }}
        }}"""
    elif isinstance(asn1_def, StringDefinition):
        if is_optional:
            value_accessor = f"(*{value_accessor})"
            
        min_val = asn1_def.min_val or 0
        if not str(min_val).isdigit(): min_val = pascal_case_converter(str(min_val))
        max_val = asn1_def.max_val or 0
        max_val_str = str(max_val).strip()
        
        ext = "true" if is_extensible else "false"
        max_str = f"int64Ptr({max_val})" if max_val_str.isdigit() else f"int64Ptr({pascal_case_converter(max_val_str)})"
        if not max_val_str.isdigit() and max_str != "nil" and max_str[9].islower():
             max_str = max_str[:9] + max_str[9].upper() + max_str[10:]
             
        constraint = f"per.SizeConstraints{{Extensible: {ext}, Min: int64Ptr({min_val}), Max: {max_str}}}"

        if asn1_def.string_type == "BIT STRING" or "BIT STRING" in getattr(asn1_def, 'alias_of', ''):
             return f'if err = w.EncodeBitString({value_accessor}.Value, {constraint}); err != nil {{ return fmt.Errorf("encode {field_name} failed: %w", err) }}'
        else:
             return f'if err = w.EncodeOctetString([]byte({value_accessor}.Value), {constraint}); err != nil {{ return fmt.Errorf("encode {field_name} failed: %w", err) }}'

    else:
        return f'if err = s.{field_name}.Encode(w); err != nil {{ return fmt.Errorf("encode {field_name} failed: %w", err) }}'

def _generate_direct_decode_call(
    field_name, go_type, ie, asn1_def, parser, is_struct_extensible
):
    pascal_case_converter = parser.pascal_case_converter
    is_optional = ie.presence in ["optional", "conditional"]
    is_extensible = getattr(asn1_def, "is_extensible", False) if asn1_def else False
    
    prefix = ""
    if is_optional:
        prefix = f"s.{field_name} = new({pascal_case_converter(ie.type)})\n\t\t"

    if isinstance(asn1_def, IntegerDefinition):
        min_val = asn1_def.min_val if asn1_def.min_val is not None else 0
        if not str(min_val).isdigit(): 
            min_val = pascal_case_converter(str(min_val))
            if min_val and min_val[0].islower(): min_val = min_val[0].upper() + min_val[1:]
        max_val = asn1_def.max_val if asn1_def.max_val is not None else 0
        max_val_str = str(max_val).strip()
        if max_val_str == "18446744073709551615" or max_val_str == "4294967295":
             constraint = "per.Unconstrained()"
        elif not max_val_str.isdigit(): 
            max_val = pascal_case_converter(max_val_str)
            if max_val and max_val[0].islower(): max_val = max_val[0].upper() + max_val[1:]
            constraint = f"per.Constrained({min_val}, {max_val})"
        else:
            constraint = f"per.Constrained({min_val}, {max_val})"

        if is_extensible: constraint = constraint.replace("Constrained", "ConstrainedExtensible")
        
        decode_log = f"""
        {{
            val, err := r.DecodeInteger({constraint})
            if err != nil {{ return fmt.Errorf("decode {field_name} failed: %w", err) }}
            s.{field_name}.Value = val
        }}"""
        if is_optional:
            return f"{prefix}{decode_log}"
        else:
            return decode_log
            
    elif isinstance(asn1_def, EnumDefinition):
        num_enums = len(asn1_def.enum_values)
        decode_log = f"""
        {{
            enumC := per.EnumeratedConstraints{{ Extensible: {str(is_extensible).lower()}, RootValues: make([]int64, {num_enums}), ExtValues: nil }}
            val, err := r.DecodeEnumerated(enumC)
            if err != nil {{ return fmt.Errorf("decode {field_name} failed: %w", err) }}
            s.{field_name}.Value = val
        }}"""
        
        if is_optional:
            return f"{prefix}{decode_log}"
        else:
            return decode_log
            
    elif isinstance(asn1_def, StringDefinition):
        min_val = asn1_def.min_val if asn1_def.min_val is not None else 0
        if not str(min_val).isdigit(): min_val = pascal_case_converter(str(min_val))
        max_val = asn1_def.max_val if asn1_def.max_val is not None else 0
        max_val_str = str(max_val).strip()
        
        ext = "true" if is_extensible else "false"
        max_str = f"int64Ptr({max_val})" if max_val_str.isdigit() else f"int64Ptr({pascal_case_converter(max_val_str)})"
        if not max_val_str.isdigit() and max_str != "nil" and max_str[9].islower():
             max_str = max_str[:9] + max_str[9].upper() + max_str[10:]
             
        constraint = f"per.SizeConstraints{{Extensible: {ext}, Min: int64Ptr({min_val}), Max: {max_str}}}"

        if asn1_def.string_type == "BIT STRING" or "BIT STRING" in getattr(asn1_def, 'alias_of', ''):
            decode_log = f"""
        {{
            val, err := r.DecodeBitString({constraint})
            if err != nil {{ return fmt.Errorf("decode {field_name} failed: %w", err) }}
            s.{field_name}.Value = val
        }}"""
            if is_optional:
                return f"{prefix}{decode_log}"
            else:
                return decode_log
        else:
            decode_log = f"""
        {{
            val, err := r.DecodeOctetString({constraint})
            if err != nil {{ return fmt.Errorf("decode {field_name} failed: %w", err) }}
            s.{field_name}.Value = val
        }}"""
            if is_optional:
                return f"{prefix}{decode_log}"
            else:
                return decode_log

    else:
        decode_log = f'if err = s.{field_name}.Decode(r); err != nil {{ return fmt.Errorf("Decode {field_name} failed: %w", err) }}'
        if is_optional:
            return f"{prefix}{decode_log}"
        else:
            return decode_log


def _generate_sequence_encode_body(item: SequenceDefinition, parser: ASN1Parser) -> str:
    pascal_case_converter = parser.pascal_case_converter
    is_extensible = str(getattr(item, "is_extensible", False)).lower()
    
    root_components_go = []
    
    for ie in item.ies:
        is_optional = str(ie.presence in ["optional", "conditional"]).lower()
        root_components_go.append(f'per.ComponentInfo{{Name: "{ie.ie}", Optional: {is_optional}}},')
        
    components_code = "\n\t\t\t".join(root_components_go)
    
    encode_parts = [f"""
    c := per.SequenceConstraints{{
        Extensible: {is_extensible},
        RootComponents: []per.ComponentInfo{{
            {components_code}
        }},
    }}
    seqEncoder := w.NewSequenceEncoder(c)
    if err := seqEncoder.EncodeExtensionBit(false); err != nil {{
        return err
    }}
    
    optionalBitmap := make([]bool, 0)
"""]

    for ie in item.ies:
        if ie.presence in ["optional", "conditional"]:
            field_name = pascal_case_converter(ie.ie)
            encode_parts.append(f"""
    if s.{field_name} != nil {{
        optionalBitmap = append(optionalBitmap, true)
    }} else {{
        optionalBitmap = append(optionalBitmap, false)
    }}""")

    encode_parts.append("""
    if err := seqEncoder.EncodePreamble(optionalBitmap); err != nil {
        return err
    }
""")

    for ie in item.ies:
        field_name = pascal_case_converter(ie.ie)
        is_optional = ie.presence in ["optional", "conditional"]
        _, asn1_def = parser.go_type_resolver(ie.type)
        
        encode_call = _generate_direct_encode_call(field_name, ie, asn1_def, parser, False)
        
        if is_optional:
            encode_parts.append(f"""
    if s.{field_name} != nil {{
        {encode_call}
    }}""")
        else:
            encode_parts.append(f"    {encode_call}")

    if item.is_extensible:
         encode_parts.append(f"""
    if err := seqEncoder.EncodeExtensionAdditions([]bool{{}}, [][]byte{{}}); err != nil {{
        return err
    }}
""")

    return "\n".join(encode_parts)

def _generate_sequence_decode_body(
    go_name: str, item: SequenceDefinition, parser: ASN1Parser
    ) -> str:
    pascal_case_converter = parser.pascal_case_converter
    is_extensible = str(getattr(item, "is_extensible", False)).lower()
    
    root_components_go = []
    for ie in item.ies:
        is_optional = str(ie.presence in ["optional", "conditional"]).lower()
        root_components_go.append(f'per.ComponentInfo{{Name: "{ie.ie}", Optional: {is_optional}}},')
        
    components_code = "\n\t\t\t".join(root_components_go)
    
    decode_parts = [f"""
    c := per.SequenceConstraints{{
        Extensible: {is_extensible},
        RootComponents: []per.ComponentInfo{{
            {components_code}
        }},
    }}
    seqDecoder := r.NewSequenceDecoder(c)
    if err := seqDecoder.DecodeExtensionBit(); err != nil {{
        return err
    }}
    
    if err := seqDecoder.DecodePreamble(); err != nil {{
        return err
    }}
"""]

    for i, ie in enumerate(item.ies):
        field_name = pascal_case_converter(ie.ie)
        is_optional = ie.presence in ["optional", "conditional"]
        go_type, asn1_def = parser.go_type_resolver(ie.type)
        if not asn1_def and ie.type == "ProtocolExtensionContainer" and hasattr(ie, "extension_set_name"):
            go_type = f"{go_name}Extensions"
        
        decode_call = _generate_direct_decode_call(field_name, go_type, ie, asn1_def, parser, False)
        
        if is_optional:
            decode_parts.append(f"""
    if seqDecoder.IsComponentPresent({i}) {{
        {decode_call}
    }}""")
        else:
            decode_parts.append(f"    {decode_call}")

    if item.is_extensible:
         decode_parts.append(f"""
    if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {{
        return err
    }}
""")

    return "\n".join(decode_parts)



def _generate_choice_encode_body(
    go_name: str, item: ChoiceDefinition, parser: ASN1Parser
) -> str:
    pascal_case_converter = parser.pascal_case_converter

    num_choices = len(item.ies)
    ub_choices = num_choices - 1 if num_choices > 0 else 0
    is_extensible = str(item.is_extensible).lower()

    # Generate choice constraints
    root_alts_go = []
    for ie in item.ies:
        root_alts_go.append(f'per.AlternativeInfo{{Name: "{ie.ie}", Tag: 0}},')

    alts_code = "\n\t\t\t".join(root_alts_go)

    encode_parts = [f"""
    c := per.ChoiceConstraints{{
        Extensible: {is_extensible},
        RootAlternatives: []per.AlternativeInfo{{
            {alts_code}
        }},
    }}
    choiceEncoder := w.NewChoiceEncoder(c)
"""]

    encode_switch_cases = []
    for i, choice_ie in enumerate(item.ies):
        choice_name = pascal_case_converter(choice_ie.ie)
        const_name = f"{go_name}Present{choice_name}"
        
        base_go_type, concrete_def = parser.go_type_resolver(choice_ie.type)
        
        encode_call = ""
        
        if base_go_type == "string":
            encode_call = f"""
        if err = choiceEncoder.EncodeChoice({i}, false, nil); err != nil {{ return err }}
        if err = w.EncodeOctetString([]byte(*s.{choice_name}), per.SizeConstraints{{Extensible: false, Min: nil, Max: nil}}); err != nil {{
            return fmt.Errorf("encode {choice_name} failed: %w", err)
        }}"""
        elif isinstance(concrete_def, IntegerDefinition):
            min_val = parser.pascal_case_converter(str(getattr(concrete_def, 'min_val', '0')))
            if not str(min_val).isdigit() and min_val and min_val[0].islower(): min_val = min_val[0].upper() + min_val[1:]
            
            max_val = parser.pascal_case_converter(str(getattr(concrete_def, 'max_val', '0')))
            max_val_str = str(max_val).strip()
            
            if max_val_str == "18446744073709551615" or max_val_str == "4294967295":
                 constraint = "per.Unconstrained()"
            elif not max_val_str.isdigit(): 
                max_val = parser.pascal_case_converter(max_val_str)
                if max_val and max_val[0].islower(): max_val = max_val[0].upper() + max_val[1:]
                constraint = f"per.Constrained({min_val}, {max_val})"
            else:
                constraint = f"per.Constrained({min_val}, {max_val})"

            choice_is_extensible = getattr(concrete_def, 'is_extensible', False)
            if choice_is_extensible:
                constraint = constraint.replace("Constrained", "ConstrainedExtensible")
                
            encode_call = f"""
        if err = choiceEncoder.EncodeChoice({i}, false, nil); err != nil {{ return err }}
        if err = w.EncodeInteger(int64(s.{choice_name}.Value), {constraint}); err != nil {{
            return fmt.Errorf("encode {choice_name} failed: %w", err)
        }}"""
        else:
            encode_call = f"""
        if err = choiceEncoder.EncodeChoice({i}, false, nil); err != nil {{ return err }}
        if err = s.{choice_name}.Encode(w); err != nil {{
            return fmt.Errorf("encode {choice_name} failed: %w", err)
        }}"""
        
        encode_switch_cases.append(f"case {const_name}:{encode_call}")

    encode_switch_body = "\n    ".join(encode_switch_cases)

    encode_parts.append(f"""
    switch s.Choice {{
    {encode_switch_body}
    default:
        return fmt.Errorf("Encode choice of {go_name} with unknown choice value %d", s.Choice)
    }}""")

    return "".join(encode_parts)

def _generate_choice_decode_body(
    go_name: str, item: ChoiceDefinition, parser: ASN1Parser
) -> str:
    pascal_case_converter = parser.pascal_case_converter

    num_choices = len(item.ies)
    ub_choices = num_choices - 1 if num_choices > 0 else 0
    is_extensible = str(item.is_extensible).lower()

    root_alts_go = []
    for ie in item.ies:
        root_alts_go.append(f'per.AlternativeInfo{{Name: "{ie.ie}", Tag: 0}},')

    alts_code = "\n\t\t\t".join(root_alts_go)

    decode_parts = [f"""
    c := per.ChoiceConstraints{{
        Extensible: {is_extensible},
        RootAlternatives: []per.AlternativeInfo{{
            {alts_code}
        }},
    }}
    choiceDecoder := r.NewChoiceDecoder(c)
    
    choiceIndex, isExtension, _, err := choiceDecoder.DecodeChoice()
    if err != nil {{
        return fmt.Errorf("decode choice index failed: %w", err)
    }}
    
    if isExtension {{
        return fmt.Errorf("extension choices are not fully supported yet")
    }}
    
    s.Choice = uint64(choiceIndex + 1) // 1-based internal Choice enum
"""]

    decode_switch_cases = []
    for i, choice_ie in enumerate(item.ies):
        choice_name = pascal_case_converter(choice_ie.ie)
        base_go_type, concrete_def = parser.go_type_resolver(choice_ie.type)
        
        alloc_and_decode = ""
        
        if base_go_type == "string":
            alloc_and_decode = f"""
        val, err := r.DecodeOctetString(per.SizeConstraints{{Extensible: false, Min: nil, Max: nil}})
        if err != nil {{ return fmt.Errorf("decode {choice_name} failed: %w", err) }}
        tmpStr := string(val)
        s.{choice_name} = &tmpStr"""
        elif isinstance(concrete_def, IntegerDefinition):
            min_val = parser.pascal_case_converter(str(getattr(concrete_def, 'min_val', '0')))
            if not str(min_val).isdigit() and min_val and min_val[0].islower(): min_val = min_val[0].upper() + min_val[1:]
            
            max_val = parser.pascal_case_converter(str(getattr(concrete_def, 'max_val', '0')))
            max_val_str = str(max_val).strip()
            
            if max_val_str == "18446744073709551615" or max_val_str == "4294967295":
                 constraint = "per.Unconstrained()"
            elif not max_val_str.isdigit(): 
                max_val = parser.pascal_case_converter(max_val_str)
                if max_val and max_val[0].islower(): max_val = max_val[0].upper() + max_val[1:]
                constraint = f"per.Constrained({min_val}, {max_val})"
            else:
                constraint = f"per.Constrained({min_val}, {max_val})"

            choice_is_extensible = getattr(concrete_def, 'is_extensible', False)
            if choice_is_extensible: constraint = constraint.replace("Constrained", "ConstrainedExtensible")
                
            alloc_and_decode = f"""
        s.{choice_name} = new({base_go_type})
        val, err := r.DecodeInteger({constraint})
        if err != nil {{ return fmt.Errorf("decode {choice_name} failed: %w", err) }}
        s.{choice_name}.Value = val"""
        else:
            alloc_and_decode = f"""
        s.{choice_name} = new({base_go_type})
        if err = s.{choice_name}.Decode(r); err != nil {{ return fmt.Errorf("decode {choice_name} failed: %w", err) }}"""
        
        decode_switch_cases.append(f"case {i}:{alloc_and_decode}")
        
    decode_switch_body = "\n    ".join(decode_switch_cases)

    decode_parts.append(f"""
    switch choiceIndex {{
    {decode_switch_body}
    default:
        return fmt.Errorf("decode choice of {go_name} with unknown choice index %d", choiceIndex)
    }}""")

    return "".join(decode_parts)


def _generate_pdu_encode(go_name: str, item: SequenceDefinition, message_to_procedure_map: dict, procedures: dict, parser: ASN1Parser) -> str:
    # Existing internal encode
    internal_encode = f"""
func (msg *{go_name}) EncodeWithEncoder(e *per.Encoder) (err error) {{
	ies, err := msg.toIes()
	if err != nil {{
		return err
	}}

	sizeC := per.SizeConstraints{{Extensible: false, Min: int64Ptr(0), Max: int64Ptr(65535)}}
	if err = e.EncodeLengthDeterminant(int64(len(ies)), sizeC); err != nil {{
		return fmt.Errorf("encode IE count failed: %w", err)
	}}
	for i := range ies {{
		if err = ies[i].Encode(e); err != nil {{
			return fmt.Errorf("encode IE %d failed: %w", i, err)
		}}
	}}
	return nil
}}"""

    # New public encode
    public_encode = f"""
func (msg *{go_name}) Encode(w io.Writer) error {{
	e := per.NewEncoder(per.APER)
	if err := msg.EncodeWithEncoder(e); err != nil {{
		return err
	}}
	_, err := w.Write(e.Bytes())
	return err
}}"""

    return internal_encode + "\n" + public_encode





def _generate_pdu_decode(go_name: str, item: SequenceDefinition, parser: ASN1Parser) -> Tuple[str, Set[str]]:
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
		if err == nil {{
			err = fmt.Errorf("mandatory field {field_name} is missing")
		}}
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
// Decode implements the MessageUnmarshaller interface for {go_name}.
func (msg *{go_name}) Decode(data []byte) (diagList []CriticalityDiagnosticsIEItem, err error) {{
	r := per.NewDecoder(data, per.APER)
	return msg.DecodeFromDecoder(r)
}}

func (msg *{go_name}) DecodeFromDecoder(r *per.Decoder) (diagList []CriticalityDiagnosticsIEItem, err error) {{

	defer func() {{
		if err != nil {{
			err = fmt.Errorf("decode {go_name} failed: %w", err)
		}}
	}}()

	decoder := {decoder_name}{{
		msg:  msg,
		list: make(map[ProtocolIEID]*E1APMessageIE),
	}}
	
	c := per.SizeConstraints{{Extensible: false, Min: int64Ptr(0), Max: int64Ptr(65535)}}
	length, err := r.DecodeLengthDeterminant(c)
	if err != nil {{
		return
	}}

	for i := int64(0); i < length; i++ {{
		if _, err = decoder.decodeIE(r); err != nil {{
			return
		}}
	}}

    // After decoding all present IEs, validate that mandatory ones were found.
    {validation_block}

	return
}}""", set()

def _generate_decoder_helper(go_name: str, item: SequenceDefinition, parser: ASN1Parser) -> Tuple[str, Set[str]]:
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
            return nil, fmt.Errorf("not comprehended IE ID %d (criticality: reject)", msgIe.ID.Value)
        case CriticalityNotify:
            decoder.diagList = append(decoder.diagList, CriticalityDiagnosticsIEItem{{
                IECriticality: msgIe.Criticality,
                IEID:          msgIe.ID,
                TypeOfError:   TypeOfError{{Value: TypeOfErrorNotUnderstood}},
            }})
        case CriticalityIgnore:
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

func (decoder *{decoder_name}) decodeIE(r *per.Decoder) (msgIe *E1APMessageIE, err error) {{
	id, err := r.DecodeInteger(per.Constrained(0, 65535))
	if err != nil {{
		return nil, err
	}}
	msgIe = new(E1APMessageIE)
	msgIe.ID = ProtocolIEID{{Value: id}}
	
	enumC := per.EnumeratedConstraints{{Extensible: false, RootValues: make([]int64, 3)}}
	c, err := r.DecodeEnumerated(enumC)
	if err != nil {{
		return nil, err
	}}
	msgIe.Criticality = Criticality{{Value: c}}

	buf, err := r.DecodeOctetString(per.SizeConstraints{{Extensible: false, Min: nil, Max: nil}})
	if err != nil {{
		return nil, err
	}}

	ieId := msgIe.ID
	if _, ok := decoder.list[ieId]; ok {{
		return nil, fmt.Errorf("duplicated protocol IE ID %d", ieId.Value)
	}}
	decoder.list[ieId] = msgIe

	ieR := per.NewDecoder(buf, per.APER)
    msg := decoder.msg
    {unused_var_logic}
    switch msgIe.ID.Value {{
	{switch_body}
	}}
	return msgIe, nil
}}""", {"bytes"}

def _generate_pdu_decode_call(field_name, go_type, ie, asn1_def, parser):
    is_optional = ie.presence in ["optional", "conditional"]
    is_extensible = getattr(asn1_def, 'is_extensible', False) if asn1_def else False
    pascal_case_converter = parser.pascal_case_converter

    if isinstance(asn1_def, IntegerDefinition):
        min_val = pascal_case_converter(str(asn1_def.min_val or 0)) if not str(asn1_def.min_val or 0).isdigit() else asn1_def.min_val or 0
        if not str(min_val).isdigit() and min_val and min_val[0].islower(): min_val = min_val[0].upper() + min_val[1:]
        max_val_str = str(asn1_def.max_val or 0).strip()
        
        if max_val_str == "18446744073709551615" or max_val_str == "4294967295":
             constraint = "per.Unconstrained()"
        else:
            max_val = pascal_case_converter(max_val_str) if not max_val_str.isdigit() else max_val_str
            if not str(max_val).isdigit() and max_val and max_val[0].islower(): max_val = max_val[0].upper() + max_val[1:]
            constraint = f"per.Constrained({min_val}, {max_val})"

        if is_extensible:
            constraint = constraint.replace("Constrained", "ConstrainedExtensible")
        
        assignment = f"msg.{field_name}.Value = val"
        if is_optional:
            assignment = f"msg.{field_name} = new({pascal_case_converter(ie.type)}); msg.{field_name}.Value = val"
        return f"""
        {{
            val, err := ieR.DecodeInteger({constraint})
            if err != nil {{
                 return nil, fmt.Errorf("decode {field_name} failed: %w", err)
            }}
            {assignment}
        }}"""

    elif isinstance(asn1_def, EnumDefinition):
        num_enums = len(asn1_def.enum_values)
        upper_bound = num_enums - 1 if num_enums > 0 else 0
        ext_str = str(is_extensible).lower()
        decode_block = f"""
        {{
            c := per.EnumeratedConstraints{{Extensible: {ext_str}, RootValues: make([]int64, {num_enums})}}
            val, err := ieR.DecodeEnumerated(c)
            if err != nil {{
                return nil, fmt.Errorf("decode {field_name} failed: %w", err)
            }}
            msg.{field_name}.Value = val
        }}"""
        if is_optional:
            return f"msg.{field_name} = new({go_type});\n\t\t{decode_block}"
        return decode_block

    elif isinstance(asn1_def, (StringDefinition, BuiltinDefinition)) and "STRING" in asn1_def.name.upper():
        min_val = getattr(asn1_def, 'min_val', 0) or 0
        max_val = getattr(asn1_def, 'max_val', 0) or 0
        ext_str = "true" if is_extensible else "false"
        max_str = f"int64Ptr({max_val})" if str(max_val).isdigit() else "nil"
        constraint = f"per.SizeConstraints{{Extensible: {ext_str}, Min: int64Ptr({min_val}), Max: {max_str}}}"

        if "BIT STRING" in asn1_def.name.upper():
            assignment = f"msg.{field_name}.Value = val"
            if is_optional:
                assignment = f"msg.{field_name} = new({go_type}); msg.{field_name}.Value = val"
            return f"""
            {{
                val, err := ieR.DecodeBitString({constraint})
                if err != nil {{ return nil, fmt.Errorf("decode {field_name} failed: %w", err) }}
                {assignment}
            }}"""
        else: 
            assignment = f"msg.{field_name}.Value = val"
            if is_optional:
                assignment = f"msg.{field_name} = new({go_type}); msg.{field_name}.Value = val"
            return f"""
            {{
                val, err := ieR.DecodeOctetString({constraint})
                if err != nil {{ return nil, fmt.Errorf("decode {field_name} failed: %w", err) }}
                {assignment}
            }}"""

    elif isinstance(asn1_def, ListDefinition):
        of_type_go_name, _ = parser.go_type_resolver(asn1_def.of_type)
        item_type = of_type_go_name
        
        min_val = asn1_def.min_val if asn1_def.min_val is not None else 0
        if not str(min_val).isdigit(): min_val = pascal_case_converter(str(min_val))
        
        max_val = asn1_def.max_val if asn1_def.max_val is not None else 0
        if not str(max_val).isdigit(): max_val = pascal_case_converter(str(max_val))

        decode_block = f"""
        {{
            itemDecoder := func(r *per.Decoder) (*{item_type}, error) {{
                item := new({item_type})
                if err := item.Decode(r); err != nil {{
                    return nil, err
                }}
                return item, nil
            }}
            
            c := per.SizeConstraints{{Extensible: {str(is_extensible).lower()}, Min: int64Ptr({min_val}), Max: int64Ptr({max_val})}}
            length, err := ieR.DecodeLengthDeterminant(c)
            if err != nil {{ return nil, fmt.Errorf("decode struct list length failed: %w", err) }}
            for i := int64(0); i < length; i++ {{
                item, err := itemDecoder(ieR)
                if err != nil {{ return nil, fmt.Errorf("decode item failed: %w", err) }}
                msg.{field_name}.Value = append(msg.{field_name}.Value, *item)
            }}
        }}"""
        if is_optional:
             # Use the IE type name for the struct allocation
             ie_type_name = pascal_case_converter(ie.type)
             return f"msg.{field_name} = new({ie_type_name});\n\t\t{decode_block}"
        return decode_block


    else:
        decode_block = f"""
        if err = msg.{field_name}.Decode(ieR); err != nil {{
             return nil, fmt.Errorf("decode {field_name} failed: %w", err)
        }}"""
        if is_optional:
             return f"msg.{field_name} = new({go_type});\n\t\t{decode_block}"
        return decode_block
