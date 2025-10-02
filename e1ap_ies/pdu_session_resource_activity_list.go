package e1ap_ies

// PDUSessionResourceActivityList From: 9_4_5_Information_Element_Definitions.txt:1653
type PDUSessionResourceActivityList []PDUSessionResourceActivityItem

// PDUSessionResourceActivityItem From: 9_4_5_Information_Element_Definitions.txt:1655
type PDUSessionResourceActivityItem struct {
	PDUSessionID               int64                      `asn1:"lb:0,ub:255,mandatory,ext"`
	PDUSessionResourceActivity PDUSessionResourceActivity `asn1:"mandatory,ext"`
}
