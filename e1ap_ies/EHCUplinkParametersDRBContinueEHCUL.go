package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// EHCUplinkParametersDRBContinueEHCUL is a generated ENUMERATED type.
type EHCUplinkParametersDRBContinueEHCUL struct {
	Value aper.Enumerated
}

const (
	EHCUplinkParametersDRBContinueEHCULTrue aper.Enumerated = 0
)

// Encode implements the aper.AperMarshaller interface.
func (e *EHCUplinkParametersDRBContinueEHCUL) Encode(w *aper.AperWriter) error {
	return w.WriteEnumerate(uint64(e.Value), aper.Constraint{Lb: 0, Ub: 0}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (e *EHCUplinkParametersDRBContinueEHCUL) Decode(r *aper.AperReader) error {

	val, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true)
	if err != nil {
		return err
	}
	e.Value = aper.Enumerated(val)
	return nil
}
