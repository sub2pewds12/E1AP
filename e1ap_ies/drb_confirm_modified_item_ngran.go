package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBConfirmModifiedItemNGRAN From: 9_4_5_Information_Element_Definitions.txt:453
// ASN.1 Data Type: SEQUENCE
type DRBConfirmModifiedItemNGRAN struct {
	DRBID                aper.Integer               `aper:"mandatory,ext"`
	CellGroupInformation []CellGroupInformationItem `aper:"optional,ext"`
}
