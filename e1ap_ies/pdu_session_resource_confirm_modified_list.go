package e1ap_ies

// PDUSessionResourceConfirmModifiedList From: 9_4_5_Information_Element_Definitions.txt:1667
type PDUSessionResourceConfirmModifiedList []PDUSessionResourceConfirmModifiedItem

// PDUSessionResourceConfirmModifiedItem From: 9_4_5_Information_Element_Definitions.txt:1669
type PDUSessionResourceConfirmModifiedItem struct {
	PDUSessionID                int64                         `asn1:"lb:0,ub:255,mandatory,ext"`
	DRBConfirmModifiedListNGRAN []DRBConfirmModifiedItemNGRAN `asn1:"optional,ext"`
}
