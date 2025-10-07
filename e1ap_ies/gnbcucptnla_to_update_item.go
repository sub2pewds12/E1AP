package e1ap_ies

// GNBCUCPTNLAToUpdateItem is a generated SEQUENCE type.
type GNBCUCPTNLAToUpdateItem struct {
	TNLAssociationTransportLayerAddress CPTNLInformation            `aper:"mandatory"`
	TNLAssociationUsage                 *TNLAssociationUsage        `aper:"optional"`
	IEExtensions                        *ProtocolExtensionContainer `aper:"optional"`
}
