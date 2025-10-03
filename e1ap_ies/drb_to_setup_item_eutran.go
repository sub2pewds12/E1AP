package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBToSetupItemEUTRAN From: 9_4_5_Information_Element_Definitions.txt:840
// ASN.1 Data Type: SEQUENCE
type DRBToSetupItemEUTRAN struct {
	DRBID                            aper.Integer                      `aper:"mandatory,ext"`
	PDCPConfiguration                PDCPConfiguration                 `aper:"mandatory,ext"`
	EUTRANQOS                        EUTRANQOS                         `aper:"mandatory,ext"`
	S1ULUPTNLInformation             UPTNLInformation                  `aper:"mandatory,ext"`
	DataForwardingInformationRequest *DataForwardingInformationRequest `aper:"optional,ext"`
	CellGroupInformation             []CellGroupInformationItem        `aper:"mandatory,ext"`
	DLUPParameters                   []UPParametersItem                `aper:"optional,ext"`
	DRBInactivityTimer               *aper.Integer                     `aper:"optional,reject,ext"`
	ExistingAllocatedS1DLUPTNLInfo   *UPTNLInformation                 `aper:"optional,ext"`
}
