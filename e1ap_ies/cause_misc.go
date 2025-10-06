package e1ap_ies

import (
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
