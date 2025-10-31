package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DLTXStop is a generated ENUMERATED type.
type DLTXStop struct {
	Value aper.Enumerated
}

const (
	DLTXStopStop   aper.Enumerated = 0
	DLTXStopResume aper.Enumerated = 1
)

func (e *DLTXStop) Encode(w *aper.AperWriter) error {
	// Encode logic for enum DLTXStop to be generated here.
	return fmt.Errorf("Encode not implemented for enum DLTXStop")
}

func (e *DLTXStop) Decode(r *aper.AperReader) error {
	// Decode logic for enum DLTXStop to be generated here.
	return fmt.Errorf("Decode not implemented for enum DLTXStop")
}
