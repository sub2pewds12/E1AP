package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// RLCMode is a generated ENUMERATED type.
type RLCMode struct {
	Value aper.Enumerated
}

const (
	RLCModeRlcTm                 aper.Enumerated = 0
	RLCModeRlcAm                 aper.Enumerated = 1
	RLCModeRlcUmBidirectional    aper.Enumerated = 2
	RLCModeRlcUmUnidirectionalUl aper.Enumerated = 3
	RLCModeRlcUmUnidirectionalDl aper.Enumerated = 4
)

// Encode implements the aper.AperMarshaller interface.
func (e *RLCMode) Encode(w *aper.AperWriter) error {
	return w.WriteEnumerate(uint64(e.Value), aper.Constraint{Lb: 0, Ub: 4}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (e *RLCMode) Decode(r *aper.AperReader) error {

	val, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 4}, true)
	if err != nil {
		return err
	}
	e.Value = aper.Enumerated(val)
	return nil
}
