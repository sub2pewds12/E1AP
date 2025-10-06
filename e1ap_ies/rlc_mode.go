package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// RLCMode is a generated ENUMERATED type.
type RLCMode struct {
	Value aper.Enumerated
}

const (
	RLCModeRlcTm                 aper.Enumerated = 0
	RLCModeRlcAm                 aper.Enumerated = 1
	RLCModeRlcUmBidirectional    aper.Enumerated = 2
	RLCModeRlcUmUnidirectionalUl aper.Enumerated = 3
	RLCModeRlcUmUnidirectionalDl aper.Enumerated = 4
)
