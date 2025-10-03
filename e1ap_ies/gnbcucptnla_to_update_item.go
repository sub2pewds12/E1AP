package e1ap_ies

// GNBCUCPTNLAToUpdateItem From: 9_4_5_Information_Element_Definitions.txt:1145
// ASN.1 Data Type: SEQUENCE
type GNBCUCPTNLAToUpdateItem struct {
	TNLAssociationTransportLayerAddress CPTNLInformation     `aper:"mandatory"`
	TNLAssociationUsage                 *TNLAssociationUsage `aper:"optional"`
}
