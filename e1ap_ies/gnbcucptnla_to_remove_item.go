package e1ap_ies

// GNBCUCPTNLAToRemoveItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1135
type GNBCUCPTNLAToRemoveItem struct {
	TNLAssociationTransportLayerAddress CPTNLInformation `asn1:"mandatory"`
}
