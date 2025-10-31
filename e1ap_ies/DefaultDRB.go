package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DefaultDRB is a generated ENUMERATED type.
type DefaultDRB struct {
	Value aper.Enumerated
}

const (
	DefaultDRBTrue  aper.Enumerated = 0
	DefaultDRBFalse aper.Enumerated = 1
)

func (e *DefaultDRB) Encode(w *aper.AperWriter) error {
	// Encode logic for enum DefaultDRB to be generated here.
	return fmt.Errorf("Encode not implemented for enum DefaultDRB")
}

func (e *DefaultDRB) Decode(r *aper.AperReader) error {
	// Decode logic for enum DefaultDRB to be generated here.
	return fmt.Errorf("Decode not implemented for enum DefaultDRB")
}
