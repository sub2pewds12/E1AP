package e1ap_ies

import "github.com/lvdund/ngap/aper"

// GNBCUCPConfigurationUpdateAcknowledge From: 9_4_4_PDU_Definitions.txt:584
// ASN.1 Data Type: SEQUENCE
type GNBCUCPConfigurationUpdateAcknowledge struct {
	TransactionID                aper.Integer                   `aper:"mandatory,reject,ext"`
	CriticalityDiagnostics       *CriticalityDiagnostics        `aper:"optional,ignore,ext"`
	GNBCUCPTNLASetupList         []GNBCUCPTNLASetupItem         `aper:"optional,ignore,ext"`
	GNBCUCPTNLAFailedToSetupList []GNBCUCPTNLAFailedToSetupItem `aper:"optional,ignore,ext"`
	TransportLayerAddressInfo    *TransportLayerAddressInfo     `aper:"optional,ignore,ext"`
}
