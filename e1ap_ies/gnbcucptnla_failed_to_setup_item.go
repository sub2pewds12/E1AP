package e1ap_ies

// GNBCUCPTNLAFailedToSetupItem represents the ASN.1 definition from 9_4_5_Information_Element_Definitions.txt:1115
type GNBCUCPTNLAFailedToSetupItem struct {
	TNLAssociationTransportLayerAddress CPTNLInformation `asn1:"mandatory"`
	Cause                               Cause            `asn1:"mandatory"`
}
