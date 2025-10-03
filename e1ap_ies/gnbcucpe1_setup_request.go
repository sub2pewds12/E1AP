package e1ap_ies

import "github.com/lvdund/ngap/aper"

// GNBCUCPE1SetupRequest From: 9_4_4_PDU_Definitions.txt:427
// ASN.1 Data Type: SEQUENCE
type GNBCUCPE1SetupRequest struct {
	TransactionID             aper.Integer               `aper:"mandatory,reject,ext"`
	GNBCUCPName               *aper.OctetString          `aper:"optional,ignore,ext"`
	TransportLayerAddressInfo *TransportLayerAddressInfo `aper:"optional,ignore,ext"`
}
