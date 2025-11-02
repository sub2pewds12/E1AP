package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// IntegrityProtectionResult is a generated ENUMERATED type.
type IntegrityProtectionResult struct {
	Value aper.Enumerated
}

const (
	IntegrityProtectionResultPerformed    aper.Enumerated = 0
	IntegrityProtectionResultNotPerformed aper.Enumerated = 1
)

// Encode implements the aper.AperMarshaller interface.
func (e *IntegrityProtectionResult) Encode(w *aper.AperWriter) error {
	return w.WriteEnumerate(uint64(e.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (e *IntegrityProtectionResult) Decode(r *aper.AperReader) error {

	val, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	if err != nil {
		return err
	}
	e.Value = aper.Enumerated(val)
	return nil
}
