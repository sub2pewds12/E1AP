package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBToModifyItemEUTRAN From: 9_4_5_Information_Element_Definitions.txt:740
// ASN.1 Data Type: SEQUENCE
type DRBToModifyItemEUTRAN struct {
	DRBID                     aper.Integer               `aper:"mandatory,ext"`
	PDCPConfiguration         *PDCPConfiguration         `aper:"optional,ext"`
	EUTRANQOS                 *EUTRANQOS                 `aper:"optional,ext"`
	S1ULUPTNLInformation      *UPTNLInformation          `aper:"optional,ext"`
	DataForwardingInformation *DataForwardingInformation `aper:"optional,ext"`
	PDCPSNStatusRequest       *PDCPSNStatusRequest       `aper:"optional,ext"`
	PDCPSNStatusInformation   *PDCPSNStatusInformation   `aper:"optional,ext"`
	DLUPParameters            []UPParametersItem         `aper:"optional,ext"`
	CellGroupToAdd            []CellGroupInformationItem `aper:"optional,ext"`
	CellGroupToModify         []CellGroupInformationItem `aper:"optional,ext"`
	CellGroupToRemove         []CellGroupInformationItem `aper:"optional,ext"`
	DRBInactivityTimer        *aper.Integer              `aper:"optional,reject,ext"`
}
