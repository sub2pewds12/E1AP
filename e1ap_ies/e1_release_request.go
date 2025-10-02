package e1ap_ies

// E1ReleaseRequest From: 9_4_4_PDU_Definitions.txt:633
type E1ReleaseRequest struct {
	TransactionID int64 `asn1:"mandatory,reject,ext"`
	Cause         Cause `asn1:"mandatory,ignore,ext"`
}
