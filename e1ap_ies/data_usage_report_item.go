package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DataUsageReportItem From: 9_4_5_Information_Element_Definitions.txt:373
// ASN.1 Data Type: SEQUENCE
type DataUsageReportItem struct {
	DRBID              aper.Integer         `aper:"mandatory,ext"`
	RATType            RATType              `aper:"mandatory,ext"`
	DRBUsageReportList []DRBUsageReportItem `aper:"lb:1,ub:Maxnooftimeperiods,mandatory,ext"`
}
