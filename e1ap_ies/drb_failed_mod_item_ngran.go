package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBFailedModItemNGRAN From: 9_4_5_Information_Element_Definitions.txt:505
// ASN.1 Data Type: SEQUENCE
type DRBFailedModItemNGRAN struct {
	DRBID aper.Integer `aper:"mandatory,ext"`
	Cause Cause        `aper:"mandatory,ignore,ext"`
}
