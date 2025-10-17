package e1ap_ies

import "github.com/lvdund/ngap/aper"

// InitiatingMessage is a generated SEQUENCE type.
type InitiatingMessage struct {
	ProcedureCode ProcedureCode    `aper:"lb:0,ub:255,mandatory"`
	Criticality   Criticality      `aper:"mandatory"`
	Value         aper.OctetString `aper:"mandatory"`
}
