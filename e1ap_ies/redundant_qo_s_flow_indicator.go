package e1ap_ies

import "github.com/lvdund/ngap/aper"

// RedundantQoSFlowIndicator From: 9_4_5_Information_Element_Definitions.txt:2065
const (
	RedundantQoSFlowIndicatorTrue  aper.Enumerated = 0
	RedundantQoSFlowIndicatorFalse aper.Enumerated = 1
)

type RedundantQoSFlowIndicator struct {
	Value aper.Enumerated
}
