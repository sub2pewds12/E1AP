package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBRequiredToModifyItemEUTRAN From: 9_4_5_Information_Element_Definitions.txt:592
// ASN.1 Data Type: SEQUENCE
type DRBRequiredToModifyItemEUTRAN struct {
	DRBID                                aper.Integer                               `aper:"mandatory,ext"`
	S1DLUPTNLInformation                 *UPTNLInformation                          `aper:"optional,ext"`
	GNBCUUPCellGroupRelatedConfiguration []GNBCUUPCellGroupRelatedConfigurationItem `aper:"optional,ext"`
	Cause                                *Cause                                     `aper:"optional,ignore,ext"`
}
