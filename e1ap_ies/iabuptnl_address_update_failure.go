package e1ap_ies

import "github.com/lvdund/ngap/aper"

// IABUPTNLAddressUpdateFailure From: 9_4_4_PDU_Definitions.txt:1470
// ASN.1 Data Type: SEQUENCE
type IABUPTNLAddressUpdateFailure struct {
	TransactionID          aper.Integer            `aper:"mandatory,reject,ext"`
	Cause                  Cause                   `aper:"mandatory,ignore,ext"`
	TimeToWait             *TimeToWait             `aper:"optional,ignore,ext"`
	CriticalityDiagnostics *CriticalityDiagnostics `aper:"optional,ignore,ext"`
}
