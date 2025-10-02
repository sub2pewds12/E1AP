package e1ap_ies

import "github.com/lvdund/ngap/aper"

// DAPSRequestInfoDapsIndicator From: 9_4_5_Information_Element_Definitions.txt:307
const (
	DAPSRequestInfoDapsIndicatorDapsHORequired aper.Enumerated = 0
)

type DAPSRequestInfoDapsIndicator struct {
	Value aper.Enumerated
}
