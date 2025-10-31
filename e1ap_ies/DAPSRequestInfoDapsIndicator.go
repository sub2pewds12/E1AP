package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DAPSRequestInfoDapsIndicator is a generated ENUMERATED type.
type DAPSRequestInfoDapsIndicator struct {
	Value aper.Enumerated
}

const (
	DAPSRequestInfoDapsIndicatorDapsHORequired aper.Enumerated = 0
)

func (e *DAPSRequestInfoDapsIndicator) Encode(w *aper.AperWriter) error {
	// Encode logic for enum DAPSRequestInfoDapsIndicator to be generated here.
	return fmt.Errorf("Encode not implemented for enum DAPSRequestInfoDapsIndicator")
}

func (e *DAPSRequestInfoDapsIndicator) Decode(r *aper.AperReader) error {
	// Decode logic for enum DAPSRequestInfoDapsIndicator to be generated here.
	return fmt.Errorf("Decode not implemented for enum DAPSRequestInfoDapsIndicator")
}
