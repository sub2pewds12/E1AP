package e1ap_ies

// PDUSessionToNotifyList From: 9_4_5_Information_Element_Definitions.txt:1886
type PDUSessionToNotifyList []PDUSessionToNotifyItem

// PDUSessionToNotifyItem From: 9_4_5_Information_Element_Definitions.txt:1888
type PDUSessionToNotifyItem struct {
	PDUSessionID int64         `asn1:"lb:0,ub:255,mandatory,ext"`
	QOSFlowList  []QOSFlowItem `asn1:"mandatory,ext"`
}
