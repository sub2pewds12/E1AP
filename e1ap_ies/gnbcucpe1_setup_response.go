package e1ap_ies

import "github.com/lvdund/ngap/aper"

// GNBCUCPE1SetupResponse From: 9_4_4_PDU_Definitions.txt:445
// ASN.1 Data Type: SEQUENCE
type GNBCUCPE1SetupResponse struct {
	TransactionID             aper.Integer               `aper:"mandatory,reject,ext"`
	GNBCUUPID                 aper.Integer               `aper:"lb:0,ub:68719476735,mandatory,reject,ext"`
	GNBCUUPName               *aper.OctetString          `aper:"optional,ignore,ext"`
	CNSupport                 CNSupport                  `aper:"mandatory,reject,ext"`
	SupportedPLMNs            []SupportedPLMNsItem       `aper:"lb:1,ub:MaxnoofSPLMNs,mandatory,reject,ext"`
	GNBCUUPCapacity           *aper.Integer              `aper:"lb:0,ub:255,optional,ignore,ext"`
	TransportLayerAddressInfo *TransportLayerAddressInfo `aper:"optional,ignore,ext"`
}
