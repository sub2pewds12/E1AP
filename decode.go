package e1ap

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
	"github.com/sub2pewds12/E1AP/e1ap_ies"
)

func E1apDecode(buf []byte) (pdu E1AP_PDU, diagnostics []e1ap_ies.CriticalityDiagnosticsIEItem, err error) {
	r := per.NewDecoder(buf, per.APER)

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			{Name: "initiatingMessage", Tag: 0},
			{Name: "successfulOutcome", Tag: 1},
			{Name: "unsuccessfulOutcome", Tag: 2},
		},
	}
	choiceDecoder := r.NewChoiceDecoder(c)
	choice, _, _, err := choiceDecoder.DecodeChoice()
	if err != nil {
		return
	}
	pdu.Present = uint8(choice)

	v, err := r.DecodeInteger(per.Constrained(0, 255))
	if err != nil {
		return
	}
	var procedureCode e1ap_ies.ProcedureCode = e1ap_ies.ProcedureCode{Value: v}

	enumC := per.EnumeratedConstraints{Extensible: false, RootValues: make([]int64, 3)}
	e, err := r.DecodeEnumerated(enumC)
	if err != nil {
		return
	}
	var criticality e1ap_ies.Criticality = e1ap_ies.Criticality{Value: e}

	var containerBytes []byte
	if containerBytes, err = r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: nil, Max: nil}); err != nil {
		return
	}

	message := createMessage(pdu.Present, procedureCode)
	if message == nil {
		err = fmt.Errorf("unknown message: procedureCode=%d, present=%d", procedureCode.Value, pdu.Present)
		return
	}

	if diagnostics, err = message.Decode(containerBytes); err != nil {
		return
	}

	switch pdu.Present {
	case PduChoiceInitiatingMessage:
		pdu.Message = &InitiatingMessage{
			ProcedureCode: procedureCode,
			Criticality:   criticality,
			Value:         message,
		}
	case PduChoiceSuccessfulOutcome:
		pdu.Message = &SuccessfulOutcome{
			ProcedureCode: procedureCode,
			Criticality:   criticality,
			Value:         message,
		}
	case PduChoiceUnsuccessfulOutcome:
		pdu.Message = &UnsuccessfulOutcome{
			ProcedureCode: procedureCode,
			Criticality:   criticality,
			Value:         message,
		}
	}

	return
}