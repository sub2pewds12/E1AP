package e1ap_ies

// PDUSessionResourceFailedModList From: 9_4_5_Information_Element_Definitions.txt:1693
type PDUSessionResourceFailedModList []PDUSessionResourceFailedModItem

// PDUSessionResourceFailedModItem From: 9_4_5_Information_Element_Definitions.txt:1695
type PDUSessionResourceFailedModItem struct {
	PDUSessionID int64 `asn1:"lb:0,ub:255,mandatory,ext"`
	Cause        Cause `asn1:"mandatory,ext"`
}
