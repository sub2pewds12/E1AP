package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// Cause is a generated CHOICE type.
type Cause struct {
	Choice          uint64 `json:"-"`
	RadioNetwork    *CauseRadioNetwork
	Transport       *CauseTransport
	Protocol        *CauseProtocol
	Misc            *CauseMisc
	ChoiceExtension *ProtocolIESingleContainer
}

const (
	CausePresentNothing uint64 = iota
	CausePresentRadioNetwork
	CausePresentTransport
	CausePresentProtocol
	CausePresentMisc
	CausePresentChoiceExtension
)

// Encode implements the aper.AperMarshaller interface.
func (s *Cause) Encode(w *aper.AperWriter) (err error) {

	// 1. Write the choice index.
	if err = w.WriteChoice(uint64(s.Choice-1), 4, false); err != nil {
		return fmt.Errorf("Encode choice index failed: %w", err)
	}

	// 2. Encode the selected member.
	switch s.Choice {
	case CausePresentRadioNetwork:
		if err = s.RadioNetwork.Encode(w); err != nil {
			return fmt.Errorf("Encode RadioNetwork failed: %w", err)
		}
	case CausePresentTransport:
		if err = s.Transport.Encode(w); err != nil {
			return fmt.Errorf("Encode Transport failed: %w", err)
		}
	case CausePresentProtocol:
		if err = s.Protocol.Encode(w); err != nil {
			return fmt.Errorf("Encode Protocol failed: %w", err)
		}
	case CausePresentMisc:
		if err = s.Misc.Encode(w); err != nil {
			return fmt.Errorf("Encode Misc failed: %w", err)
		}
	case CausePresentChoiceExtension:
		if err = s.ChoiceExtension.Encode(w); err != nil {
			return fmt.Errorf("Encode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of Cause with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *Cause) Decode(r *aper.AperReader) (err error) {

	// 1. Read the choice index.
	var choice uint64
	if choice, err = r.ReadChoice(4, false); err != nil {
		return fmt.Errorf("Read choice index failed: %w", err)
	}

	// 2. Decode the selected member.
	switch choice {
	case 0:
		s.RadioNetwork = new(CauseRadioNetwork)
		if err = s.RadioNetwork.Decode(r); err != nil {
			return fmt.Errorf("Decode RadioNetwork failed: %w", err)
		}
	case 1:
		s.Transport = new(CauseTransport)
		if err = s.Transport.Decode(r); err != nil {
			return fmt.Errorf("Decode Transport failed: %w", err)
		}
	case 2:
		s.Protocol = new(CauseProtocol)
		if err = s.Protocol.Decode(r); err != nil {
			return fmt.Errorf("Decode Protocol failed: %w", err)
		}
	case 3:
		s.Misc = new(CauseMisc)
		if err = s.Misc.Decode(r); err != nil {
			return fmt.Errorf("Decode Misc failed: %w", err)
		}
	case 4:
		s.ChoiceExtension = new(ProtocolIESingleContainer)
		if err = s.ChoiceExtension.Decode(r); err != nil {
			return fmt.Errorf("Decode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("Decode choice of Cause with unknown choice index %d", choice)
	}
	return nil
}
