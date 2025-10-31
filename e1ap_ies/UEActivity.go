package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// UEActivity is a generated ENUMERATED type.
type UEActivity struct {
	Value aper.Enumerated
}

const (
	UEActivityActive    aper.Enumerated = 0
	UEActivityNotActive aper.Enumerated = 1
)

func (e *UEActivity) Encode(w *aper.AperWriter) error {
	// Encode logic for enum UEActivity to be generated here.
	return fmt.Errorf("Encode not implemented for enum UEActivity")
}

func (e *UEActivity) Decode(r *aper.AperReader) error {
	// Decode logic for enum UEActivity to be generated here.
	return fmt.Errorf("Decode not implemented for enum UEActivity")
}
