package e1ap_ies

// GNBCUCPConfigurationUpdateAcknowledge From: 9_4_4_PDU_Definitions.txt:584
type GNBCUCPConfigurationUpdateAcknowledge struct {
	TransactionID                int64                          `asn1:"mandatory,reject,ext"`
	CriticalityDiagnostics       *CriticalityDiagnostics        `asn1:"optional,ignore,ext"`
	GNBCUCPTNLASetupList         []GNBCUCPTNLASetupItem         `asn1:"optional,ignore,ext"`
	GNBCUCPTNLAFailedToSetupList []GNBCUCPTNLAFailedToSetupItem `asn1:"optional,ignore,ext"`
	TransportLayerAddressInfo    *TransportLayerAddressInfo     `asn1:"optional,ignore,ext"`
}
