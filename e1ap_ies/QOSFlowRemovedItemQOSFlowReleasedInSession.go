package e1ap_ies

import (
	"fmt"

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

func (e *QOSFlowRemovedItemQOSFlowReleasedInSession) Encode(w *aper.AperWriter) error {
	// Encode logic for enum QOSFlowRemovedItemQOSFlowReleasedInSession to be generated here.
	return fmt.Errorf("Encode not implemented for enum QOSFlowRemovedItemQOSFlowReleasedInSession")
}

func (e *QOSFlowRemovedItemQOSFlowReleasedInSession) Decode(r *aper.AperReader) error {
	// Decode logic for enum QOSFlowRemovedItemQOSFlowReleasedInSession to be generated here.
	return fmt.Errorf("Decode not implemented for enum QOSFlowRemovedItemQOSFlowReleasedInSession")
}
