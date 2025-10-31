package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// NewULTNLInformationRequired is a generated ENUMERATED type.
type NewULTNLInformationRequired struct {
	Value aper.Enumerated
}

const (
	NewULTNLInformationRequiredRequired aper.Enumerated = 0
)

func (e *NewULTNLInformationRequired) Encode(w *aper.AperWriter) error {
	// Encode logic for enum NewULTNLInformationRequired to be generated here.
	return fmt.Errorf("Encode not implemented for enum NewULTNLInformationRequired")
}

func (e *NewULTNLInformationRequired) Decode(r *aper.AperReader) error {
	// Decode logic for enum NewULTNLInformationRequired to be generated here.
	return fmt.Errorf("Decode not implemented for enum NewULTNLInformationRequired")
}
