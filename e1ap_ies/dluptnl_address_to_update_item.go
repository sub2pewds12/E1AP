package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DLUPTNLAddressToUpdateItem From: 9_4_5_Information_Element_Definitions.txt:402
// ASN.1 Data Type: SEQUENCE
type DLUPTNLAddressToUpdateItem struct {
	OldTNLAdress aper.BitString `aper:"mandatory,ignore,ext"`
	NewTNLAdress aper.BitString `aper:"mandatory,ignore,ext"`
}
