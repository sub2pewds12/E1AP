package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// TimeToWait is a generated ENUMERATED type.
type TimeToWait struct {
	Value aper.Enumerated
}

const (
	TimeToWaitV1s  aper.Enumerated = 0
	TimeToWaitV2s  aper.Enumerated = 1
	TimeToWaitV5s  aper.Enumerated = 2
	TimeToWaitV10s aper.Enumerated = 3
	TimeToWaitV20s aper.Enumerated = 4
	TimeToWaitV60s aper.Enumerated = 5
)
