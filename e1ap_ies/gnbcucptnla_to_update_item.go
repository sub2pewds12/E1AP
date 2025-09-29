package e1ap_ies

// GNBCUCPTNLAToUpdateItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1145
type GNBCUCPTNLAToUpdateItem struct {
	TNLAssociationTransportLayerAddress CPTNLInformation     `asn1:"mandatory"`
	TNLAssociationUsage                 *TNLAssociationUsage `asn1:"optional"`
}
