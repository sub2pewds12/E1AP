package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// ConfidentialityProtectionResult is a generated ENUMERATED type.
type ConfidentialityProtectionResult struct {
	Value aper.Enumerated
}

const (
	ConfidentialityProtectionResultPerformed    aper.Enumerated = 0
	ConfidentialityProtectionResultNotPerformed aper.Enumerated = 1
)

func (e *ConfidentialityProtectionResult) Encode(w *aper.AperWriter) error {
	// Encode logic for enum ConfidentialityProtectionResult to be generated here.
	return fmt.Errorf("Encode not implemented for enum ConfidentialityProtectionResult")
}

func (e *ConfidentialityProtectionResult) Decode(r *aper.AperReader) error {
	// Decode logic for enum ConfidentialityProtectionResult to be generated here.
	return fmt.Errorf("Decode not implemented for enum ConfidentialityProtectionResult")
}
