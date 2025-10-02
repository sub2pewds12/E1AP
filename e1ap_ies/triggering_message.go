package e1ap_ies

import "github.com/lvdund/ngap/aper"

// TriggeringMessage From: 9_4_6_Common_Definitions.txt:49
const (
	TriggeringMessageInitiatingMessage   aper.Enumerated = 0
	TriggeringMessageSuccessfulOutcome   aper.Enumerated = 1
	TriggeringMessageUnsuccessfulOutcome aper.Enumerated = 2
)

type TriggeringMessage struct {
	Value aper.Enumerated
}
