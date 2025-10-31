package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// ResetAll is a generated ENUMERATED type.
type ResetAll struct {
	Value aper.Enumerated
}

const (
	ResetAllResetAll aper.Enumerated = 0
)

func (e *ResetAll) Encode(w *aper.AperWriter) error {
	// Encode logic for enum ResetAll to be generated here.
	return fmt.Errorf("Encode not implemented for enum ResetAll")
}

func (e *ResetAll) Decode(r *aper.AperReader) error {
	// Decode logic for enum ResetAll to be generated here.
	return fmt.Errorf("Decode not implemented for enum ResetAll")
}
