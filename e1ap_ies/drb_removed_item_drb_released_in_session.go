package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// DRBRemovedItemDRBReleasedInSession is a generated ENUMERATED type.
type DRBRemovedItemDRBReleasedInSession struct {
	Value aper.Enumerated
}

const (
	DRBRemovedItemDRBReleasedInSessionReleasedInSession    aper.Enumerated = 0
	DRBRemovedItemDRBReleasedInSessionNotReleasedInSession aper.Enumerated = 1
)
