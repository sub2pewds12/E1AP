package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBToSetupItemNGRAN is a generated SEQUENCE type.
type DRBToSetupItemNGRAN struct {
	DRBID                               aper.Integer                      `aper:"lb:1,ub:32,mandatory,ext"`
	SDAPConfiguration                   SDAPConfiguration                 `aper:"mandatory,ext"`
	PDCPConfiguration                   PDCPConfiguration                 `aper:"mandatory,ext"`
	CellGroupInformation                []CellGroupInformationItem        `aper:"mandatory,ext"`
	QOSFlowInformationToBeSetup         []QOSFlowQOSParameterItem         `aper:"mandatory,ext"`
	DRBDataForwardingInformationRequest *DataForwardingInformationRequest `aper:"optional,ext"`
	DRBInactivityTimer                  *aper.Integer                     `aper:"lb:1,ub:7200,optional,reject,ext"`
	PDCPSNStatusInformation             *PDCPSNStatusInformation          `aper:"optional,ext"`
	IEExtensions                        *ProtocolExtensionContainer       `aper:"optional,ext"`
}
