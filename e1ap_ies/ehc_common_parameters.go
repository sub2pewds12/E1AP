package e1ap_ies

// EHCCommonParameters From: 9_4_5_Information_Element_Definitions.txt:977
type EHCCommonParameters struct {
	EhcCIDLength EHCCommonParametersEhcCIDLength `asn1:"mandatory,ext"`
}
