package e1ap_ies

// UplinkOnlyROHC From: 9_4_5_Information_Element_Definitions.txt:2430
type UplinkOnlyROHC struct {
	MaxCID       int64                       `asn1:"mandatory,ext"`
	ROHCProfiles int64                       `asn1:"mandatory,ext"`
	ContinueROHC *UplinkOnlyROHCContinueROHC `asn1:"optional,ext"`
}
