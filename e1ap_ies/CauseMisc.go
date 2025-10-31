package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// CauseMisc is a generated ENUMERATED type.
type CauseMisc struct {
	Value aper.Enumerated
}

const (
	CauseMiscControlProcessingOverload             aper.Enumerated = 0
	CauseMiscNotEnoughUserPlaneProcessingResources aper.Enumerated = 1
	CauseMiscHardwareFailure                       aper.Enumerated = 2
	CauseMiscOmIntervention                        aper.Enumerated = 3
	CauseMiscUnspecified                           aper.Enumerated = 4
)

func (e *CauseMisc) Encode(w *aper.AperWriter) error {
	// Encode logic for enum CauseMisc to be generated here.
	return fmt.Errorf("Encode not implemented for enum CauseMisc")
}

func (e *CauseMisc) Decode(r *aper.AperReader) error {
	// Decode logic for enum CauseMisc to be generated here.
	return fmt.Errorf("Decode not implemented for enum CauseMisc")
}
