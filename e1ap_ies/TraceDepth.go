package e1ap_ies

import (
	"fmt"

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

func (e *TraceDepth) Encode(w *aper.AperWriter) error {
	// Encode logic for enum TraceDepth to be generated here.
	return fmt.Errorf("Encode not implemented for enum TraceDepth")
}

func (e *TraceDepth) Decode(r *aper.AperReader) error {
	// Decode logic for enum TraceDepth to be generated here.
	return fmt.Errorf("Decode not implemented for enum TraceDepth")
}
