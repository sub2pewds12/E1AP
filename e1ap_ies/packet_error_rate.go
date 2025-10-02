package e1ap_ies

// PacketErrorRate From: 9_4_5_Information_Element_Definitions.txt:1525
type PacketErrorRate struct {
	PERScalar   int64 `asn1:"mandatory,ext"`
	PERExponent int64 `asn1:"mandatory,ext"`
}
