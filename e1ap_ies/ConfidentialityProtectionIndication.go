package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// ConfidentialityProtectionIndication is a generated ENUMERATED type.
type ConfidentialityProtectionIndication struct {
	Value aper.Enumerated
}

const (
	ConfidentialityProtectionIndicationRequired  aper.Enumerated = 0
	ConfidentialityProtectionIndicationPreferred aper.Enumerated = 1
	ConfidentialityProtectionIndicationNotNeeded aper.Enumerated = 2
)

// Encode implements the aper.AperMarshaller interface.
func (e *ConfidentialityProtectionIndication) Encode(w *aper.AperWriter) error {
	return w.WriteEnumerate(uint64(e.Value), aper.Constraint{Lb: 0, Ub: 2}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (e *ConfidentialityProtectionIndication) Decode(r *aper.AperReader) error {

	val, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true)
	if err != nil {
		return err
	}
	e.Value = aper.Enumerated(val)
	return nil
}
