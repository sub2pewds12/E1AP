package e1ap_ies

// DRBRequiredToModifyListEUTRAN From: 9_4_5_Information_Element_Definitions.txt:590
type DRBRequiredToModifyListEUTRAN []DRBRequiredToModifyItemEUTRAN

// DRBRequiredToModifyItemEUTRAN From: 9_4_5_Information_Element_Definitions.txt:592
type DRBRequiredToModifyItemEUTRAN struct {
	DRBID                                int64                                      `asn1:"mandatory,ext"`
	S1DLUPTNLInformation                 *UPTNLInformation                          `asn1:"optional,ext"`
	GNBCUUPCellGroupRelatedConfiguration []GNBCUUPCellGroupRelatedConfigurationItem `asn1:"optional,ext"`
	Cause                                *Cause                                     `asn1:"optional,ext"`
}
