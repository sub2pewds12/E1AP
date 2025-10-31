package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// MDTActivation is a generated ENUMERATED type.
type MDTActivation struct {
	Value aper.Enumerated
}

const (
	MDTActivationImmediateMDTOnly     aper.Enumerated = 0
	MDTActivationImmediateMDTAndTrace aper.Enumerated = 1
)

func (e *MDTActivation) Encode(w *aper.AperWriter) error {
	// Encode logic for enum MDTActivation to be generated here.
	return fmt.Errorf("Encode not implemented for enum MDTActivation")
}

func (e *MDTActivation) Decode(r *aper.AperReader) error {
	// Decode logic for enum MDTActivation to be generated here.
	return fmt.Errorf("Decode not implemented for enum MDTActivation")
}
