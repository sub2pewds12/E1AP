package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// DLTXStop is a generated ENUMERATED type.
type DLTXStop struct {
	Value aper.Enumerated
}

const (
	DLTXStopStop   aper.Enumerated = 0
	DLTXStopResume aper.Enumerated = 1
)
