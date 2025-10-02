package e1ap_ies

// PDUSessionResourceFailedList From: 9_4_5_Information_Element_Definitions.txt:1680
type PDUSessionResourceFailedList []PDUSessionResourceFailedItem

// PDUSessionResourceFailedItem From: 9_4_5_Information_Element_Definitions.txt:1682
type PDUSessionResourceFailedItem struct {
	PDUSessionID int64 `asn1:"lb:0,ub:255,mandatory,ext"`
	Cause        Cause `asn1:"mandatory,ext"`
}
