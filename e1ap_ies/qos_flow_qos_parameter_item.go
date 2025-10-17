package e1ap_ies

// QOSFlowQOSParameterItem is a generated SEQUENCE type.
type QOSFlowQOSParameterItem struct {
	QOSFlowIdentifier         QOSFlowIdentifier           `aper:"lb:0,ub:63,mandatory,ext"`
	QoSFlowLevelQoSParameters QoSFlowLevelQoSParameters   `aper:"mandatory,ext"`
	QoSFlowMappingIndication  *QOSFlowMappingIndication   `aper:"optional,ext"`
	IEExtensions              *ProtocolExtensionContainer `aper:"optional,ext"`
}
