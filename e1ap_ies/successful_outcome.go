package e1ap_ies

import "github.com/lvdund/ngap/aper"

// SuccessfulOutcome is a generated SEQUENCE type.
type SuccessfulOutcome struct {
	ProcedureCode aper.Integer     `aper:"lb:0,ub:255,mandatory"`
	Criticality   Criticality      `aper:"mandatory"`
	Value         aper.OctetString `aper:"mandatory"`
}
