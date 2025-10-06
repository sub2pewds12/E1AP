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
