package e1ap_ies

// PDUSessionResourceRequiredToModifyItem is a generated SEQUENCE type.
type PDUSessionResourceRequiredToModifyItem struct {
	PDUSessionID                 PDUSessionID                   `aper:"lb:0,ub:255,mandatory,ext"`
	NGDLUPTNLInformation         *UPTNLInformation              `aper:"optional,ext"`
	DRBRequiredToModifyListNGRAN []DRBRequiredToModifyItemNGRAN `aper:"optional,ext"`
	DRBRequiredToRemoveListNGRAN []DRBRequiredToRemoveItemNGRAN `aper:"optional,ext"`
	IEExtensions                 *ProtocolExtensionContainer    `aper:"optional,ext"`
	// Possible extensions:
	// - UPTNLInformation (ID: id-redundant-nG-DL-UP-TNL-Information)
}
