package e1ap_ies

import (
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
