package e1ap_ies

// GNBCUUPConfigurationUpdate From: 9_4_4_PDU_Definitions.txt:492
type GNBCUUPConfigurationUpdate struct {
	TransactionID             int64                      `asn1:"mandatory,reject,ext"`
	GNBCUUPID                 int64                      `asn1:"lb:0,ub:68719476735,mandatory,reject,ext"`
	GNBCUUPName               *[]byte                    `asn1:"optional,ignore,ext"`
	SupportedPLMNs            []SupportedPLMNsItem       `asn1:"lb:1,ub:MaxnoofSPLMNs,optional,reject,ext"`
	GNBCUUPCapacity           *int64                     `asn1:"lb:0,ub:255,optional,ignore,ext"`
	GNBCUUPTNLAToRemoveList   []GNBCUUPTNLAToRemoveItem  `asn1:"optional,reject,ext"`
	TransportLayerAddressInfo *TransportLayerAddressInfo `asn1:"optional,ignore,ext"`
}
