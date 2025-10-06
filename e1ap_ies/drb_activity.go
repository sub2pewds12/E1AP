package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// DRBActivity is a generated ENUMERATED type.
type DRBActivity struct {
	Value aper.Enumerated
}

const (
	DRBActivityActive    aper.Enumerated = 0
	DRBActivityNotActive aper.Enumerated = 1
)
