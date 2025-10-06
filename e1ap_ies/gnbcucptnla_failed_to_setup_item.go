package e1ap_ies

// GNBCUCPTNLAFailedToSetupItem is a generated SEQUENCE type.
type GNBCUCPTNLAFailedToSetupItem struct {
	TNLAssociationTransportLayerAddress CPTNLInformation `aper:"mandatory"`
	Cause                               Cause            `aper:"mandatory,ignore"`
}
