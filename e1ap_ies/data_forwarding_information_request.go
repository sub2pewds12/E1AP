package e1ap_ies

// DataForwardingInformationRequest is a generated SEQUENCE type.
type DataForwardingInformationRequest struct {
	DataForwardingRequest         DataForwardingRequest       `aper:"mandatory,ext"`
	QOSFlowsForwardedOnFwdTunnels []QOSFlowMappingItem        `aper:"optional,ext"`
	IEExtensions                  *ProtocolExtensionContainer `aper:"optional,ext"`
}
