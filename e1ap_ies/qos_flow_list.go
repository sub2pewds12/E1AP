package e1ap_ies

// QOSFlowList From: 9_4_5_Information_Element_Definitions.txt:1948
type QOSFlowList []QOSFlowItem

// QOSFlowItem From: 9_4_5_Information_Element_Definitions.txt:1950
type QOSFlowItem struct {
	QOSFlowIdentifier int64 `asn1:"lb:0,ub:63,mandatory,ext"`
}
