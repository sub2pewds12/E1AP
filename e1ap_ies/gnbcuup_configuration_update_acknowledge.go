package e1ap_ies

// GNBCUUPConfigurationUpdateAcknowledge From: 9_4_4_PDU_Definitions.txt:516
type GNBCUUPConfigurationUpdateAcknowledge struct {
	TransactionID             int64                      `asn1:"mandatory,reject,ext"`
	CriticalityDiagnostics    *CriticalityDiagnostics    `asn1:"optional,ignore,ext"`
	TransportLayerAddressInfo *TransportLayerAddressInfo `asn1:"optional,ignore,ext"`
}
