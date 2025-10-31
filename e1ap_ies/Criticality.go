package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// Criticality is a generated ENUMERATED type.
type Criticality struct {
	Value aper.Enumerated
}

const (
	CriticalityReject aper.Enumerated = 0
	CriticalityIgnore aper.Enumerated = 1
	CriticalityNotify aper.Enumerated = 2
)

func (e *Criticality) Encode(w *aper.AperWriter) error {
	// Encode logic for enum Criticality to be generated here.
	return fmt.Errorf("Encode not implemented for enum Criticality")
}

func (e *Criticality) Decode(r *aper.AperReader) error {
	// Decode logic for enum Criticality to be generated here.
	return fmt.Errorf("Decode not implemented for enum Criticality")
}
