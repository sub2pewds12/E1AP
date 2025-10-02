package e1ap_ies

// GNBCUCPTNLAToRemoveList From: 9_4_4_PDU_Definitions.txt:575
type GNBCUCPTNLAToRemoveList []GNBCUCPTNLAToRemoveItem

// GNBCUCPTNLAToRemoveItem From: 9_4_5_Information_Element_Definitions.txt:1135
type GNBCUCPTNLAToRemoveItem struct {
	TNLAssociationTransportLayerAddress CPTNLInformation `asn1:"mandatory"`
}
