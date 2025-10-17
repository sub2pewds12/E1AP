package e1ap_ies

// QOSFlowMappingItem is a generated SEQUENCE type.
type QOSFlowMappingItem struct {
	QOSFlowIdentifier        QOSFlowIdentifier           `aper:"lb:0,ub:63,mandatory,ext"`
	QoSFlowMappingIndication *QOSFlowMappingIndication   `aper:"optional,ext"`
	IEExtensions             *ProtocolExtensionContainer `aper:"optional,ext"`
}
