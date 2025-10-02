package e1ap_ies

// DataUsagePerQOSFlowList From: 9_4_5_Information_Element_Definitions.txt:357
type DataUsagePerQOSFlowList []DataUsagePerQOSFlowItem

// DataUsagePerQOSFlowItem From: 9_4_5_Information_Element_Definitions.txt:359
type DataUsagePerQOSFlowItem struct {
	QOSFlowIdentifier      int64                                   `asn1:"lb:0,ub:63,mandatory,ext"`
	SecondaryRATType       DataUsagePerQOSFlowItemSecondaryRATType `asn1:"mandatory,ext"`
	QOSFlowTimedReportList SEQUENCE                                `asn1:"mandatory,ext"`
}
