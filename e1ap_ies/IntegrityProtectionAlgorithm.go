package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// IntegrityProtectionAlgorithm is a generated ENUMERATED type.
type IntegrityProtectionAlgorithm struct {
	Value aper.Enumerated
}

const (
	IntegrityProtectionAlgorithmNIA0     aper.Enumerated = 0
	IntegrityProtectionAlgorithmI128NIA1 aper.Enumerated = 1
	IntegrityProtectionAlgorithmI128NIA2 aper.Enumerated = 2
	IntegrityProtectionAlgorithmI128NIA3 aper.Enumerated = 3
)

func (e *IntegrityProtectionAlgorithm) Encode(w *aper.AperWriter) error {
	// Encode logic for enum IntegrityProtectionAlgorithm to be generated here.
	return fmt.Errorf("Encode not implemented for enum IntegrityProtectionAlgorithm")
}

func (e *IntegrityProtectionAlgorithm) Decode(r *aper.AperReader) error {
	// Decode logic for enum IntegrityProtectionAlgorithm to be generated here.
	return fmt.Errorf("Decode not implemented for enum IntegrityProtectionAlgorithm")
}
