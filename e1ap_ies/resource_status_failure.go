package e1ap_ies

import "github.com/lvdund/ngap/aper"

// ResourceStatusFailure From: 9_4_4_PDU_Definitions.txt:1381
// ASN.1 Data Type: SEQUENCE
type ResourceStatusFailure struct {
	TransactionID          aper.Integer            `aper:"mandatory,reject,ext"`
	Cause                  Cause                   `aper:"mandatory,ignore,ext"`
	CriticalityDiagnostics *CriticalityDiagnostics `aper:"optional,ignore,ext"`
}
