package e1ap_ies

// EHCCommonParameters From: 9_4_5_Information_Element_Definitions.txt:977
// ASN.1 Data Type: SEQUENCE
type EHCCommonParameters struct {
	EhcCIDLength EHCCommonParametersEhcCIDLength `aper:"mandatory,ext"`
}
