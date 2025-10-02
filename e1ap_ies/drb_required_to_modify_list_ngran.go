package e1ap_ies

// DRBRequiredToModifyListNGRAN From: 9_4_5_Information_Element_Definitions.txt:605
type DRBRequiredToModifyListNGRAN []DRBRequiredToModifyItemNGRAN

// DRBRequiredToModifyItemNGRAN From: 9_4_5_Information_Element_Definitions.txt:607
type DRBRequiredToModifyItemNGRAN struct {
	DRBID                                int64                                      `asn1:"mandatory,ext"`
	GNBCUUPCellGroupRelatedConfiguration []GNBCUUPCellGroupRelatedConfigurationItem `asn1:"optional,ext"`
	FlowToRemove                         []QOSFlowItem                              `asn1:"optional,ext"`
	Cause                                *Cause                                     `asn1:"optional,ext"`
}
