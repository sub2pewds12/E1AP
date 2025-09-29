package e1ap_ies

// EarlyForwardingCOUNTInfo represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:965
// EarlyForwardingCOUNTInfo represents a CHOICE type.
type EarlyForwardingCOUNTInfo struct {
	FirstDLCount      FirstDLCount              `asn1:"mandatory"`
	DLDiscardingCount DLDiscarding              `asn1:"mandatory"`
	ChoiceExtension   ProtocolIESingleContainer `asn1:"mandatory"`
}
