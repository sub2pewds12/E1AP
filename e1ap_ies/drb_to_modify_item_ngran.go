package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBToModifyItemNGRAN From: 9_4_5_Information_Element_Definitions.txt:763
// ASN.1 Data Type: SEQUENCE
type DRBToModifyItemNGRAN struct {
	DRBID                        aper.Integer               `aper:"mandatory,ext"`
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
	DRBInactivityTimer           *aper.Integer              `aper:"optional,reject,ext"`
}
