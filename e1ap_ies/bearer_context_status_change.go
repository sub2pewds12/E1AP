package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// BearerContextStatusChange From: 9_4_5_Information_Element_Definitions.txt:131
// ASN.1 Data Type: ENUMERATED
const (
	BearerContextStatusChangeSuspend aper.Enumerated = 0
	BearerContextStatusChangeResume  aper.Enumerated = 1
)

type BearerContextStatusChange struct {
	Value aper.Enumerated
}
