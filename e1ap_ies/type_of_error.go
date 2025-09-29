package e1ap_ies

// TypeOfError represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2314
type TypeOfError int32

const (
	TypeOfError_NotUnderstood TypeOfError = 0
	TypeOfError_Missing       TypeOfError = 1
)
