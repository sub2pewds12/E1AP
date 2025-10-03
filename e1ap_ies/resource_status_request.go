package e1ap_ies

import "github.com/lvdund/ngap/aper"

// ResourceStatusRequest From: 9_4_4_PDU_Definitions.txt:1340
// ASN.1 Data Type: SEQUENCE
type ResourceStatusRequest struct {
	TransactionID         aper.Integer          `aper:"mandatory,reject,ext"`
	RegistrationRequest   RegistrationRequest   `aper:"mandatory,reject,ext"`
	ReportCharacteristics *aper.BitString       `aper:"lb:36,ub:36,conditional,reject,ext"`
	ReportingPeriodicity  *ReportingPeriodicity `aper:"optional,reject,ext"`
}
