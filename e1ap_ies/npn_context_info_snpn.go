package e1ap_ies

// NPNContextInfoSNPN From: 9_4_5_Information_Element_Definitions.txt:1481
type NPNContextInfoSNPN struct {
	NID []byte `asn1:"lb:44,ub:44,mandatory"`
}
