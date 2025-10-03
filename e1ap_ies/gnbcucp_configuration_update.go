package e1ap_ies

import "github.com/lvdund/ngap/aper"

// GNBCUCPConfigurationUpdate From: 9_4_4_PDU_Definitions.txt:559
// ASN.1 Data Type: SEQUENCE
type GNBCUCPConfigurationUpdate struct {
	TransactionID             aper.Integer               `aper:"mandatory,reject,ext"`
	GNBCUCPName               *aper.OctetString          `aper:"optional,ignore,ext"`
	GNBCUCPTNLAToAddList      []GNBCUCPTNLAToAddItem     `aper:"optional,ignore,ext"`
	GNBCUCPTNLAToRemoveList   []GNBCUCPTNLAToRemoveItem  `aper:"optional,ignore,ext"`
	GNBCUCPTNLAToUpdateList   []GNBCUCPTNLAToUpdateItem  `aper:"optional,ignore,ext"`
	TransportLayerAddressInfo *TransportLayerAddressInfo `aper:"optional,ignore,ext"`
}
