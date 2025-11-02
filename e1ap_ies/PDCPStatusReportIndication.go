package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// PDCPStatusReportIndication is a generated ENUMERATED type.
type PDCPStatusReportIndication struct {
	Value aper.Enumerated
}

const (
	PDCPStatusReportIndicationDownlink aper.Enumerated = 0
	PDCPStatusReportIndicationUplink   aper.Enumerated = 1
	PDCPStatusReportIndicationBoth     aper.Enumerated = 2
)

// Encode implements the aper.AperMarshaller interface.
func (e *PDCPStatusReportIndication) Encode(w *aper.AperWriter) error {
	return w.WriteEnumerate(uint64(e.Value), aper.Constraint{Lb: 0, Ub: 2}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (e *PDCPStatusReportIndication) Decode(r *aper.AperReader) error {

	val, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true)
	if err != nil {
		return err
	}
	e.Value = aper.Enumerated(val)
	return nil
}
