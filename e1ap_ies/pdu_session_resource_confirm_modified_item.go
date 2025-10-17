package e1ap_ies

// PDUSessionResourceConfirmModifiedItem is a generated SEQUENCE type.
type PDUSessionResourceConfirmModifiedItem struct {
	PDUSessionID                PDUSessionID                  `aper:"lb:0,ub:255,mandatory,ext"`
	DRBConfirmModifiedListNGRAN []DRBConfirmModifiedItemNGRAN `aper:"optional,ext"`
	IEExtensions                *ProtocolExtensionContainer   `aper:"optional,ext"`
}
