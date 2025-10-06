package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBRequiredToModifyItemEUTRAN is a generated SEQUENCE type.
type DRBRequiredToModifyItemEUTRAN struct {
	DRBID                                aper.Integer                               `aper:"lb:1,ub:32,mandatory,ext"`
	S1DLUPTNLInformation                 *UPTNLInformation                          `aper:"optional,ext"`
	GNBCUUPCellGroupRelatedConfiguration []GNBCUUPCellGroupRelatedConfigurationItem `aper:"optional,ext"`
	Cause                                *Cause                                     `aper:"optional,ignore,ext"`
}
