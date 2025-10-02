package e1ap_ies

import "github.com/lvdund/ngap/aper"

// QosMonitoringRequest From: 9_4_5_Information_Element_Definitions.txt:2035
const (
	QosMonitoringRequestUl   aper.Enumerated = 0
	QosMonitoringRequestDl   aper.Enumerated = 1
	QosMonitoringRequestBoth aper.Enumerated = 2
)

type QosMonitoringRequest struct {
	Value aper.Enumerated
}
