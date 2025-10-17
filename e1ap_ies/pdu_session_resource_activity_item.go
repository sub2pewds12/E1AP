package e1ap_ies

// PDUSessionResourceActivityItem is a generated SEQUENCE type.
type PDUSessionResourceActivityItem struct {
	PDUSessionID               PDUSessionID                `aper:"lb:0,ub:255,mandatory,ext"`
	PDUSessionResourceActivity PDUSessionResourceActivity  `aper:"mandatory,ext"`
	IEExtensions               *ProtocolExtensionContainer `aper:"optional,ext"`
}
