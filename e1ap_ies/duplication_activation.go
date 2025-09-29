package e1ap_ies

// DuplicationActivation represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:933
type DuplicationActivation int32

const (
	DuplicationActivation_Active   DuplicationActivation = 0
	DuplicationActivation_Inactive DuplicationActivation = 1
)
