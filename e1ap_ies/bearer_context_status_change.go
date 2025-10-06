package e1ap_ies

import (
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
