package e1ap_ies

// NPNSupportInfoSNPN represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1462
type NPNSupportInfoSNPN struct {
	NID []byte `asn1:"lb:44,ub:44,mandatory"`
}
