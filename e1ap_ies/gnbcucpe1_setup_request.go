package e1ap_ies

// GNBCUCPE1SetupRequest From: 9_4_4_PDU_Definitions.txt:427
type GNBCUCPE1SetupRequest struct {
	TransactionID             int64                      `asn1:"mandatory,reject,ext"`
	GNBCUCPName               *[]byte                    `asn1:"optional,ignore,ext"`
	TransportLayerAddressInfo *TransportLayerAddressInfo `asn1:"optional,ignore,ext"`
}
