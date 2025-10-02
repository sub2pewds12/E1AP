package e1ap_ies

// CriticalityDiagnostics From: 9_4_5_Information_Element_Definitions.txt:277
type CriticalityDiagnostics struct {
	ProcedureCode             *int64                         `asn1:"lb:0,ub:255,optional,ext"`
	TriggeringMessage         *TriggeringMessage             `asn1:"optional,ext"`
	ProcedureCriticality      *Criticality                   `asn1:"optional,ext"`
	TransactionID             *int64                         `asn1:"optional,ext"`
	IEsCriticalityDiagnostics []CriticalityDiagnosticsIEItem `asn1:"optional,ext"`
}
