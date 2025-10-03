package e1ap_ies

import "github.com/lvdund/ngap/aper"

// GNBCUUPConfigurationUpdateAcknowledge From: 9_4_4_PDU_Definitions.txt:516
// ASN.1 Data Type: SEQUENCE
type GNBCUUPConfigurationUpdateAcknowledge struct {
	TransactionID             aper.Integer               `aper:"mandatory,reject,ext"`
	CriticalityDiagnostics    *CriticalityDiagnostics    `aper:"optional,ignore,ext"`
	TransportLayerAddressInfo *TransportLayerAddressInfo `aper:"optional,ignore,ext"`
}
