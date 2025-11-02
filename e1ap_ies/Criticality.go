package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// Criticality is a generated ENUMERATED type.
type Criticality struct {
	Value aper.Enumerated
}

const (
	CriticalityReject aper.Enumerated = 0
	CriticalityIgnore aper.Enumerated = 1
	CriticalityNotify aper.Enumerated = 2
)

// Encode implements the aper.AperMarshaller interface.
func (e *Criticality) Encode(w *aper.AperWriter) error {
	return w.WriteEnumerate(uint64(e.Value), aper.Constraint{Lb: 0, Ub: 2}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (e *Criticality) Decode(r *aper.AperReader) error {

	val, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false)
	if err != nil {
		return err
	}
	e.Value = aper.Enumerated(val)
	return nil
}
