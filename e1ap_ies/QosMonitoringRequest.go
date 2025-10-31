package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// QosMonitoringRequest is a generated ENUMERATED type.
type QosMonitoringRequest struct {
	Value aper.Enumerated
}

const (
	QosMonitoringRequestUl   aper.Enumerated = 0
	QosMonitoringRequestDl   aper.Enumerated = 1
	QosMonitoringRequestBoth aper.Enumerated = 2
)

func (e *QosMonitoringRequest) Encode(w *aper.AperWriter) error {
	// Encode logic for enum QosMonitoringRequest to be generated here.
	return fmt.Errorf("Encode not implemented for enum QosMonitoringRequest")
}

func (e *QosMonitoringRequest) Decode(r *aper.AperReader) error {
	// Decode logic for enum QosMonitoringRequest to be generated here.
	return fmt.Errorf("Decode not implemented for enum QosMonitoringRequest")
}
