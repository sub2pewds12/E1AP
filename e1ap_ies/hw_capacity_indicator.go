package e1ap_ies

import "github.com/lvdund/ngap/aper"

// HWCapacityIndicator From: 9_4_5_Information_Element_Definitions.txt:1228
// ASN.1 Data Type: SEQUENCE
type HWCapacityIndicator struct {
	OfferedThroughput   aper.Integer `aper:"mandatory,ext"`
	AvailableThroughput aper.Integer `aper:"mandatory,ext"`
}
