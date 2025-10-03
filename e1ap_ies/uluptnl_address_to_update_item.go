package e1ap_ies

import "github.com/lvdund/ngap/aper"

// ULUPTNLAddressToUpdateItem From: 9_4_5_Information_Element_Definitions.txt:2382
// ASN.1 Data Type: SEQUENCE
type ULUPTNLAddressToUpdateItem struct {
	OldTNLAdress aper.BitString `aper:"mandatory,ignore,ext"`
	NewTNLAdress aper.BitString `aper:"mandatory,ignore,ext"`
}
