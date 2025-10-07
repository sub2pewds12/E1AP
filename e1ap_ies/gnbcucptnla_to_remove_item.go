package e1ap_ies

// GNBCUCPTNLAToRemoveItem is a generated SEQUENCE type.
type GNBCUCPTNLAToRemoveItem struct {
	TNLAssociationTransportLayerAddress CPTNLInformation            `aper:"mandatory"`
	IEExtensions                        *ProtocolExtensionContainer `aper:"optional"`
}
