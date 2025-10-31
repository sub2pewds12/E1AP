package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// ConfidentialityProtectionIndication is a generated ENUMERATED type.
type ConfidentialityProtectionIndication struct {
	Value aper.Enumerated
}

const (
	ConfidentialityProtectionIndicationRequired  aper.Enumerated = 0
	ConfidentialityProtectionIndicationPreferred aper.Enumerated = 1
	ConfidentialityProtectionIndicationNotNeeded aper.Enumerated = 2
)

func (e *ConfidentialityProtectionIndication) Encode(w *aper.AperWriter) error {
	// Encode logic for enum ConfidentialityProtectionIndication to be generated here.
	return fmt.Errorf("Encode not implemented for enum ConfidentialityProtectionIndication")
}

func (e *ConfidentialityProtectionIndication) Decode(r *aper.AperReader) error {
	// Decode logic for enum ConfidentialityProtectionIndication to be generated here.
	return fmt.Errorf("Decode not implemented for enum ConfidentialityProtectionIndication")
}
