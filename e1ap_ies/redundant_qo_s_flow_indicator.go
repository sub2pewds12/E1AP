package e1ap_ies

import (
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
