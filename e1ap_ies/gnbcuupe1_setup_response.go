package e1ap_ies

// GNBCUUPE1SetupResponse From: 9_4_4_PDU_Definitions.txt:384
type GNBCUUPE1SetupResponse struct {
	TransactionID             int64                      `asn1:"mandatory,reject,ext"`
	GNBCUCPName               *[]byte                    `asn1:"optional,ignore,ext"`
	TransportLayerAddressInfo *TransportLayerAddressInfo `asn1:"optional,ignore,ext"`
}
