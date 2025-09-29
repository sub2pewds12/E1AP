package e1ap_ies

// GNBCUUPTNLAToRemoveItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1155
type GNBCUUPTNLAToRemoveItem struct {
	TNLAssociationTransportLayerAddress        CPTNLInformation  `asn1:"mandatory"`
	TNLAssociationTransportLayerAddressgNBCUCP *CPTNLInformation `asn1:"optional"`
}
