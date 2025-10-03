package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBConfirmModifiedItemEUTRAN From: 9_4_5_Information_Element_Definitions.txt:440
// ASN.1 Data Type: SEQUENCE
type DRBConfirmModifiedItemEUTRAN struct {
	DRBID                aper.Integer               `aper:"mandatory,ext"`
	CellGroupInformation []CellGroupInformationItem `aper:"optional,ext"`
}
