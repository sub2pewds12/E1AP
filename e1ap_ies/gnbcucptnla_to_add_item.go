package e1ap_ies

// GNBCUCPTNLAToAddItem is a generated SEQUENCE type.
type GNBCUCPTNLAToAddItem struct {
	TNLAssociationTransportLayerAddress CPTNLInformation    `aper:"mandatory"`
	TNLAssociationUsage                 TNLAssociationUsage `aper:"mandatory"`
}
