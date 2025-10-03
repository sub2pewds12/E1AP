package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBFailedToModifyItemEUTRAN From: 9_4_5_Information_Element_Definitions.txt:518
// ASN.1 Data Type: SEQUENCE
type DRBFailedToModifyItemEUTRAN struct {
	DRBID aper.Integer `aper:"mandatory,ext"`
	Cause Cause        `aper:"mandatory,ignore,ext"`
}
