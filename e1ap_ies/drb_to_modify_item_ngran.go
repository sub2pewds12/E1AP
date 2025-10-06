package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBToModifyItemNGRAN is a generated SEQUENCE type.
type DRBToModifyItemNGRAN struct {
	DRBID                        aper.Integer               `aper:"lb:1,ub:32,mandatory,ext"`
	SDAPConfiguration            *SDAPConfiguration         `aper:"optional,ext"`
	PDCPConfiguration            *PDCPConfiguration         `aper:"optional,ext"`
	DRBDataForwardingInformation *DataForwardingInformation `aper:"optional,ext"`
	PDCPSNStatusRequest          *PDCPSNStatusRequest       `aper:"optional,ext"`
	PdcpSNStatusInformation      *PDCPSNStatusInformation   `aper:"optional,ext"`
	DLUPParameters               []UPParametersItem         `aper:"optional,ext"`
	CellGroupToAdd               []CellGroupInformationItem `aper:"optional,ext"`
	CellGroupToModify            []CellGroupInformationItem `aper:"optional,ext"`
	CellGroupToRemove            []CellGroupInformationItem `aper:"optional,ext"`
	FlowMappingInformation       []QOSFlowQOSParameterItem  `aper:"optional,ext"`
	DRBInactivityTimer           *aper.Integer              `aper:"lb:1,ub:7200,optional,reject,ext"`
}
