package e1ap_ies

// GNBCUCPTNLASetupItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1105
type GNBCUCPTNLASetupItem struct {
	TNLAssociationTransportLayerAddress CPTNLInformation `asn1:"mandatory,ext"`
}
