package e1ap_ies

// DRBActivity represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:419
type DRBActivity int32

const (
	DRBActivity_Active    DRBActivity = 0
	DRBActivity_NotActive DRBActivity = 1
)
