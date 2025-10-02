package e1ap_ies

// E1ReleaseResponse From: 9_4_4_PDU_Definitions.txt:650
type E1ReleaseResponse struct {
	TransactionID int64 `asn1:"mandatory,reject,ext"`
}
