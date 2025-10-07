package e1ap_ies

import "github.com/lvdund/ngap/aper"

// PDUSessionResourceRequiredToModifyItem is a generated SEQUENCE type.
type PDUSessionResourceRequiredToModifyItem struct {
	PDUSessionID                 aper.Integer                   `aper:"lb:0,ub:255,mandatory,ext"`
	NGDLUPTNLInformation         *UPTNLInformation              `aper:"optional,ext"`
	DRBRequiredToModifyListNGRAN []DRBRequiredToModifyItemNGRAN `aper:"optional,ext"`
	DRBRequiredToRemoveListNGRAN []DRBRequiredToRemoveItemNGRAN `aper:"optional,ext"`
	IEExtensions                 *ProtocolExtensionContainer    `aper:"optional,ext"`
}
