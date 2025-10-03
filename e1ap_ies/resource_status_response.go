package e1ap_ies

import "github.com/lvdund/ngap/aper"

// ResourceStatusResponse From: 9_4_4_PDU_Definitions.txt:1362
// ASN.1 Data Type: SEQUENCE
type ResourceStatusResponse struct {
	TransactionID          aper.Integer            `aper:"mandatory,reject,ext"`
	CriticalityDiagnostics *CriticalityDiagnostics `aper:"optional,ignore,ext"`
}
