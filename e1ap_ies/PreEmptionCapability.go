package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// PreEmptionCapability is a generated ENUMERATED type.
type PreEmptionCapability struct {
	Value aper.Enumerated
}

const (
	PreEmptionCapabilityShallNotTriggerPreEmption aper.Enumerated = 0
	PreEmptionCapabilityMayTriggerPreEmption      aper.Enumerated = 1
)

// Encode implements the aper.AperMarshaller interface.
func (e *PreEmptionCapability) Encode(w *aper.AperWriter) error {
	return w.WriteEnumerate(uint64(e.Value), aper.Constraint{Lb: 0, Ub: 1}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (e *PreEmptionCapability) Decode(r *aper.AperReader) error {

	val, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, false)
	if err != nil {
		return err
	}
	e.Value = aper.Enumerated(val)
	return nil
}
