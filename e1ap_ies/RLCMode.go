package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// RLCMode is a generated ENUMERATED type.
type RLCMode struct {
	Value aper.Enumerated
}

const (
	RLCModeRlcTm                 aper.Enumerated = 0
	RLCModeRlcAm                 aper.Enumerated = 1
	RLCModeRlcUmBidirectional    aper.Enumerated = 2
	RLCModeRlcUmUnidirectionalUl aper.Enumerated = 3
	RLCModeRlcUmUnidirectionalDl aper.Enumerated = 4
)

func (e *RLCMode) Encode(w *aper.AperWriter) error {
	// Encode logic for enum RLCMode to be generated here.
	return fmt.Errorf("Encode not implemented for enum RLCMode")
}

func (e *RLCMode) Decode(r *aper.AperReader) error {
	// Decode logic for enum RLCMode to be generated here.
	return fmt.Errorf("Decode not implemented for enum RLCMode")
}
