package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// PDUSessionType is a generated ENUMERATED type.
type PDUSessionType struct {
	Value aper.Enumerated
}

const (
	PDUSessionTypeIpv4         aper.Enumerated = 0
	PDUSessionTypeIpv6         aper.Enumerated = 1
	PDUSessionTypeIpv4v6       aper.Enumerated = 2
	PDUSessionTypeEthernet     aper.Enumerated = 3
	PDUSessionTypeUnstructured aper.Enumerated = 4
)

// Encode implements the aper.AperMarshaller interface.
func (e *PDUSessionType) Encode(w *aper.AperWriter) error {
	return w.WriteEnumerate(uint64(e.Value), aper.Constraint{Lb: 0, Ub: 4}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (e *PDUSessionType) Decode(r *aper.AperReader) error {

	val, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 4}, true)
	if err != nil {
		return err
	}
	e.Value = aper.Enumerated(val)
	return nil
}
