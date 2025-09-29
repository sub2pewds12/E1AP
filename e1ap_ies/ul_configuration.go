package e1ap_ies

// ULConfiguration represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:2375
type ULConfiguration int32

const (
	ULConfiguration_NoData ULConfiguration = 0
	ULConfiguration_Shared ULConfiguration = 1
	ULConfiguration_Only   ULConfiguration = 2
)
