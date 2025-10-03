package e1ap_ies

import "github.com/lvdund/ngap/aper"

// MRDCDataUsageReportItem From: 9_4_5_Information_Element_Definitions.txt:1315
// ASN.1 Data Type: SEQUENCE
type MRDCDataUsageReportItem struct {
	StartTimeStamp aper.OctetString `aper:"lb:4,ub:4,mandatory,ext"`
	EndTimeStamp   aper.OctetString `aper:"lb:4,ub:4,mandatory,ext"`
	UsageCountUL   aper.Integer     `aper:"lb:0,ub:18446744073709551615,mandatory,ext"`
	UsageCountDL   aper.Integer     `aper:"lb:0,ub:18446744073709551615,mandatory,ext"`
}
