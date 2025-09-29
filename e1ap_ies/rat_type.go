package e1ap_ies

// RATType represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2059
type RATType int32

const (
	RATType_EUTRA RATType = 0
	RATType_NR    RATType = 1
)
