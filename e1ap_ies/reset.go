package e1ap_ies

// Reset From: 9_4_4_PDU_Definitions.txt:249
type Reset struct {
	TransactionID int64     `asn1:"mandatory,reject,ext"`
	Cause         Cause     `asn1:"mandatory,ignore,ext"`
	ResetType     ResetType `asn1:"mandatory,reject,ext"`
}
