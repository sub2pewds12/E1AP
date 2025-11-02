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

// Encode implements the aper.AperMarshaller interface.
func (e *QosMonitoringRequest) Encode(w *aper.AperWriter) error {
	return w.WriteEnumerate(uint64(e.Value), aper.Constraint{Lb: 0, Ub: 2}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (e *QosMonitoringRequest) Decode(r *aper.AperReader) error {

	val, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false)
	if err != nil {
		return err
	}
	e.Value = aper.Enumerated(val)
	return nil
}
