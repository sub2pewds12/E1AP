package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// PDCPReestablishment is a generated ENUMERATED type.
type PDCPReestablishment struct {
	Value aper.Enumerated
}

const (
	PDCPReestablishmentTrue aper.Enumerated = 0
)

func (e *PDCPReestablishment) Encode(w *aper.AperWriter) error {
	// Encode logic for enum PDCPReestablishment to be generated here.
	return fmt.Errorf("Encode not implemented for enum PDCPReestablishment")
}

func (e *PDCPReestablishment) Decode(r *aper.AperReader) error {
	// Decode logic for enum PDCPReestablishment to be generated here.
	return fmt.Errorf("Decode not implemented for enum PDCPReestablishment")
}
