package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// PrivacyIndicator is a generated ENUMERATED type.
type PrivacyIndicator struct {
	Value aper.Enumerated
}

const (
	PrivacyIndicatorImmediateMDT aper.Enumerated = 0
	PrivacyIndicatorLoggedMDT    aper.Enumerated = 1
)

func (e *PrivacyIndicator) Encode(w *aper.AperWriter) error {
	// Encode logic for enum PrivacyIndicator to be generated here.
	return fmt.Errorf("Encode not implemented for enum PrivacyIndicator")
}

func (e *PrivacyIndicator) Decode(r *aper.AperReader) error {
	// Decode logic for enum PrivacyIndicator to be generated here.
	return fmt.Errorf("Decode not implemented for enum PrivacyIndicator")
}
