package e1ap_ies

// GNBCUCPTNLASetupItem From: 9_4_5_Information_Element_Definitions.txt:1105
// ASN.1 Data Type: SEQUENCE
type GNBCUCPTNLASetupItem struct {
	TNLAssociationTransportLayerAddress CPTNLInformation `aper:"mandatory,ext"`
}
