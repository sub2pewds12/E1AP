package e1ap_ies

// QOSFlowItem is a generated SEQUENCE type.
type QOSFlowItem struct {
	QOSFlowIdentifier QOSFlowIdentifier           `aper:"lb:0,ub:63,mandatory,ext"`
	IEExtensions      *ProtocolExtensionContainer `aper:"optional,ext"`
}
