package e1ap_ies

import "github.com/lvdund/ngap/aper"

// CriticalityDiagnostics From: 9_4_5_Information_Element_Definitions.txt:277
// ASN.1 Data Type: SEQUENCE
type CriticalityDiagnostics struct {
	ProcedureCode             *aper.Integer                  `aper:"lb:0,ub:255,optional,ext"`
	TriggeringMessage         *TriggeringMessage             `aper:"optional,ext"`
	ProcedureCriticality      *Criticality                   `aper:"optional,ext"`
	TransactionID             *aper.Integer                  `aper:"optional,reject,ext"`
	IEsCriticalityDiagnostics []CriticalityDiagnosticsIEItem `aper:"optional,ext"`
}
