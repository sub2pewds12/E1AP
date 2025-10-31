package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// Presence is a generated ENUMERATED type.
type Presence struct {
	Value aper.Enumerated
}

const (
	PresenceOptional    aper.Enumerated = 0
	PresenceConditional aper.Enumerated = 1
	PresenceMandatory   aper.Enumerated = 2
)

func (e *Presence) Encode(w *aper.AperWriter) error {
	// Encode logic for enum Presence to be generated here.
	return fmt.Errorf("Encode not implemented for enum Presence")
}

func (e *Presence) Decode(r *aper.AperReader) error {
	// Decode logic for enum Presence to be generated here.
	return fmt.Errorf("Decode not implemented for enum Presence")
}
