package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// AdditionalPDCPduplicationInformation is a generated ENUMERATED type.
type AdditionalPDCPduplicationInformation struct {
	Value aper.Enumerated
}

const (
	AdditionalPDCPduplicationInformationThree aper.Enumerated = 0
	AdditionalPDCPduplicationInformationFour  aper.Enumerated = 1
)

// Encode implements the aper.AperMarshaller interface.
func (e *AdditionalPDCPduplicationInformation) Encode(w *aper.AperWriter) error {
	return w.WriteEnumerate(uint64(e.Value), aper.Constraint{Lb: 0, Ub: 1}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (e *AdditionalPDCPduplicationInformation) Decode(r *aper.AperReader) error {

	val, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true)
	if err != nil {
		return err
	}
	e.Value = aper.Enumerated(val)
	return nil
}
