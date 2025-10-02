package e1ap_ies

import "github.com/lvdund/ngap/aper"

// IntegrityProtectionAlgorithm From: 9_4_5_Information_Element_Definitions.txt:1248
const (
	IntegrityProtectionAlgorithmNIA0     aper.Enumerated = 0
	IntegrityProtectionAlgorithmI128NIA1 aper.Enumerated = 1
	IntegrityProtectionAlgorithmI128NIA2 aper.Enumerated = 2
	IntegrityProtectionAlgorithmI128NIA3 aper.Enumerated = 3
)

type IntegrityProtectionAlgorithm struct {
	Value aper.Enumerated
}
