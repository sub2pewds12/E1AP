package e1ap_ies

import "github.com/lvdund/ngap/aper"

// UnsuccessfulOutcome From: 9_4_3_Elementary_Procedure_Definitions.txt:154
// ASN.1 Data Type: SEQUENCE
type UnsuccessfulOutcome struct {
	ProcedureCode aper.Integer     `aper:"lb:0,ub:255,mandatory"`
	Criticality   Criticality      `aper:"mandatory"`
	Value         aper.OctetString `aper:"mandatory"`
}
