package e1ap_ies

// GBRQosInformation represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1165
type GBRQosInformation struct {
	ERABMaximumBitrateDL    int64 `asn1:"lb:0,ub:4000000000000,mandatory,ext"`
	ERABMaximumBitrateUL    int64 `asn1:"lb:0,ub:4000000000000,mandatory,ext"`
	ERABGuaranteedBitrateDL int64 `asn1:"lb:0,ub:4000000000000,mandatory,ext"`
	ERABGuaranteedBitrateUL int64 `asn1:"lb:0,ub:4000000000000,mandatory,ext"`
}
