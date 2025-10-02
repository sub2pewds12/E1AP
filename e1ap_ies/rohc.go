package e1ap_ies

// ROHC From: 9_4_5_Information_Element_Definitions.txt:2115
type ROHC struct {
	MaxCID       int64             `asn1:"mandatory,ext"`
	ROHCProfiles int64             `asn1:"mandatory,ext"`
	ContinueROHC *ROHCContinueROHC `asn1:"optional,ext"`
}
