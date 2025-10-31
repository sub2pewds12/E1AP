package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// E1APPDU is a generated CHOICE type.
type E1APPDU struct {
	Choice              uint64 `json:"-"`
	InitiatingMessage   *InitiatingMessage
	SuccessfulOutcome   *SuccessfulOutcome
	UnsuccessfulOutcome *UnsuccessfulOutcome
}

const (
	E1APPDUPresentNothing uint64 = iota
	E1APPDUPresentInitiatingMessage
	E1APPDUPresentSuccessfulOutcome
	E1APPDUPresentUnsuccessfulOutcome
)

// Encode implements the aper.AperMarshaller interface.
func (s *E1APPDU) Encode(w *aper.AperWriter) (err error) {

	// 1. Write the choice index.
	if err = w.WriteChoice(uint64(s.Choice-1), 2, true); err != nil {
		return fmt.Errorf("Encode choice index failed: %w", err)
	}

	// 2. Encode the selected member.
	switch s.Choice {
	case E1APPDUPresentInitiatingMessage:
		if err = s.InitiatingMessage.Encode(w); err != nil {
			return fmt.Errorf("Encode InitiatingMessage failed: %w", err)
		}
	case E1APPDUPresentSuccessfulOutcome:
		if err = s.SuccessfulOutcome.Encode(w); err != nil {
			return fmt.Errorf("Encode SuccessfulOutcome failed: %w", err)
		}
	case E1APPDUPresentUnsuccessfulOutcome:
		if err = s.UnsuccessfulOutcome.Encode(w); err != nil {
			return fmt.Errorf("Encode UnsuccessfulOutcome failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of E1APPDU with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *E1APPDU) Decode(r *aper.AperReader) (err error) {

	// 1. Read the choice index.
	var choice uint64
	if choice, err = r.ReadChoice(2, true); err != nil {
		return fmt.Errorf("Read choice index failed: %w", err)
	}

	// 2. Decode the selected member.
	switch choice {
	case 0:
		s.InitiatingMessage = new(InitiatingMessage)
		if err = s.InitiatingMessage.Decode(r); err != nil {
			return fmt.Errorf("Decode InitiatingMessage failed: %w", err)
		}
	case 1:
		s.SuccessfulOutcome = new(SuccessfulOutcome)
		if err = s.SuccessfulOutcome.Decode(r); err != nil {
			return fmt.Errorf("Decode SuccessfulOutcome failed: %w", err)
		}
	case 2:
		s.UnsuccessfulOutcome = new(UnsuccessfulOutcome)
		if err = s.UnsuccessfulOutcome.Decode(r); err != nil {
			return fmt.Errorf("Decode UnsuccessfulOutcome failed: %w", err)
		}
	default:
		return fmt.Errorf("Decode choice of E1APPDU with unknown choice index %d", choice)
	}
	return nil
}
