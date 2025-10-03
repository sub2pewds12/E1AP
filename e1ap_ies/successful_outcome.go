package e1ap_ies

import "github.com/lvdund/ngap/aper"

// SuccessfulOutcome From: 9_4_3_Elementary_Procedure_Definitions.txt:148
// ASN.1 Data Type: SEQUENCE
type SuccessfulOutcome struct {
	ProcedureCode aper.Integer     `aper:"lb:0,ub:255,mandatory"`
	Criticality   Criticality      `aper:"mandatory"`
	Value         aper.OctetString `aper:"mandatory"`
}
