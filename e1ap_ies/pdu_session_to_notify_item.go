package e1ap_ies

// PDUSessionToNotifyItem is a generated SEQUENCE type.
type PDUSessionToNotifyItem struct {
	PDUSessionID PDUSessionID                `aper:"lb:0,ub:255,mandatory,ext"`
	QOSFlowList  []QOSFlowItem               `aper:"mandatory,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}
