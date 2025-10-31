package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// QosMonitoringDisabled is a generated ENUMERATED type.
type QosMonitoringDisabled struct {
	Value aper.Enumerated
}

const (
	QosMonitoringDisabledTrue aper.Enumerated = 0
)

func (e *QosMonitoringDisabled) Encode(w *aper.AperWriter) error {
	// Encode logic for enum QosMonitoringDisabled to be generated here.
	return fmt.Errorf("Encode not implemented for enum QosMonitoringDisabled")
}

func (e *QosMonitoringDisabled) Decode(r *aper.AperReader) error {
	// Decode logic for enum QosMonitoringDisabled to be generated here.
	return fmt.Errorf("Decode not implemented for enum QosMonitoringDisabled")
}
