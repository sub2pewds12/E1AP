package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBRequiredToModifyItemNGRAN is a generated SEQUENCE type.
type DRBRequiredToModifyItemNGRAN struct {
	DRBID                                aper.Integer                               `aper:"lb:1,ub:32,mandatory,ext"`
	GNBCUUPCellGroupRelatedConfiguration []GNBCUUPCellGroupRelatedConfigurationItem `aper:"optional,ext"`
	FlowToRemove                         []QOSFlowItem                              `aper:"optional,ext"`
	Cause                                *Cause                                     `aper:"optional,ignore,ext"`
}
