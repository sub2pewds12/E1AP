package e1ap_ies

// PDUSessionResourceToRemoveList From: 9_4_5_Information_Element_Definitions.txt:1821
type PDUSessionResourceToRemoveList []PDUSessionResourceToRemoveItem

// PDUSessionResourceToRemoveItem From: 9_4_5_Information_Element_Definitions.txt:1823
type PDUSessionResourceToRemoveItem struct {
	PDUSessionID int64 `asn1:"lb:0,ub:255,mandatory,ext"`
}
