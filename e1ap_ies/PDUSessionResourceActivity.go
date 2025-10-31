package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// PDUSessionResourceActivity is a generated ENUMERATED type.
type PDUSessionResourceActivity struct {
	Value aper.Enumerated
}

const (
	PDUSessionResourceActivityActive    aper.Enumerated = 0
	PDUSessionResourceActivityNotActive aper.Enumerated = 1
)

func (e *PDUSessionResourceActivity) Encode(w *aper.AperWriter) error {
	// Encode logic for enum PDUSessionResourceActivity to be generated here.
	return fmt.Errorf("Encode not implemented for enum PDUSessionResourceActivity")
}

func (e *PDUSessionResourceActivity) Decode(r *aper.AperReader) error {
	// Decode logic for enum PDUSessionResourceActivity to be generated here.
	return fmt.Errorf("Decode not implemented for enum PDUSessionResourceActivity")
}
