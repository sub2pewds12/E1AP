package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// BearerContextStatusChange is a generated ENUMERATED type.
type BearerContextStatusChange struct {
	Value aper.Enumerated
}

const (
	BearerContextStatusChangeSuspend aper.Enumerated = 0
	BearerContextStatusChangeResume  aper.Enumerated = 1
)

func (e *BearerContextStatusChange) Encode(w *aper.AperWriter) error {
	// Encode logic for enum BearerContextStatusChange to be generated here.
	return fmt.Errorf("Encode not implemented for enum BearerContextStatusChange")
}

func (e *BearerContextStatusChange) Decode(r *aper.AperReader) error {
	// Decode logic for enum BearerContextStatusChange to be generated here.
	return fmt.Errorf("Decode not implemented for enum BearerContextStatusChange")
}
