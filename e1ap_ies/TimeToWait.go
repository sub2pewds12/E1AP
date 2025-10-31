package e1ap_ies

import (
	"fmt"

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

func (e *TimeToWait) Encode(w *aper.AperWriter) error {
	// Encode logic for enum TimeToWait to be generated here.
	return fmt.Errorf("Encode not implemented for enum TimeToWait")
}

func (e *TimeToWait) Decode(r *aper.AperReader) error {
	// Decode logic for enum TimeToWait to be generated here.
	return fmt.Errorf("Decode not implemented for enum TimeToWait")
}
