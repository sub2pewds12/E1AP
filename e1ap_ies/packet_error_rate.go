package e1ap_ies

// PacketErrorRate represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1525
type PacketErrorRate struct {
	PERScalar   int64 `asn1:"lb:0,ub:9,mandatory,ext"`
	PERExponent int64 `asn1:"lb:0,ub:9,mandatory,ext"`
}
