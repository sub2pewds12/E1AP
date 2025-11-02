package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// Dynamic5QIDescriptorDelayCritical is a generated ENUMERATED type.
type Dynamic5QIDescriptorDelayCritical struct {
	Value aper.Enumerated
}

const (
	Dynamic5QIDescriptorDelayCriticalDelayCritical    aper.Enumerated = 0
	Dynamic5QIDescriptorDelayCriticalNonDelayCritical aper.Enumerated = 1
)

// Encode implements the aper.AperMarshaller interface.
func (e *Dynamic5QIDescriptorDelayCritical) Encode(w *aper.AperWriter) error {
	return w.WriteEnumerate(uint64(e.Value), aper.Constraint{Lb: 0, Ub: 1}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (e *Dynamic5QIDescriptorDelayCritical) Decode(r *aper.AperReader) error {

	val, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false)
	if err != nil {
		return err
	}
	e.Value = aper.Enumerated(val)
	return nil
}
