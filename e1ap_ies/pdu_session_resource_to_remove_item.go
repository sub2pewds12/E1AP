package e1ap_ies

// PDUSessionResourceToRemoveItem is a generated SEQUENCE type.
type PDUSessionResourceToRemoveItem struct {
	PDUSessionID PDUSessionID                `aper:"lb:0,ub:255,mandatory,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}
