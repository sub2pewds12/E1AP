package e1ap_ies

// GNBCUCPTNLAFailedToSetupItem From: 9_4_5_Information_Element_Definitions.txt:1115
// ASN.1 Data Type: SEQUENCE
type GNBCUCPTNLAFailedToSetupItem struct {
	TNLAssociationTransportLayerAddress CPTNLInformation `aper:"mandatory"`
	Cause                               Cause            `aper:"mandatory,ignore"`
}
