package e1ap_ies

// CriticalityDiagnostics is a generated SEQUENCE type.
type CriticalityDiagnostics struct {
	ProcedureCode             *ProcedureCode                 `aper:"lb:0,ub:255,optional,ext"`
	TriggeringMessage         *TriggeringMessage             `aper:"optional,ext"`
	ProcedureCriticality      *Criticality                   `aper:"optional,ext"`
	TransactionID             TransactionID                  `aper:"lb:0,ub:255,mandatory,reject,ext"`
	IEsCriticalityDiagnostics []CriticalityDiagnosticsIEItem `aper:"optional,ext"`
	IEExtensions              *ProtocolExtensionContainer    `aper:"optional,ext"`
}
