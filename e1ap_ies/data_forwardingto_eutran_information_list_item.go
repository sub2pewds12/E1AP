package e1ap_ies

// DataForwardingtoEUTRANInformationListItem is a generated SEQUENCE type.
type DataForwardingtoEUTRANInformationListItem struct {
	DataForwardingTunnelInformation UPTNLInformation            `aper:"mandatory,ext"`
	QOSFlowsToBeForwardedList       []QOSFlowsToBeForwardedItem `aper:"mandatory,ext"`
	IEExtensions                    *ProtocolExtensionContainer `aper:"optional,ext"`
}
