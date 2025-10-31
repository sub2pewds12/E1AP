package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// MaxIPrate is a generated ENUMERATED type.
type MaxIPrate struct {
	Value aper.Enumerated
}

const (
	MaxIPrateBitrate64kbs aper.Enumerated = 0
	MaxIPrateMaxUErate    aper.Enumerated = 1
)

func (e *MaxIPrate) Encode(w *aper.AperWriter) error {
	// Encode logic for enum MaxIPrate to be generated here.
	return fmt.Errorf("Encode not implemented for enum MaxIPrate")
}

func (e *MaxIPrate) Decode(r *aper.AperReader) error {
	// Decode logic for enum MaxIPrate to be generated here.
	return fmt.Errorf("Decode not implemented for enum MaxIPrate")
}
