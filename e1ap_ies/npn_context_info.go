package e1ap_ies

// NPNContextInfo represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1472
// NPNContextInfo represents a CHOICE type.
type NPNContextInfo struct {
	SNPN NPNContextInfoSNPN `asn1:"mandatory"`
}
