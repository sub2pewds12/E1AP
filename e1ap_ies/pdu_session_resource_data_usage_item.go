package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDUSessionResourceDataUsageItem From: 9_4_5_Information_Element_Definitions.txt:1597
// ASN.1 Data Type: SEQUENCE
type PDUSessionResourceDataUsageItem struct {
	PDUSessionID         aper.Integer         `aper:"lb:0,ub:255,mandatory,ext"`
	MRDCUsageInformation MRDCUsageInformation `aper:"mandatory,ext"`
}
