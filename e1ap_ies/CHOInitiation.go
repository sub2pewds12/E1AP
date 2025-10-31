package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// CHOInitiation is a generated ENUMERATED type.
type CHOInitiation struct {
	Value aper.Enumerated
}

const (
	CHOInitiationTrue aper.Enumerated = 0
)

func (e *CHOInitiation) Encode(w *aper.AperWriter) error {
	// Encode logic for enum CHOInitiation to be generated here.
	return fmt.Errorf("Encode not implemented for enum CHOInitiation")
}

func (e *CHOInitiation) Decode(r *aper.AperReader) error {
	// Decode logic for enum CHOInitiation to be generated here.
	return fmt.Errorf("Decode not implemented for enum CHOInitiation")
}
