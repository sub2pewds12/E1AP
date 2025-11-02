package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// CNSupport is a generated ENUMERATED type.
type CNSupport struct {
	Value aper.Enumerated
}

const (
	CNSupportCEpc aper.Enumerated = 0
	CNSupportC5gc aper.Enumerated = 1
	CNSupportBoth aper.Enumerated = 2
)

// Encode implements the aper.AperMarshaller interface.
func (e *CNSupport) Encode(w *aper.AperWriter) error {
	return w.WriteEnumerate(uint64(e.Value), aper.Constraint{Lb: 0, Ub: 2}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (e *CNSupport) Decode(r *aper.AperReader) error {

	val, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true)
	if err != nil {
		return err
	}
	e.Value = aper.Enumerated(val)
	return nil
}
