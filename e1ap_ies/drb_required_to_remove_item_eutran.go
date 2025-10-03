package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBRequiredToRemoveItemEUTRAN From: 9_4_5_Information_Element_Definitions.txt:802
// ASN.1 Data Type: SEQUENCE
type DRBRequiredToRemoveItemEUTRAN struct {
	DRBID aper.Integer `aper:"mandatory,ext"`
	Cause Cause        `aper:"mandatory,ignore,ext"`
}
