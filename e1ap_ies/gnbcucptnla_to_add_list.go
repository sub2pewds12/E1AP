package e1ap_ies

// GNBCUCPTNLAToAddList From: 9_4_4_PDU_Definitions.txt:574
type GNBCUCPTNLAToAddList []GNBCUCPTNLAToAddItem

// GNBCUCPTNLAToAddItem From: 9_4_5_Information_Element_Definitions.txt:1125
type GNBCUCPTNLAToAddItem struct {
	TNLAssociationTransportLayerAddress CPTNLInformation    `asn1:"mandatory"`
	TNLAssociationUsage                 TNLAssociationUsage `asn1:"mandatory"`
}
