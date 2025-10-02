package e1ap_ies

import "github.com/lvdund/ngap/aper"

// TraceDepth From: 9_4_5_Information_Element_Definitions.txt:2285
const (
	TraceDepthMinimum                               aper.Enumerated = 0
	TraceDepthMedium                                aper.Enumerated = 1
	TraceDepthMaximum                               aper.Enumerated = 2
	TraceDepthMinimumWithoutVendorSpecificExtension aper.Enumerated = 3
	TraceDepthMediumWithoutVendorSpecificExtension  aper.Enumerated = 4
	TraceDepthMaximumWithoutVendorSpecificExtension aper.Enumerated = 5
)

type TraceDepth struct {
	Value aper.Enumerated
}
