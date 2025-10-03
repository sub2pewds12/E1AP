package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBRequiredToRemoveItemNGRAN From: 9_4_5_Information_Element_Definitions.txt:827
// ASN.1 Data Type: SEQUENCE
type DRBRequiredToRemoveItemNGRAN struct {
	DRBID aper.Integer `aper:"mandatory,ext"`
	Cause Cause        `aper:"mandatory,ignore,ext"`
}
