package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// M4period is a generated ENUMERATED type.
type M4period struct {
	Value aper.Enumerated
}

const (
	M4periodMs1024  aper.Enumerated = 0
	M4periodMs2048  aper.Enumerated = 1
	M4periodMs5120  aper.Enumerated = 2
	M4periodMs10240 aper.Enumerated = 3
	M4periodMin1    aper.Enumerated = 4
)

func (e *M4period) Encode(w *aper.AperWriter) error {
	// Encode logic for enum M4period to be generated here.
	return fmt.Errorf("Encode not implemented for enum M4period")
}

func (e *M4period) Decode(r *aper.AperReader) error {
	// Decode logic for enum M4period to be generated here.
	return fmt.Errorf("Decode not implemented for enum M4period")
}
