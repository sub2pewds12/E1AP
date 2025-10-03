package e1ap_ies

import "github.com/lvdund/ngap/aper"

// InitiatingMessage From: 9_4_3_Elementary_Procedure_Definitions.txt:142
// ASN.1 Data Type: SEQUENCE
type InitiatingMessage struct {
	ProcedureCode aper.Integer     `aper:"lb:0,ub:255,mandatory"`
	Criticality   Criticality      `aper:"mandatory"`
	Value         aper.OctetString `aper:"mandatory"`
}
