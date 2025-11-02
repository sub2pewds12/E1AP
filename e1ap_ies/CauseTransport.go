package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// CauseTransport is a generated ENUMERATED type.
type CauseTransport struct {
	Value aper.Enumerated
}

const (
	CauseTransportUnspecified                  aper.Enumerated = 0
	CauseTransportTransportResourceUnavailable aper.Enumerated = 1
	CauseTransportUnknownTNLAddressForIAB      aper.Enumerated = 2
)

// Encode implements the aper.AperMarshaller interface.
func (e *CauseTransport) Encode(w *aper.AperWriter) error {
	return w.WriteEnumerate(uint64(e.Value), aper.Constraint{Lb: 0, Ub: 2}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (e *CauseTransport) Decode(r *aper.AperReader) error {

	val, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true)
	if err != nil {
		return err
	}
	e.Value = aper.Enumerated(val)
	return nil
}
