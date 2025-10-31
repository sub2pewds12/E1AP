package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// IntegrityProtectionResult is a generated ENUMERATED type.
type IntegrityProtectionResult struct {
	Value aper.Enumerated
}

const (
	IntegrityProtectionResultPerformed    aper.Enumerated = 0
	IntegrityProtectionResultNotPerformed aper.Enumerated = 1
)

func (e *IntegrityProtectionResult) Encode(w *aper.AperWriter) error {
	// Encode logic for enum IntegrityProtectionResult to be generated here.
	return fmt.Errorf("Encode not implemented for enum IntegrityProtectionResult")
}

func (e *IntegrityProtectionResult) Decode(r *aper.AperReader) error {
	// Decode logic for enum IntegrityProtectionResult to be generated here.
	return fmt.Errorf("Decode not implemented for enum IntegrityProtectionResult")
}
