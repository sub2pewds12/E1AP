package e1ap_ies

// UPTNLInformation represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2421
// UPTNLInformation represents a CHOICE type.
type UPTNLInformation struct {
	GTPTunnel GTPTunnel `asn1:"mandatory"`
}
