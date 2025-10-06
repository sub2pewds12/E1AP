package e1ap_ies

// GNBCUUPTNLAToRemoveItem is a generated SEQUENCE type.
type GNBCUUPTNLAToRemoveItem struct {
	TNLAssociationTransportLayerAddress        CPTNLInformation  `aper:"mandatory"`
	TNLAssociationTransportLayerAddressgNBCUCP *CPTNLInformation `aper:"optional"`
}
