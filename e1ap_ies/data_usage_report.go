package e1ap_ies

// DataUsageReport From: 9_4_4_PDU_Definitions.txt:1153
type DataUsageReport struct {
	GNBCUCPUEE1APID     int64                 `asn1:"lb:0,ub:4294967295,mandatory,reject,ext"`
	GNBCUUPUEE1APID     int64                 `asn1:"lb:0,ub:4294967295,mandatory,reject,ext"`
	DataUsageReportList []DataUsageReportItem `asn1:"mandatory,ignore,ext"`
}
