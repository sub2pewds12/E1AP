package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// TriggeringMessage is a generated ENUMERATED type.
type TriggeringMessage struct {
	Value aper.Enumerated
}

const (
	TriggeringMessageInitiatingMessage   aper.Enumerated = 0
	TriggeringMessageSuccessfulOutcome   aper.Enumerated = 1
	TriggeringMessageUnsuccessfulOutcome aper.Enumerated = 2
)

// Encode implements the aper.AperMarshaller interface.
func (e *TriggeringMessage) Encode(w *aper.AperWriter) error {
	return w.WriteEnumerate(uint64(e.Value), aper.Constraint{Lb: 0, Ub: 2}, false)
}

// Decode implements the aper.AperUnmarshaller interface.
func (e *TriggeringMessage) Decode(r *aper.AperReader) error {

	val, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, false)
	if err != nil {
		return err
	}
	e.Value = aper.Enumerated(val)
	return nil
}
