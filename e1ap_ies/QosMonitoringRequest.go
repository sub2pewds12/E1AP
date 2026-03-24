package e1ap_ies

import (
	"asn1go/per"
)

// QosMonitoringRequest is a generated ENUMERATED type.
type QosMonitoringRequest struct {
	Value int64
}

const (
	QosMonitoringRequestUl   int64 = 0
	QosMonitoringRequestDl   int64 = 1
	QosMonitoringRequestBoth int64 = 2
)

// Encode implements the MessageEncoder interface for QosMonitoringRequest.
func (e *QosMonitoringRequest) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: false, RootValues: make([]int64, 3), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for QosMonitoringRequest.
func (e *QosMonitoringRequest) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: false, RootValues: make([]int64, 3), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
