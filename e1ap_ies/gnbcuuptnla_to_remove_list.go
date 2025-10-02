package e1ap_ies

// GNBCUUPTNLAToRemoveList From: 9_4_4_PDU_Definitions.txt:508
type GNBCUUPTNLAToRemoveList []GNBCUUPTNLAToRemoveItem

// GNBCUUPTNLAToRemoveItem From: 9_4_5_Information_Element_Definitions.txt:1155
type GNBCUUPTNLAToRemoveItem struct {
	TNLAssociationTransportLayerAddress        CPTNLInformation  `asn1:"mandatory"`
	TNLAssociationTransportLayerAddressgNBCUCP *CPTNLInformation `asn1:"optional"`
}
