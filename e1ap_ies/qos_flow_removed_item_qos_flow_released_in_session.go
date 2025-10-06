package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// QOSFlowRemovedItemQOSFlowReleasedInSession is a generated ENUMERATED type.
type QOSFlowRemovedItemQOSFlowReleasedInSession struct {
	Value aper.Enumerated
}

const (
	QOSFlowRemovedItemQOSFlowReleasedInSessionReleasedInSession    aper.Enumerated = 0
	QOSFlowRemovedItemQOSFlowReleasedInSessionNotReleasedInSession aper.Enumerated = 1
)
