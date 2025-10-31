package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// RegistrationRequest is a generated ENUMERATED type.
type RegistrationRequest struct {
	Value aper.Enumerated
}

const (
	RegistrationRequestStart aper.Enumerated = 0
	RegistrationRequestStop  aper.Enumerated = 1
)

func (e *RegistrationRequest) Encode(w *aper.AperWriter) error {
	// Encode logic for enum RegistrationRequest to be generated here.
	return fmt.Errorf("Encode not implemented for enum RegistrationRequest")
}

func (e *RegistrationRequest) Decode(r *aper.AperReader) error {
	// Decode logic for enum RegistrationRequest to be generated here.
	return fmt.Errorf("Decode not implemented for enum RegistrationRequest")
}
