package e1ap_ies

// GNBCUCPTNLAFailedToSetupList From: 9_4_4_PDU_Definitions.txt:599
type GNBCUCPTNLAFailedToSetupList []GNBCUCPTNLAFailedToSetupItem

// GNBCUCPTNLAFailedToSetupItem From: 9_4_5_Information_Element_Definitions.txt:1115
type GNBCUCPTNLAFailedToSetupItem struct {
	TNLAssociationTransportLayerAddress CPTNLInformation `asn1:"mandatory"`
	Cause                               Cause            `asn1:"mandatory"`
}
