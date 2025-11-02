package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// Presence is a generated ENUMERATED type.
type Presence struct {
	Value aper.Enumerated
}

const (
	PresenceOptional    aper.Enumerated = 0
	PresenceConditional aper.Enumerated = 1
	PresenceMandatory   aper.Enumerated = 2
)

// Encode implements the aper.AperMarshaller interface.
func (e *Presence) Encode(w *aper.AperWriter) error {
	return w.WriteEnumerate(uint64(e.Value), aper.Constraint{Lb: 0, Ub: 2}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (e *Presence) Decode(r *aper.AperReader) error {

	val, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false)
	if err != nil {
		return err
	}
	e.Value = aper.Enumerated(val)
	return nil
}
