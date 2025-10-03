package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBToRemoveItemEUTRAN From: 9_4_5_Information_Element_Definitions.txt:790
// ASN.1 Data Type: SEQUENCE
type DRBToRemoveItemEUTRAN struct {
	DRBID aper.Integer `aper:"mandatory,ext"`
}
