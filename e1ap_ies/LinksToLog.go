package e1ap_ies

import (
	"fmt"

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

func (e *LinksToLog) Encode(w *aper.AperWriter) error {
	// Encode logic for enum LinksToLog to be generated here.
	return fmt.Errorf("Encode not implemented for enum LinksToLog")
}

func (e *LinksToLog) Decode(r *aper.AperReader) error {
	// Decode logic for enum LinksToLog to be generated here.
	return fmt.Errorf("Decode not implemented for enum LinksToLog")
}
