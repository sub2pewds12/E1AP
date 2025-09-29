package e1ap_ies

// CPTNLInformation represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:266
// CPTNLInformation represents a CHOICE type.
type CPTNLInformation struct {
	EndpointIPAddress []byte `asn1:"lb:1,ub:160,mandatory"`
}
