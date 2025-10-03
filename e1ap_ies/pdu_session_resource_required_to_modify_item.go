package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDUSessionResourceRequiredToModifyItem From: 9_4_5_Information_Element_Definitions.txt:1741
// ASN.1 Data Type: SEQUENCE
type PDUSessionResourceRequiredToModifyItem struct {
	PDUSessionID                 aper.Integer                   `aper:"lb:0,ub:255,mandatory,ext"`
	NGDLUPTNLInformation         *UPTNLInformation              `aper:"optional,ext"`
	DRBRequiredToModifyListNGRAN []DRBRequiredToModifyItemNGRAN `aper:"optional,ext"`
	DRBRequiredToRemoveListNGRAN []DRBRequiredToRemoveItemNGRAN `aper:"optional,ext"`
}
