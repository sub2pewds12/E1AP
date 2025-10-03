package e1ap_ies

import "github.com/lvdund/ngap/aper"

// TNLAvailableCapacityIndicator From: 9_4_5_Information_Element_Definitions.txt:2233
// ASN.1 Data Type: SEQUENCE
type TNLAvailableCapacityIndicator struct {
	DLTNLOfferedCapacity   aper.Integer `aper:"mandatory,ext"`
	DLTNLAvailableCapacity aper.Integer `aper:"mandatory,ext"`
	ULTNLOfferedCapacity   aper.Integer `aper:"mandatory,ext"`
	ULTNLAvailableCapacity aper.Integer `aper:"mandatory,ext"`
}
