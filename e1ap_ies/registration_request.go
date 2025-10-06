package e1ap_ies

import (
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
