package e1ap_ies

// GNBCUCPTNLAToAddItem From: 9_4_5_Information_Element_Definitions.txt:1125
// ASN.1 Data Type: SEQUENCE
type GNBCUCPTNLAToAddItem struct {
	TNLAssociationTransportLayerAddress CPTNLInformation    `aper:"mandatory"`
	TNLAssociationUsage                 TNLAssociationUsage `aper:"mandatory"`
}
