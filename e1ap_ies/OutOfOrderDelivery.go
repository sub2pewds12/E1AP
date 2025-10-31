package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// OutOfOrderDelivery is a generated ENUMERATED type.
type OutOfOrderDelivery struct {
	Value aper.Enumerated
}

const (
	OutOfOrderDeliveryTrue aper.Enumerated = 0
)

func (e *OutOfOrderDelivery) Encode(w *aper.AperWriter) error {
	// Encode logic for enum OutOfOrderDelivery to be generated here.
	return fmt.Errorf("Encode not implemented for enum OutOfOrderDelivery")
}

func (e *OutOfOrderDelivery) Decode(r *aper.AperReader) error {
	// Decode logic for enum OutOfOrderDelivery to be generated here.
	return fmt.Errorf("Decode not implemented for enum OutOfOrderDelivery")
}
