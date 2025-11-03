// e1ap/decode.go
package e1ap

import (
	"bytes"
	"fmt"

	"github.com/lvdund/ngap/aper"
	"github.com/sub2pewds12/E1AP/e1ap_ies"
)

// E1apDecode is the main entry point for decoding a raw E1AP message.
func E1apDecode(buf []byte) (pdu E1AP_PDU, err error, diagnostics []e1ap_ies.CriticalityDiagnosticsIEItem) {
	r := aper.NewReader(bytes.NewBuffer(buf))

	// 1. Decode the PDU CHOICE index (Initiating, Successful, or Unsuccessful)
	choice, err := r.ReadChoice(2, false) // E1AP PDU is not extensible
	if err != nil {
		return
	}
	pdu.Present = uint8(choice)

	// 2. Decode the Procedure Code
	v, err := r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, false)
	if err != nil {
		return
	}
	var procedureCode e1ap_ies.ProcedureCode = e1ap_ies.ProcedureCode{Value: aper.Integer(v)}

	// 3. Decode the Criticality
	e, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false)
	if err != nil {
		return
	}
	var criticality e1ap_ies.Criticality = e1ap_ies.Criticality{Value: aper.Enumerated(e)}

	// 4. Decode the message content (the Value) as an Open Type
	var containerBytes []byte
	if containerBytes, err = r.ReadOpenType(); err != nil {
		return
	}

	// 5. Use the factory to get an empty struct of the correct PDU type
	message := createMessage(pdu.Present, procedureCode)
	if message == nil {
		err = fmt.Errorf("unknown message: procedureCode=%d, present=%d", procedureCode.Value, pdu.Present)
		return
	}

	// 6. Call the PDU's specific Decode method to parse the containerBytes
	if err, diagnostics = message.Decode(containerBytes); err != nil {
		return
	}

	// 7. Assemble the final PDU struct
	pdu.Message = &InitiatingMessage{ // This should be more specific
		ProcedureCode: procedureCode,
		Criticality:   criticality,
		Value:         message,
	}
	// A more complete implementation would switch on pdu.Present to set the correct
	// wrapper type (InitiatingMessage, SuccessfulOutcome, etc.) in pdu.Message

	return
}
