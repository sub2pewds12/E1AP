package e1ap_ies

// GBRQosInformation From: 9_4_5_Information_Element_Definitions.txt:1165
type GBRQosInformation struct {
	ERABMaximumBitrateDL    int64 `asn1:"mandatory,ext"`
	ERABMaximumBitrateUL    int64 `asn1:"mandatory,ext"`
	ERABGuaranteedBitrateDL int64 `asn1:"mandatory,ext"`
	ERABGuaranteedBitrateUL int64 `asn1:"mandatory,ext"`
}
