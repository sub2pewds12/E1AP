package e1ap_ies

// ResourceStatusFailure From: 9_4_4_PDU_Definitions.txt:1381
type ResourceStatusFailure struct {
	TransactionID          int64                   `asn1:"mandatory,reject,ext"`
	Cause                  Cause                   `asn1:"mandatory,ignore,ext"`
	CriticalityDiagnostics *CriticalityDiagnostics `asn1:"optional,ignore,ext"`
}
