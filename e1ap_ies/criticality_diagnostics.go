package e1ap_ies

import "github.com/lvdund/ngap/aper"

// CriticalityDiagnostics is a generated SEQUENCE type.
type CriticalityDiagnostics struct {
	ProcedureCode             *aper.Integer                  `aper:"lb:0,ub:255,optional,ext"`
	TriggeringMessage         *TriggeringMessage             `aper:"optional,ext"`
	ProcedureCriticality      *Criticality                   `aper:"optional,ext"`
	TransactionID             *aper.Integer                  `aper:"lb:0,ub:255,optional,reject,ext"`
	IEsCriticalityDiagnostics []CriticalityDiagnosticsIEItem `aper:"optional,ext"`
}
