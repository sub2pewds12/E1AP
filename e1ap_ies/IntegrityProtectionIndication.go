package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// IntegrityProtectionIndication is a generated ENUMERATED type.
type IntegrityProtectionIndication struct {
	Value aper.Enumerated
}

const (
	IntegrityProtectionIndicationRequired  aper.Enumerated = 0
	IntegrityProtectionIndicationPreferred aper.Enumerated = 1
	IntegrityProtectionIndicationNotNeeded aper.Enumerated = 2
)

func (e *IntegrityProtectionIndication) Encode(w *aper.AperWriter) error {
	// Encode logic for enum IntegrityProtectionIndication to be generated here.
	return fmt.Errorf("Encode not implemented for enum IntegrityProtectionIndication")
}

func (e *IntegrityProtectionIndication) Decode(r *aper.AperReader) error {
	// Decode logic for enum IntegrityProtectionIndication to be generated here.
	return fmt.Errorf("Decode not implemented for enum IntegrityProtectionIndication")
}
