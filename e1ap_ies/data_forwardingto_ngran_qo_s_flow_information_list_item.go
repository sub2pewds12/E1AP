package e1ap_ies

// DataForwardingtoNGRANQoSFlowInformationListItem is a generated SEQUENCE type.
type DataForwardingtoNGRANQoSFlowInformationListItem struct {
	QOSFlowIdentifier QOSFlowIdentifier           `aper:"lb:0,ub:63,mandatory,ext"`
	IEExtensions      *ProtocolExtensionContainer `aper:"optional,ext"`
}
