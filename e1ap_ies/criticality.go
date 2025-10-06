package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// Criticality is a generated ENUMERATED type.
type Criticality struct {
	Value aper.Enumerated
}

const (
	CriticalityReject aper.Enumerated = 0
	CriticalityIgnore aper.Enumerated = 1
	CriticalityNotify aper.Enumerated = 2
)
