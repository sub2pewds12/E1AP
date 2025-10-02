package e1ap_ies

// ResourceStatusResponse From: 9_4_4_PDU_Definitions.txt:1362
type ResourceStatusResponse struct {
	TransactionID          int64                   `asn1:"mandatory,reject,ext"`
	CriticalityDiagnostics *CriticalityDiagnostics `asn1:"optional,ignore,ext"`
}
