package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// DataForwardingRequest is a generated ENUMERATED type.
type DataForwardingRequest struct {
	Value aper.Enumerated
}

const (
	DataForwardingRequestUL   aper.Enumerated = 0
	DataForwardingRequestDL   aper.Enumerated = 1
	DataForwardingRequestBoth aper.Enumerated = 2
)

// Encode implements the aper.AperMarshaller interface.
func (e *DataForwardingRequest) Encode(w *aper.AperWriter) error {
	return w.WriteEnumerate(uint64(e.Value), aper.Constraint{Lb: 0, Ub: 2}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (e *DataForwardingRequest) Decode(r *aper.AperReader) error {

	val, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true)
	if err != nil {
		return err
	}
	e.Value = aper.Enumerated(val)
	return nil
}
