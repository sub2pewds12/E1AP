package e1ap_ies

// GNBCUCPTNLAToUpdateList From: 9_4_4_PDU_Definitions.txt:576
type GNBCUCPTNLAToUpdateList []GNBCUCPTNLAToUpdateItem

// GNBCUCPTNLAToUpdateItem From: 9_4_5_Information_Element_Definitions.txt:1145
type GNBCUCPTNLAToUpdateItem struct {
	TNLAssociationTransportLayerAddress CPTNLInformation     `asn1:"mandatory"`
	TNLAssociationUsage                 *TNLAssociationUsage `asn1:"optional"`
}
