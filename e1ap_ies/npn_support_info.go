package e1ap_ies

// NPNSupportInfo represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1453
// NPNSupportInfo represents a CHOICE type.
type NPNSupportInfo struct {
	SNPN NPNSupportInfoSNPN `asn1:"mandatory"`
}
