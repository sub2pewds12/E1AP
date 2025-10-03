package e1ap_ies

import "github.com/lvdund/ngap/aper"

// GNBCUCPE1SetupFailure From: 9_4_4_PDU_Definitions.txt:467
// ASN.1 Data Type: SEQUENCE
type GNBCUCPE1SetupFailure struct {
	TransactionID          aper.Integer            `aper:"mandatory,reject,ext"`
	Cause                  Cause                   `aper:"mandatory,ignore,ext"`
	TimeToWait             *TimeToWait             `aper:"optional,ignore,ext"`
	CriticalityDiagnostics *CriticalityDiagnostics `aper:"optional,ignore,ext"`
}
