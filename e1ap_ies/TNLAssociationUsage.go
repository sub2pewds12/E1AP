package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// TNLAssociationUsage is a generated ENUMERATED type.
type TNLAssociationUsage struct {
	Value aper.Enumerated
}

const (
	TNLAssociationUsageUe    aper.Enumerated = 0
	TNLAssociationUsageNonUe aper.Enumerated = 1
	TNLAssociationUsageBoth  aper.Enumerated = 2
)

// Encode implements the aper.AperMarshaller interface.
func (e *TNLAssociationUsage) Encode(w *aper.AperWriter) error {
	return w.WriteEnumerate(uint64(e.Value), aper.Constraint{Lb: 0, Ub: 2}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (e *TNLAssociationUsage) Decode(r *aper.AperReader) error {

	val, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true)
	if err != nil {
		return err
	}
	e.Value = aper.Enumerated(val)
	return nil
}
