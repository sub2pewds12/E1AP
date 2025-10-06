package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// UEActivity is a generated ENUMERATED type.
type UEActivity struct {
	Value aper.Enumerated
}

const (
	UEActivityActive    aper.Enumerated = 0
	UEActivityNotActive aper.Enumerated = 1
)
