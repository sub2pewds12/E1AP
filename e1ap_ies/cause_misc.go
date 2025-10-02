package e1ap_ies

import "github.com/lvdund/ngap/aper"

// CauseMisc From: 9_4_5_Information_Element_Definitions.txt:153
const (
	CauseMiscControlProcessingOverload             aper.Enumerated = 0
	CauseMiscNotEnoughUserPlaneProcessingResources aper.Enumerated = 1
	CauseMiscHardwareFailure                       aper.Enumerated = 2
	CauseMiscOmIntervention                        aper.Enumerated = 3
	CauseMiscUnspecified                           aper.Enumerated = 4
)

type CauseMisc struct {
	Value aper.Enumerated
}
