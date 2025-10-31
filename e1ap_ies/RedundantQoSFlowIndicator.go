package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// RedundantQoSFlowIndicator is a generated ENUMERATED type.
type RedundantQoSFlowIndicator struct {
	Value aper.Enumerated
}

const (
	RedundantQoSFlowIndicatorTrue  aper.Enumerated = 0
	RedundantQoSFlowIndicatorFalse aper.Enumerated = 1
)

func (e *RedundantQoSFlowIndicator) Encode(w *aper.AperWriter) error {
	// Encode logic for enum RedundantQoSFlowIndicator to be generated here.
	return fmt.Errorf("Encode not implemented for enum RedundantQoSFlowIndicator")
}

func (e *RedundantQoSFlowIndicator) Decode(r *aper.AperReader) error {
	// Decode logic for enum RedundantQoSFlowIndicator to be generated here.
	return fmt.Errorf("Decode not implemented for enum RedundantQoSFlowIndicator")
}
