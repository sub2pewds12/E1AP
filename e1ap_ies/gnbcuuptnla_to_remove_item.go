package e1ap_ies

// GNBCUUPTNLAToRemoveItem From: 9_4_5_Information_Element_Definitions.txt:1155
// ASN.1 Data Type: SEQUENCE
type GNBCUUPTNLAToRemoveItem struct {
	TNLAssociationTransportLayerAddress        CPTNLInformation  `aper:"mandatory"`
	TNLAssociationTransportLayerAddressgNBCUCP *CPTNLInformation `aper:"optional"`
}
