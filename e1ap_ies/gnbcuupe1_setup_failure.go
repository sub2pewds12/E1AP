package e1ap_ies

// GNBCUUPE1SetupFailure From: 9_4_4_PDU_Definitions.txt:402
type GNBCUUPE1SetupFailure struct {
	TransactionID          int64                   `asn1:"mandatory,reject,ext"`
	Cause                  Cause                   `asn1:"mandatory,ignore,ext"`
	TimeToWait             *TimeToWait             `asn1:"optional,ignore,ext"`
	CriticalityDiagnostics *CriticalityDiagnostics `asn1:"optional,ignore,ext"`
}
