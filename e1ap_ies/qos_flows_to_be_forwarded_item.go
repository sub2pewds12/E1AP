package e1ap_ies

// QOSFlowsToBeForwardedItem is a generated SEQUENCE type.
type QOSFlowsToBeForwardedItem struct {
	QOSFlowIdentifier QOSFlowIdentifier           `aper:"lb:0,ub:63,mandatory,ext"`
	IEExtensions      *ProtocolExtensionContainer `aper:"optional,ext"`
}
