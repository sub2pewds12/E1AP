package e1ap_ies

// ROHCParameters represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2105
// ROHCParameters represents a CHOICE type.
type ROHCParameters struct {
	ROHC            ROHC                      `asn1:"mandatory"`
	UPlinkOnlyROHC  UplinkOnlyROHC            `asn1:"mandatory"`
	ChoiceExtension ProtocolIESingleContainer `asn1:"mandatory"`
}
