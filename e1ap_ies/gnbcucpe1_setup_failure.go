package e1ap_ies

// GNBCUCPE1SetupFailure From: 9_4_4_PDU_Definitions.txt:467
type GNBCUCPE1SetupFailure struct {
	TransactionID          int64                   `asn1:"mandatory,reject,ext"`
	Cause                  Cause                   `asn1:"mandatory,ignore,ext"`
	TimeToWait             *TimeToWait             `asn1:"optional,ignore,ext"`
	CriticalityDiagnostics *CriticalityDiagnostics `asn1:"optional,ignore,ext"`
}
