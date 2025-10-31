package e1ap_ies

import (
	"fmt"

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

func (e *TriggeringMessage) Encode(w *aper.AperWriter) error {
	// Encode logic for enum TriggeringMessage to be generated here.
	return fmt.Errorf("Encode not implemented for enum TriggeringMessage")
}

func (e *TriggeringMessage) Decode(r *aper.AperReader) error {
	// Decode logic for enum TriggeringMessage to be generated here.
	return fmt.Errorf("Decode not implemented for enum TriggeringMessage")
}
