package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// RATType is a generated ENUMERATED type.
type RATType struct {
	Value aper.Enumerated
}

const (
	RATTypeEUTRA aper.Enumerated = 0
	RATTypeNR    aper.Enumerated = 1
)
