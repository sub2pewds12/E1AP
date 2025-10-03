package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DRBActivityItem From: 9_4_5_Information_Element_Definitions.txt:427
// ASN.1 Data Type: SEQUENCE
type DRBActivityItem struct {
	DRBID       aper.Integer `aper:"mandatory,ext"`
	DRBActivity DRBActivity  `aper:"mandatory,ext"`
}
