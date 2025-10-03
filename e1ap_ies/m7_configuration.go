package e1ap_ies

import "github.com/lvdund/ngap/aper"

// M7Configuration From: 9_4_5_Information_Element_Definitions.txt:1365
// ASN.1 Data Type: SEQUENCE
type M7Configuration struct {
	M7period     aper.Integer `aper:"mandatory,ext"`
	M7LinksToLog LinksToLog   `aper:"mandatory,ext"`
}
