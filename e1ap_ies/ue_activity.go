package e1ap_ies

// UEActivity represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2358
type UEActivity int32

const (
	UEActivity_Active    UEActivity = 0
	UEActivity_NotActive UEActivity = 1
)
