package e1ap_ies

// GNBCUCPConfigurationUpdate From: 9_4_4_PDU_Definitions.txt:559
type GNBCUCPConfigurationUpdate struct {
	TransactionID             int64                      `asn1:"mandatory,reject,ext"`
	GNBCUCPName               *[]byte                    `asn1:"optional,ignore,ext"`
	GNBCUCPTNLAToAddList      []GNBCUCPTNLAToAddItem     `asn1:"optional,ignore,ext"`
	GNBCUCPTNLAToRemoveList   []GNBCUCPTNLAToRemoveItem  `asn1:"optional,ignore,ext"`
	GNBCUCPTNLAToUpdateList   []GNBCUCPTNLAToUpdateItem  `asn1:"optional,ignore,ext"`
	TransportLayerAddressInfo *TransportLayerAddressInfo `asn1:"optional,ignore,ext"`
}
