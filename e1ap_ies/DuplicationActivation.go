package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DuplicationActivation is a generated ENUMERATED type.
type DuplicationActivation struct {
	Value aper.Enumerated
}

const (
	DuplicationActivationActive   aper.Enumerated = 0
	DuplicationActivationInactive aper.Enumerated = 1
)

func (e *DuplicationActivation) Encode(w *aper.AperWriter) error {
	// Encode logic for enum DuplicationActivation to be generated here.
	return fmt.Errorf("Encode not implemented for enum DuplicationActivation")
}

func (e *DuplicationActivation) Decode(r *aper.AperReader) error {
	// Decode logic for enum DuplicationActivation to be generated here.
	return fmt.Errorf("Decode not implemented for enum DuplicationActivation")
}
