package e1ap_ies

import (
	"bytes"
	"io"

	"github.com/lvdund/ngap/aper" // Your aper path
)

// PDU Choice constants
const (
	E1apPduInitiatingMessage   uint8 = 0
	E1apPduSuccessfulOutcome   uint8 = 1
	E1apPduUnsuccessfulOutcome uint8 = 2
)

func encodeMessage(w io.Writer, present uint8, procedureCode ProcedureCode, criticality Criticality, ies []E1APMessageIE) (err error) {
	aw := aper.NewWriter(w)

	// The APER spec for the PDU itself might not have this first bit, let's assume it's just the choice.
	// if err = aw.WriteBool(false); err != nil { return }

	if err = aw.WriteChoice(uint64(present), 2, false); err != nil { // Extensibility might be false
		return
	}
	if err = procedureCode.Encode(aw); err != nil {
		return
	}
	if err = criticality.Encode(aw); err != nil {
		return
	}

	// Encode IEs to a temporary buffer
	var buf bytes.Buffer
	cW := aper.NewWriter(&buf)

	// Create a slice of interfaces for the generic function
	ieInterfaces := make([]aper.IE, len(ies))
	for i := range ies {
		ieInterfaces[i] = &ies[i]
	}

	if err = aper.WriteSequenceOf(ieInterfaces, cW, &aper.Constraint{Lb: 0, Ub: 65535}, false); err != nil {
		return
	}
	if err = cW.Close(); err != nil {
		return
	}

	// Write the buffer as an Open Type
	if err = aw.WriteOpenType(buf.Bytes()); err != nil {
		return
	}
	return aw.Close()
}
