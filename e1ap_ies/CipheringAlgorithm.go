package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// CipheringAlgorithm is a generated ENUMERATED type.
type CipheringAlgorithm struct {
	Value aper.Enumerated
}

const (
	CipheringAlgorithmNEA0     aper.Enumerated = 0
	CipheringAlgorithmC128NEA1 aper.Enumerated = 1
	CipheringAlgorithmC128NEA2 aper.Enumerated = 2
	CipheringAlgorithmC128NEA3 aper.Enumerated = 3
)

// Encode implements the aper.AperMarshaller interface.
func (e *CipheringAlgorithm) Encode(w *aper.AperWriter) error {
	return w.WriteEnumerate(uint64(e.Value), aper.Constraint{Lb: 0, Ub: 3}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (e *CipheringAlgorithm) Decode(r *aper.AperReader) error {

	val, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, true)
	if err != nil {
		return err
	}
	e.Value = aper.Enumerated(val)
	return nil
}
