package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBToSetupModItemNGRAN is a generated SEQUENCE type.
type DRBToSetupModItemNGRAN struct {
	DRBID                               aper.Integer                      `aper:"lb:1,ub:32,mandatory,ext"`
	SDAPConfiguration                   SDAPConfiguration                 `aper:"mandatory,ext"`
	PDCPConfiguration                   PDCPConfiguration                 `aper:"mandatory,ext"`
	CellGroupInformation                []CellGroupInformationItem        `aper:"mandatory,ext"`
	FlowMappingInformation              []QOSFlowQOSParameterItem         `aper:"mandatory,ext"`
	DRBDataForwardingInformationRequest *DataForwardingInformationRequest `aper:"optional,ext"`
	DRBInactivityTimer                  *aper.Integer                     `aper:"lb:1,ub:7200,optional,reject,ext"`
	PDCPSNStatusInformation             *PDCPSNStatusInformation          `aper:"optional,ext"`
}
