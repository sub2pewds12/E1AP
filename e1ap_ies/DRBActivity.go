package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DRBActivity is a generated ENUMERATED type.
type DRBActivity struct {
	Value aper.Enumerated
}

const (
	DRBActivityActive    aper.Enumerated = 0
	DRBActivityNotActive aper.Enumerated = 1
)

func (e *DRBActivity) Encode(w *aper.AperWriter) error {
	// Encode logic for enum DRBActivity to be generated here.
	return fmt.Errorf("Encode not implemented for enum DRBActivity")
}

func (e *DRBActivity) Decode(r *aper.AperReader) error {
	// Decode logic for enum DRBActivity to be generated here.
	return fmt.Errorf("Decode not implemented for enum DRBActivity")
}
