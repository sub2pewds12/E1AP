package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// PDUSessionResourceActivity is a generated ENUMERATED type.
type PDUSessionResourceActivity struct {
	Value aper.Enumerated
}

const (
	PDUSessionResourceActivityActive    aper.Enumerated = 0
	PDUSessionResourceActivityNotActive aper.Enumerated = 1
)

// Encode implements the aper.AperMarshaller interface.
func (e *PDUSessionResourceActivity) Encode(w *aper.AperWriter) error {
	return w.WriteEnumerate(uint64(e.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (e *PDUSessionResourceActivity) Decode(r *aper.AperReader) error {

	val, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	if err != nil {
		return err
	}
	e.Value = aper.Enumerated(val)
	return nil
}
