package e1ap_ies

// GNBCUCPTNLAToAddItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1125
type GNBCUCPTNLAToAddItem struct {
	TNLAssociationTransportLayerAddress CPTNLInformation    `asn1:"mandatory"`
	TNLAssociationUsage                 TNLAssociationUsage `asn1:"mandatory"`
}
