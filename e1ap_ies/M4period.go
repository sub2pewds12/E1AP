package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// M4period is a generated ENUMERATED type.
type M4period struct {
	Value aper.Enumerated
}

const (
	M4periodMs1024  aper.Enumerated = 0
	M4periodMs2048  aper.Enumerated = 1
	M4periodMs5120  aper.Enumerated = 2
	M4periodMs10240 aper.Enumerated = 3
	M4periodMin1    aper.Enumerated = 4
)

// Encode implements the aper.AperMarshaller interface.
func (e *M4period) Encode(w *aper.AperWriter) error {
	return w.WriteEnumerate(uint64(e.Value), aper.Constraint{Lb: 0, Ub: 4}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (e *M4period) Decode(r *aper.AperReader) error {

	val, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 4}, true)
	if err != nil {
		return err
	}
	e.Value = aper.Enumerated(val)
	return nil
}
