package e1ap_ies

import "github.com/lvdund/ngap/aper"

// UnsuccessfulOutcome is a generated SEQUENCE type.
type UnsuccessfulOutcome struct {
	ProcedureCode aper.Integer     `aper:"lb:0,ub:255,mandatory"`
	Criticality   Criticality      `aper:"mandatory"`
	Value         aper.OctetString `aper:"mandatory"`
}
