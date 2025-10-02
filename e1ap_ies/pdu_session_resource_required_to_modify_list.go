package e1ap_ies

// PDUSessionResourceRequiredToModifyList From: 9_4_5_Information_Element_Definitions.txt:1739
type PDUSessionResourceRequiredToModifyList []PDUSessionResourceRequiredToModifyItem

// PDUSessionResourceRequiredToModifyItem From: 9_4_5_Information_Element_Definitions.txt:1741
type PDUSessionResourceRequiredToModifyItem struct {
	PDUSessionID                 int64                          `asn1:"lb:0,ub:255,mandatory,ext"`
	NGDLUPTNLInformation         *UPTNLInformation              `asn1:"optional,ext"`
	DRBRequiredToModifyListNGRAN []DRBRequiredToModifyItemNGRAN `asn1:"optional,ext"`
	DRBRequiredToRemoveListNGRAN []DRBRequiredToRemoveItemNGRAN `asn1:"optional,ext"`
}
