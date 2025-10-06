package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// MaxIPrate is a generated ENUMERATED type.
type MaxIPrate struct {
	Value aper.Enumerated
}

const (
	MaxIPrateBitrate64kbs aper.Enumerated = 0
	MaxIPrateMaxUErate    aper.Enumerated = 1
)
