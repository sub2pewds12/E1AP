package e1ap_ies

import "github.com/lvdund/ngap/aper"

// GNBCUUPE1SetupResponse From: 9_4_4_PDU_Definitions.txt:384
// ASN.1 Data Type: SEQUENCE
type GNBCUUPE1SetupResponse struct {
	TransactionID             aper.Integer               `aper:"mandatory,reject,ext"`
	GNBCUCPName               *aper.OctetString          `aper:"optional,ignore,ext"`
	TransportLayerAddressInfo *TransportLayerAddressInfo `aper:"optional,ignore,ext"`
}
