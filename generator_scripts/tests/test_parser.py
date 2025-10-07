import os
import re
from typing import List, Tuple, Optional, Dict, Any
import pytest

from asn1_parser import ASN1Parser
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
)
import sys

sys.path.insert(0, os.path.abspath(os.path.join(os.path.dirname(__file__), "..")))


def load_test_file(filename: str) -> List[Tuple[str, str, int]]:
    """Loads a single ASN.1 sample file for testing."""
    filepath = os.path.join(os.path.dirname(__file__), "asn1_samples", filename)
    lines = []
    with open(filepath, "r") as f:
        for i, line_text in enumerate(f, 1):
            lines.append((line_text, filename, i))
    return lines


def test_parses_simple_integer():
    lines = load_test_file("simple_integer.txt")
    parser = ASN1Parser(lines)
    definitions, failures, _, _ = parser.parse()
    assert "MyInteger" in definitions
    assert not failures
    item = definitions["MyInteger"]
    assert isinstance(item, IntegerDefinition)
    assert item.name == "MyInteger"
    assert item.min_val is None
    assert item.max_val is None


def test_parses_integer_with_range():
    lines = load_test_file("integer_with_range.txt")
    parser = ASN1Parser(lines)
    definitions, failures, _, _ = parser.parse()
    assert "MyCounter" in definitions
    assert not failures
    item = definitions["MyCounter"]
    assert isinstance(item, IntegerDefinition)
    assert item.name == "MyCounter"
    assert item.min_val == "0"
    assert item.max_val == "255"


def test_parses_simple_sequence():
    lines = load_test_file("simple_sequence.txt")
    boolean_def = BuiltinDefinition("BOOLEAN", "builtin")
    parser = ASN1Parser(lines)
    parser.definitions["BOOLEAN"] = boolean_def
    definitions, failures, _, _ = parser.parse()
    assert "MyMessage" in definitions
    assert not failures
    item = definitions["MyMessage"]
    assert isinstance(item, SequenceDefinition)
    assert item.name == "MyMessage"
    assert len(item.ies) == 2
    assert item.ies[0].ie == "transactionID"
    assert item.ies[0].type == "MyMessage-transactionID"
    assert item.ies[0].presence == "mandatory"
    assert item.ies[1].ie == "isOptional"
    assert item.ies[1].type == "BOOLEAN"
    assert item.ies[1].presence == "optional"


def test_handles_malformed_definition_gracefully():
    """
    Tests that the parser doesn't crash on a known-bad definition,
    but instead adds it to the 'failures' list.
    """
    lines = load_test_file("malformed_definition.txt")
    parser = ASN1Parser(lines)
    _, failures, _, _ = parser.parse()

    assert len(failures) == 1
    assert failures[0]["name"] == "BadDef { something }"


def test_parses_simple_enumerated():
    lines = load_test_file("simple_enum.txt")
    parser = ASN1Parser(lines)
    definitions, failures, _, _ = parser.parse()
    assert "MyEnum" in definitions
    assert not failures
    item = definitions["MyEnum"]
    assert isinstance(item, EnumDefinition)
    assert item.name == "MyEnum"
    assert item.enum_values == ["value1", "value2"]


def test_parses_octet_string_with_size():
    lines = load_test_file("octet_string_with_size.txt")
    parser = ASN1Parser(lines)
    definitions, failures, _, _ = parser.parse()
    assert "TransportLayerAddress" in definitions
    assert not failures
    item = definitions["TransportLayerAddress"]
    assert isinstance(item, StringDefinition)
    assert item.string_type == "OCTET STRING"
    assert item.name == "TransportLayerAddress"
    assert item.min_val == "4"
    assert item.max_val == "4"


def test_parses_bit_string_with_size_range():
    lines = load_test_file("bit_string_with_size.txt")
    parser = ASN1Parser(lines)
    definitions, failures, _, _ = parser.parse()
    assert "DataForwardingResponseDRBList" in definitions
    assert not failures
    item = definitions["DataForwardingResponseDRBList"]
    assert isinstance(item, StringDefinition)
    assert item.string_type == "BIT STRING"
    assert item.name == "DataForwardingResponseDRBList"
    assert item.min_val == "1"
    assert item.max_val == "1024"


def test_parses_simple_choice():
    lines = load_test_file("simple_choice.txt")
    dummy_types = ["CauseRadioNetwork", "CauseTransport", "CauseProtocol", "CauseMisc"]
    dummy_defs = {name: SequenceDefinition(name, "dummy") for name in dummy_types}
    parser = ASN1Parser(lines)
    parser.definitions.update(dummy_defs)
    definitions, failures, _, _ = parser.parse()
    assert "Cause" in definitions
    assert not failures
    item = definitions["Cause"]
    assert isinstance(item, ChoiceDefinition)
    assert item.name == "Cause"
    assert len(item.ies) == 4
    assert item.ies[0].ie == "radio-network"
    assert item.ies[0].type == "CauseRadioNetwork"
    assert item.ies[3].ie == "misc"
    assert item.ies[3].type == "CauseMisc"


def test_parses_constant_integer():
    lines = load_test_file("constant_integer.txt")
    dummy_def = IntegerDefinition("ProtocolIE-ID", "dummy")
    parser = ASN1Parser(lines)
    parser.definitions["ProtocolIE-ID"] = dummy_def
    definitions, failures, _, _ = parser.parse()
    constant_name = "id-E1SetupRequest"
    assert constant_name in definitions
    assert not failures
    item = definitions[constant_name]
    assert isinstance(item, ConstantDefinition)
    assert item.name == "id-E1SetupRequest"
    assert item.const_type == "ProtocolIE-ID"
    assert item.min_val == "1"
    assert item.max_val == "1"


def test_parses_sequence_of():
    lines = load_test_file("sequence_of_integer.txt")
    dummy_def = IntegerDefinition("DRB-ID", "dummy")
    parser = ASN1Parser(lines)
    parser.definitions["DRB-ID"] = dummy_def
    definitions, failures, _, _ = parser.parse()
    assert "DRB-List" in definitions
    assert not failures
    item = definitions["DRB-List"]
    assert isinstance(item, ListDefinition)
    assert item.name == "DRB-List"
    assert item.min_val == "1"
    assert item.max_val == "1024"
    assert item.of_type == "DRB-ID"


def test_ignores_comments_in_sequence():
    lines = load_test_file("sequence_with_comment.txt")
    parser = ASN1Parser(lines)
    definitions, failures, _, _ = parser.parse()
    assert "MessageWithComment" in definitions
    assert not failures
    item = definitions["MessageWithComment"]
    assert isinstance(item, SequenceDefinition)
    assert item.name == "MessageWithComment"
    assert len(item.ies) == 2
    assert item.ies[0].ie == "first-field"
    assert item.ies[1].ie == "second-field"


def test_parses_elementary_procedure():
    """
    Tests the parsing of an E1AP-ELEMENTARY-PROCEDURE block.
    """

    lines = load_test_file("elementary_procedure.txt")

    parser = ASN1Parser(lines)
    _, failures, procedures, message_map = parser.parse()

    assert not failures
    assert "e1Setup" in procedures

    proc = procedures["e1Setup"]
    assert proc.name == "e1Setup"
    assert proc.initiating_message == "E1SetupRequest"
    assert proc.successful_outcome == "E1SetupResponse"
    assert proc.unsuccessful_outcome == "E1SetupFailure"
    assert proc.procedure_code == "id-e1Setup"

    assert message_map["E1SetupRequest"] == "InitiatingMessage"
    assert message_map["E1SetupResponse"] == "SuccessfulOutcome"
    assert message_map["E1SetupFailure"] == "UnsuccessfulOutcome"


def test_parses_boolean_as_alias():
    """
    Tests that a simple BOOLEAN definition is parsed as an alias of a built-in.
    """
    lines = load_test_file("boolean_def.txt")
    parser = ASN1Parser(lines)
    definitions, failures, _, _ = parser.parse()

    assert not failures
    assert "IsActive" in definitions
    item = definitions["IsActive"]
    assert isinstance(item, AliasDefinition)
    assert item.alias_of == "BOOLEAN"


def test_parses_sequence_with_inline_enum():
    """
    Tests that the parser correctly handles a SEQUENCE that contains
    an inline ENUMERATED definition, creating a synthetic type for it.
    """

    lines = load_test_file("sequence_with_inline_enum.txt")

    dummy_defs = {
        "DRB-ID": IntegerDefinition("DRB-ID", "dummy"),
        "maxnoofQoSFlows": ConstantDefinition("INTEGER", "maxnoofQoSFlows", "dummy"),
        "QoS-Flow-Removed-Item": SequenceDefinition("QoS-Flow-Removed-Item", "dummy"),
    }

    parser = ASN1Parser(lines)
    parser.definitions.update(dummy_defs)
    definitions, failures, _, _ = parser.parse()

    assert not failures

    assert "DRB-Removed-Item" in definitions
    main_item = definitions["DRB-Removed-Item"]
    assert isinstance(main_item, SequenceDefinition)

    assert len(main_item.ies) == 5

    enum_member = main_item.ies[1]
    assert enum_member.ie == "dRB-Released-In-Session"

    assert enum_member.type == "DRB-Removed-Item-dRB-Released-In-Session"
    assert enum_member.presence == "optional"

    synthetic_enum_name = "DRB-Removed-Item-dRB-Released-In-Session"
    assert synthetic_enum_name in definitions
    synth_item = definitions[synthetic_enum_name]
    assert isinstance(synth_item, EnumDefinition)
    assert synth_item.is_synthetic is True
    assert synth_item.enum_values == ["released-in-session", "not-released-in-session"]


def test_parses_ie_set_with_inline_integer():
    """
    Tests that the parser, when parsing an IE Set, can handle an
    inline INTEGER definition and create a synthetic type for it.
    """

    lines = load_test_file("ie_set_with_inline_integer.txt")

    parser = ASN1Parser(lines)

    definitions, failures, _, _ = parser.parse()

    assert not failures

    assert "ResourceStatusRequestIEs" in parser.ie_sets
    ie_set = parser.ie_sets["ResourceStatusRequestIEs"]
    assert len(ie_set) == 6

    measurement_ie = ie_set[1]
    assert measurement_ie["id"] == "id-gNB-CU-CP-Measurement-ID"

    assert (
        measurement_ie["type"] == "ResourceStatusRequestIEs-id-gNB-CU-CP-Measurement-ID"
    )

    synthetic_name = "ResourceStatusRequestIEs-id-gNB-CU-CP-Measurement-ID"
    assert synthetic_name in definitions
    synth_item = definitions[synthetic_name]
    assert isinstance(synth_item, IntegerDefinition)
    assert synth_item.is_synthetic is True
    assert synth_item.min_val == "1"
    assert synth_item.max_val == "4095"


def test_parses_sequence_with_secondaryrat_inline_enum():
    """
    Verifies that the existing inline enum logic correctly handles
    the secondaryRATType field.
    """

    lines = load_test_file("sequence_with_secondaryrat_enum.txt")

    dummy_defs = {
        "maxnooftimeperiods": ConstantDefinition(
            "INTEGER", "maxnooftimeperiods", "dummy"
        ),
        "MRDC-Data-Usage-Report-Item": SequenceDefinition(
            "MRDC-Data-Usage-Report-Item", "dummy"
        ),
    }

    parser = ASN1Parser(lines)
    parser.definitions.update(dummy_defs)
    definitions, failures, _, _ = parser.parse()

    assert not failures

    synthetic_enum_name = "Data-Usage-per-PDU-Session-Report-secondaryRATType"
    assert synthetic_enum_name in definitions
    synth_item = definitions[synthetic_enum_name]
    assert isinstance(synth_item, EnumDefinition)
    assert synth_item.is_synthetic is True
    assert synth_item.enum_values == ["nR", "e-UTRA"]


def test_parses_list_with_inline_sequence():
    """
    Tests that the parser correctly handles a SEQUENCE OF an inline SEQUENCE,
    creating a synthetic type for the inner sequence.
    """

    lines = load_test_file("list_with_inline_sequence.txt")

    dummy_defs = {
        "maxnoofErrors": ConstantDefinition("INTEGER", "maxnoofErrors", "dummy"),
        "Criticality": EnumDefinition("Criticality", "dummy"),
        "ProtocolIE-ID": IntegerDefinition("ProtocolIE-ID", "dummy"),
        "TypeOfError": EnumDefinition("TypeOfError", "dummy"),
    }

    parser = ASN1Parser(lines)
    parser.definitions.update(dummy_defs)
    definitions, failures, _, _ = parser.parse()

    assert not failures

    assert "CriticalityDiagnostics-IE-List" in definitions
    main_item = definitions["CriticalityDiagnostics-IE-List"]
    assert isinstance(main_item, ListDefinition)

    assert main_item.of_type == "CriticalityDiagnostics-IE-Item"
    assert main_item.max_val == "maxnoofErrors"

    assert "CriticalityDiagnostics-IE-Item" in definitions
    synth_item = definitions["CriticalityDiagnostics-IE-Item"]
    assert isinstance(synth_item, SequenceDefinition)
    assert synth_item.is_synthetic is True
    assert len(synth_item.ies) == 3
    assert synth_item.ies[0].ie == "iECriticality"
    assert synth_item.ies[0].type == "Criticality"

def test_parses_ie_extensions_field():
    """
    Tests that the parser correctly parses the 'iE-Extensions' field
    and identifies its type as 'ProtocolExtensionContainer'.
    """
    lines = load_test_file("sequence_with_inline_enum.txt")

    dummy_defs = {
        "DRB-ID": IntegerDefinition("DRB-ID", "dummy"),
        "maxnoofQoSFlows": ConstantDefinition("INTEGER", "maxnoofQoSFlows", "dummy"),
        "QoS-Flow-Removed-Item": SequenceDefinition("QoS-Flow-Removed-Item", "dummy"),
    }
    
    parser = ASN1Parser(lines)
    parser.definitions.update(dummy_defs)
    definitions, failures, _, _ = parser.parse()

    assert not failures

    main_item = definitions["DRB-Removed-Item"]
    assert isinstance(main_item, SequenceDefinition)
    assert len(main_item.ies) == 5

    extension_member = None
    for ie in main_item.ies:
        if ie.ie == "iE-Extensions":
            extension_member = ie
            break
    
    assert extension_member is not None
    assert extension_member.presence == "optional"
    assert extension_member.type == "ProtocolExtensionContainer"
