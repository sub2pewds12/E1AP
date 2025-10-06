package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// TypeOfError is a generated ENUMERATED type.
type TypeOfError struct {
	Value aper.Enumerated
}

const (
	TypeOfErrorNotUnderstood aper.Enumerated = 0
	TypeOfErrorMissing       aper.Enumerated = 1
)
