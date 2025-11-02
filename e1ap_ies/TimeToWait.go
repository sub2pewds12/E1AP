package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// TimeToWait is a generated ENUMERATED type.
type TimeToWait struct {
	Value aper.Enumerated
}

const (
	TimeToWaitV1s  aper.Enumerated = 0
	TimeToWaitV2s  aper.Enumerated = 1
	TimeToWaitV5s  aper.Enumerated = 2
	TimeToWaitV10s aper.Enumerated = 3
	TimeToWaitV20s aper.Enumerated = 4
	TimeToWaitV60s aper.Enumerated = 5
)

// Encode implements the aper.AperMarshaller interface.
func (e *TimeToWait) Encode(w *aper.AperWriter) error {
	return w.WriteEnumerate(uint64(e.Value), aper.Constraint{Lb: 0, Ub: 5}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (e *TimeToWait) Decode(r *aper.AperReader) error {

	val, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, true)
	if err != nil {
		return err
	}
	e.Value = aper.Enumerated(val)
	return nil
}
