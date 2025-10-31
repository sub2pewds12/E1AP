package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// PDCPSNSize is a generated ENUMERATED type.
type PDCPSNSize struct {
	Value aper.Enumerated
}

const (
	PDCPSNSizeS12 aper.Enumerated = 0
	PDCPSNSizeS18 aper.Enumerated = 1
)

func (e *PDCPSNSize) Encode(w *aper.AperWriter) error {
	// Encode logic for enum PDCPSNSize to be generated here.
	return fmt.Errorf("Encode not implemented for enum PDCPSNSize")
}

func (e *PDCPSNSize) Decode(r *aper.AperReader) error {
	// Decode logic for enum PDCPSNSize to be generated here.
	return fmt.Errorf("Decode not implemented for enum PDCPSNSize")
}
