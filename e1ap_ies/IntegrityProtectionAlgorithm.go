package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// IntegrityProtectionAlgorithm is a generated ENUMERATED type.
type IntegrityProtectionAlgorithm struct {
	Value aper.Enumerated
}

const (
	IntegrityProtectionAlgorithmNIA0     aper.Enumerated = 0
	IntegrityProtectionAlgorithmI128NIA1 aper.Enumerated = 1
	IntegrityProtectionAlgorithmI128NIA2 aper.Enumerated = 2
	IntegrityProtectionAlgorithmI128NIA3 aper.Enumerated = 3
)

// Encode implements the aper.AperMarshaller interface.
func (e *IntegrityProtectionAlgorithm) Encode(w *aper.AperWriter) error {
	return w.WriteEnumerate(uint64(e.Value), aper.Constraint{Lb: 0, Ub: 3}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (e *IntegrityProtectionAlgorithm) Decode(r *aper.AperReader) error {

	val, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, true)
	if err != nil {
		return err
	}
	e.Value = aper.Enumerated(val)
	return nil
}
