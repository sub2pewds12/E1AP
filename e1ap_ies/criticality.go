package e1ap_ies

import "github.com/lvdund/ngap/aper"

// Criticality From: 9_4_6_Common_Definitions.txt:34
const (
	CriticalityReject aper.Enumerated = 0
	CriticalityIgnore aper.Enumerated = 1
	CriticalityNotify aper.Enumerated = 2
)

type Criticality struct {
	Value aper.Enumerated
}
