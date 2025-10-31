package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// RATType is a generated ENUMERATED type.
type RATType struct {
	Value aper.Enumerated
}

const (
	RATTypeEUTRA aper.Enumerated = 0
	RATTypeNR    aper.Enumerated = 1
)

func (e *RATType) Encode(w *aper.AperWriter) error {
	// Encode logic for enum RATType to be generated here.
	return fmt.Errorf("Encode not implemented for enum RATType")
}

func (e *RATType) Decode(r *aper.AperReader) error {
	// Decode logic for enum RATType to be generated here.
	return fmt.Errorf("Decode not implemented for enum RATType")
}
