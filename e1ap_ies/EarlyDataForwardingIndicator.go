package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// EarlyDataForwardingIndicator is a generated ENUMERATED type.
type EarlyDataForwardingIndicator struct {
	Value aper.Enumerated
}

const (
	EarlyDataForwardingIndicatorStop aper.Enumerated = 0
)

func (e *EarlyDataForwardingIndicator) Encode(w *aper.AperWriter) error {
	// Encode logic for enum EarlyDataForwardingIndicator to be generated here.
	return fmt.Errorf("Encode not implemented for enum EarlyDataForwardingIndicator")
}

func (e *EarlyDataForwardingIndicator) Decode(r *aper.AperReader) error {
	// Decode logic for enum EarlyDataForwardingIndicator to be generated here.
	return fmt.Errorf("Decode not implemented for enum EarlyDataForwardingIndicator")
}
