package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBToRemoveItemNGRAN From: 9_4_5_Information_Element_Definitions.txt:815
// ASN.1 Data Type: SEQUENCE
type DRBToRemoveItemNGRAN struct {
	DRBID aper.Integer `aper:"mandatory,ext"`
}
