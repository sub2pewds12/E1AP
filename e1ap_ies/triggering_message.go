package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// TriggeringMessage is a generated ENUMERATED type.
type TriggeringMessage struct {
	Value aper.Enumerated
}

const (
	TriggeringMessageInitiatingMessage   aper.Enumerated = 0
	TriggeringMessageSuccessfulOutcome   aper.Enumerated = 1
	TriggeringMessageUnsuccessfulOutcome aper.Enumerated = 2
)
