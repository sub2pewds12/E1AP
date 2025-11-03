package e1ap

import (
	"io"

	"github.com/sub2pewds12/E1AP/e1ap_ies"
)

const (
	PduChoiceInitiatingMessage   uint8 = 0
	PduChoiceSuccessfulOutcome   uint8 = 1
	PduChoiceUnsuccessfulOutcome uint8 = 2
)

type MessageEncoder interface {
	Encode(io.Writer) error
}

type MessageUnmarshaller interface {
	Decode([]byte) (error, []e1ap_ies.CriticalityDiagnosticsIEItem)
}

type E1AP_PDU struct {
	Present uint8
	Message interface{}
}

type InitiatingMessage struct {
	ProcedureCode e1ap_ies.ProcedureCode
	Criticality   e1ap_ies.Criticality
	Value         MessageUnmarshaller
}

type SuccessfulOutcome struct {
	ProcedureCode e1ap_ies.ProcedureCode
	Criticality   e1ap_ies.Criticality
	Value         MessageUnmarshaller
}

type UnsuccessfulOutcome struct {
	ProcedureCode e1ap_ies.ProcedureCode
	Criticality   e1ap_ies.Criticality
	Value         MessageUnmarshaller
}
