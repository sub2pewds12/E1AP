package e1ap_ies

// PreEmptionCapability represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1916
type PreEmptionCapability int32

const (
	PreEmptionCapability_ShallNotTriggerPreEmption PreEmptionCapability = 0
	PreEmptionCapability_MayTriggerPreEmption      PreEmptionCapability = 1
)
