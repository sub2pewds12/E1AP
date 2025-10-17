package e1ap_ies

// QOSFlowFailedItem is a generated SEQUENCE type.
type QOSFlowFailedItem struct {
	QOSFlowIdentifier QOSFlowIdentifier           `aper:"lb:0,ub:63,mandatory,ext"`
	Cause             Cause                       `aper:"mandatory,ignore,ext"`
	IEExtensions      *ProtocolExtensionContainer `aper:"optional,ext"`
}
