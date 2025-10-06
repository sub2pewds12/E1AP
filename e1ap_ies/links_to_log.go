package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// LinksToLog is a generated ENUMERATED type.
type LinksToLog struct {
	Value aper.Enumerated
}

const (
	LinksToLogUplink                aper.Enumerated = 0
	LinksToLogDownlink              aper.Enumerated = 1
	LinksToLogBothUplinkAndDownlink aper.Enumerated = 2
)
