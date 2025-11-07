package e1ap

import (
	"bytes"
	"fmt"

	"github.com/lvdund/ngap/aper"
	"github.com/sub2pewds12/E1AP/e1ap_ies"
)

func E1apDecode(buf []byte) (pdu E1AP_PDU, err error, diagnostics []e1ap_ies.CriticalityDiagnosticsIEItem) {
	r := aper.NewReader(bytes.NewBuffer(buf))

	choice, err := r.ReadChoice(2, false)
	if err != nil {
		return
	}
	pdu.Present = uint8(choice)

	v, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, false)
	if err != nil {
		return
	}
	var procedureCode e1ap_ies.ProcedureCode = e1ap_ies.ProcedureCode{Value: aper.Integer(v)}

	e, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false)
	if err != nil {
		return
	}
	var criticality e1ap_ies.Criticality = e1ap_ies.Criticality{Value: aper.Enumerated(e)}

	var containerBytes []byte
	if containerBytes, err = r.ReadOpenType(); err != nil {
		return
	}

	message := createMessage(pdu.Present, procedureCode)
	if message == nil {
		err = fmt.Errorf("unknown message: procedureCode=%d, present=%d", procedureCode.Value, pdu.Present)
		return
	}

	if err, diagnostics = message.Decode(containerBytes); err != nil {
		return
	}

	pdu.Message = &InitiatingMessage{
		ProcedureCode: procedureCode,
		Criticality:   criticality,
		Value:         message,
	}

	return
}
