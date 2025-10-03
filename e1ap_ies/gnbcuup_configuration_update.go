package e1ap_ies

import "github.com/lvdund/ngap/aper"

// GNBCUUPConfigurationUpdate From: 9_4_4_PDU_Definitions.txt:492
// ASN.1 Data Type: SEQUENCE
type GNBCUUPConfigurationUpdate struct {
	TransactionID             aper.Integer               `aper:"mandatory,reject,ext"`
	GNBCUUPID                 aper.Integer               `aper:"lb:0,ub:68719476735,mandatory,reject,ext"`
	GNBCUUPName               *aper.OctetString          `aper:"optional,ignore,ext"`
	SupportedPLMNs            []SupportedPLMNsItem       `aper:"lb:1,ub:MaxnoofSPLMNs,optional,reject,ext"`
	GNBCUUPCapacity           *aper.Integer              `aper:"lb:0,ub:255,optional,ignore,ext"`
	GNBCUUPTNLAToRemoveList   []GNBCUUPTNLAToRemoveItem  `aper:"optional,reject,ext"`
	TransportLayerAddressInfo *TransportLayerAddressInfo `aper:"optional,ignore,ext"`
}
