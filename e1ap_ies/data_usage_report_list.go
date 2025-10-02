package e1ap_ies

// DataUsageReportList From: 9_4_5_Information_Element_Definitions.txt:371
type DataUsageReportList []DataUsageReportItem

// DataUsageReportItem From: 9_4_5_Information_Element_Definitions.txt:373
type DataUsageReportItem struct {
	DRBID              int64                `asn1:"mandatory,ext"`
	RATType            RATType              `asn1:"mandatory,ext"`
	DRBUsageReportList []DRBUsageReportItem `asn1:"lb:1,ub:Maxnooftimeperiods,mandatory,ext"`
}
