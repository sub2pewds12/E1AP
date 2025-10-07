package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBToSetupModItemEUTRAN is a generated SEQUENCE type.
type DRBToSetupModItemEUTRAN struct {
	DRBID                            aper.Integer                      `aper:"lb:1,ub:32,mandatory,ext"`
	PDCPConfiguration                PDCPConfiguration                 `aper:"mandatory,ext"`
	EUTRANQOS                        EUTRANQOS                         `aper:"mandatory,ext"`
	S1ULUPTNLInformation             UPTNLInformation                  `aper:"mandatory,ext"`
	DataForwardingInformationRequest *DataForwardingInformationRequest `aper:"optional,ext"`
	CellGroupInformation             []CellGroupInformationItem        `aper:"mandatory,ext"`
	DLUPParameters                   []UPParametersItem                `aper:"optional,ext"`
	DRBInactivityTimer               *aper.Integer                     `aper:"lb:1,ub:7200,optional,reject,ext"`
	IEExtensions                     *ProtocolExtensionContainer       `aper:"optional,ext"`
}
