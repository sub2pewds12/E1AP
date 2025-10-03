package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBToSetupModItemNGRAN From: 9_4_5_Information_Element_Definitions.txt:900
// ASN.1 Data Type: SEQUENCE
type DRBToSetupModItemNGRAN struct {
	DRBID                               aper.Integer                      `aper:"mandatory,ext"`
	SDAPConfiguration                   SDAPConfiguration                 `aper:"mandatory,ext"`
	PDCPConfiguration                   PDCPConfiguration                 `aper:"mandatory,ext"`
	CellGroupInformation                []CellGroupInformationItem        `aper:"mandatory,ext"`
	FlowMappingInformation              []QOSFlowQOSParameterItem         `aper:"mandatory,ext"`
	DRBDataForwardingInformationRequest *DataForwardingInformationRequest `aper:"optional,ext"`
	DRBInactivityTimer                  *aper.Integer                     `aper:"optional,reject,ext"`
	PDCPSNStatusInformation             *PDCPSNStatusInformation          `aper:"optional,ext"`
}
