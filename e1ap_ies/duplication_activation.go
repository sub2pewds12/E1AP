package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// DuplicationActivation is a generated ENUMERATED type.
type DuplicationActivation struct {
	Value aper.Enumerated
}

const (
	DuplicationActivationActive   aper.Enumerated = 0
	DuplicationActivationInactive aper.Enumerated = 1
)
