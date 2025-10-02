package e1ap_ies

// GNBCUUPStatusIndication From: 9_4_4_PDU_Definitions.txt:1224
type GNBCUUPStatusIndication struct {
	TransactionID              int64                      `asn1:"mandatory,reject,ext"`
	GNBCUUPOverloadInformation GNBCUUPOverloadInformation `asn1:"mandatory,reject,ext"`
}
