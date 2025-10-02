package e1ap_ies

// DRBUsageReportList From: 9_4_5_Information_Element_Definitions.txt:918
type DRBUsageReportList []DRBUsageReportItem

// DRBUsageReportItem From: 9_4_5_Information_Element_Definitions.txt:920
type DRBUsageReportItem struct {
	StartTimeStamp []byte `asn1:"mandatory,ext"`
	EndTimeStamp   []byte `asn1:"mandatory,ext"`
	UsageCountUL   int64  `asn1:"mandatory,ext"`
	UsageCountDL   int64  `asn1:"mandatory,ext"`
}
