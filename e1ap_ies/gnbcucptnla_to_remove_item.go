package e1ap_ies

// GNBCUCPTNLAToRemoveItem From: 9_4_5_Information_Element_Definitions.txt:1135
// ASN.1 Data Type: SEQUENCE
type GNBCUCPTNLAToRemoveItem struct {
	TNLAssociationTransportLayerAddress CPTNLInformation `aper:"mandatory"`
}
