package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DataUsageReport From: 9_4_4_PDU_Definitions.txt:1153
// ASN.1 Data Type: SEQUENCE
type DataUsageReport struct {
	GNBCUCPUEE1APID     aper.Integer          `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID     aper.Integer          `aper:"lb:0,ub:4294967295,mandatory,reject,ext"`
	DataUsageReportList []DataUsageReportItem `aper:"mandatory,ignore,ext"`
}
