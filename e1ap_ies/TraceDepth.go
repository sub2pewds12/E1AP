package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// TraceDepth is a generated ENUMERATED type.
type TraceDepth struct {
	Value aper.Enumerated
}

const (
	TraceDepthMinimum                               aper.Enumerated = 0
	TraceDepthMedium                                aper.Enumerated = 1
	TraceDepthMaximum                               aper.Enumerated = 2
	TraceDepthMinimumWithoutVendorSpecificExtension aper.Enumerated = 3
	TraceDepthMediumWithoutVendorSpecificExtension  aper.Enumerated = 4
	TraceDepthMaximumWithoutVendorSpecificExtension aper.Enumerated = 5
)

// Encode implements the aper.AperMarshaller interface.
func (e *TraceDepth) Encode(w *aper.AperWriter) error {
	return w.WriteEnumerate(uint64(e.Value), aper.Constraint{Lb: 0, Ub: 5}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (e *TraceDepth) Decode(r *aper.AperReader) error {

	val, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 5}, true)
	if err != nil {
		return err
	}
	e.Value = aper.Enumerated(val)
	return nil
}
