package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// ROHCContinueROHC is a generated ENUMERATED type.
type ROHCContinueROHC struct {
	Value aper.Enumerated
}

const (
	ROHCContinueROHCTrue aper.Enumerated = 0
)

func (e *ROHCContinueROHC) Encode(w *aper.AperWriter) error {
	// Encode logic for enum ROHCContinueROHC to be generated here.
	return fmt.Errorf("Encode not implemented for enum ROHCContinueROHC")
}

func (e *ROHCContinueROHC) Decode(r *aper.AperReader) error {
	// Decode logic for enum ROHCContinueROHC to be generated here.
	return fmt.Errorf("Decode not implemented for enum ROHCContinueROHC")
}
