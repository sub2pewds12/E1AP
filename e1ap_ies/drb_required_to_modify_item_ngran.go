package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBRequiredToModifyItemNGRAN From: 9_4_5_Information_Element_Definitions.txt:607
// ASN.1 Data Type: SEQUENCE
type DRBRequiredToModifyItemNGRAN struct {
	DRBID                                aper.Integer                               `aper:"mandatory,ext"`
	GNBCUUPCellGroupRelatedConfiguration []GNBCUUPCellGroupRelatedConfigurationItem `aper:"optional,ext"`
	FlowToRemove                         []QOSFlowItem                              `aper:"optional,ext"`
	Cause                                *Cause                                     `aper:"optional,ignore,ext"`
}
