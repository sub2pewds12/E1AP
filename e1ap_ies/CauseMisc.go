package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// CauseMisc is a generated ENUMERATED type.
type CauseMisc struct {
	Value aper.Enumerated
}

const (
	CauseMiscControlProcessingOverload             aper.Enumerated = 0
	CauseMiscNotEnoughUserPlaneProcessingResources aper.Enumerated = 1
	CauseMiscHardwareFailure                       aper.Enumerated = 2
	CauseMiscOmIntervention                        aper.Enumerated = 3
	CauseMiscUnspecified                           aper.Enumerated = 4
)

// Encode implements the aper.AperMarshaller interface.
func (e *CauseMisc) Encode(w *aper.AperWriter) error {
	return w.WriteEnumerate(uint64(e.Value), aper.Constraint{Lb: 0, Ub: 4}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (e *CauseMisc) Decode(r *aper.AperReader) error {

	val, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 4}, true)
	if err != nil {
		return err
	}
	e.Value = aper.Enumerated(val)
	return nil
}
