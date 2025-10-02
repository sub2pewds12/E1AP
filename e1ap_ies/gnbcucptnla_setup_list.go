package e1ap_ies

// GNBCUCPTNLASetupList From: 9_4_4_PDU_Definitions.txt:598
type GNBCUCPTNLASetupList []GNBCUCPTNLASetupItem

// GNBCUCPTNLASetupItem From: 9_4_5_Information_Element_Definitions.txt:1105
type GNBCUCPTNLASetupItem struct {
	TNLAssociationTransportLayerAddress CPTNLInformation `asn1:"mandatory,ext"`
}
