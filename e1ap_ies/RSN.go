package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// RSN is a generated ENUMERATED type.
type RSN struct {
	Value aper.Enumerated
}

const (
	RSNV1 aper.Enumerated = 0
	RSNV2 aper.Enumerated = 1
)

func (e *RSN) Encode(w *aper.AperWriter) error {
	// Encode logic for enum RSN to be generated here.
	return fmt.Errorf("Encode not implemented for enum RSN")
}

func (e *RSN) Decode(r *aper.AperReader) error {
	// Decode logic for enum RSN to be generated here.
	return fmt.Errorf("Decode not implemented for enum RSN")
}
