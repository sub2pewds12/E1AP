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
	// fmt.Printf("--- GO DEBUG: Encoding CHOICE Cause | Choice: %d, UpperBound: 4, Extensible: false\n", s.Choice-1) // UNCOMMENT FOR DEEP DEBUGGING
	if err = w.WriteChoice(uint64(s.Choice), 4, false); err != nil {
		return fmt.Errorf("Encode choice index failed for Cause: %w", err)
	}

	// 2. Encode the selected member.
	switch s.Choice {
	case CausePresentRadioNetwork:
		if err = s.RadioNetwork.Encode(w); err != nil {
			return fmt.Errorf("encode RadioNetwork failed: %w", err)
		}
	case CausePresentTransport:
		if err = s.Transport.Encode(w); err != nil {
			return fmt.Errorf("encode Transport failed: %w", err)
		}
	case CausePresentProtocol:
		if err = s.Protocol.Encode(w); err != nil {
			return fmt.Errorf("encode Protocol failed: %w", err)
		}
	case CausePresentMisc:
		if err = s.Misc.Encode(w); err != nil {
			return fmt.Errorf("encode Misc failed: %w", err)
		}
	case CausePresentChoiceExtension:
		if err = s.ChoiceExtension.Encode(w); err != nil {
			return fmt.Errorf("encode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of Cause with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *Cause) Decode(r *aper.AperReader) (err error) {

	// 1. Read the choice index (0-based) and assign it to the struct's Choice field.
	choice, err := r.ReadChoice(4, false)
	if err != nil {
		return fmt.Errorf("read choice index failed: %w", err)
	}
	s.Choice = choice // Choice is 1-based from ReadChoice

	// 2. Decode the selected member.
	switch choice {
	case 1:
		s.RadioNetwork = new(CauseRadioNetwork)
		if err = s.RadioNetwork.Decode(r); err != nil {
			return fmt.Errorf("decode RadioNetwork failed: %w", err)
		}
	case 2:
		s.Transport = new(CauseTransport)
		if err = s.Transport.Decode(r); err != nil {
			return fmt.Errorf("decode Transport failed: %w", err)
		}
	case 3:
		s.Protocol = new(CauseProtocol)
		if err = s.Protocol.Decode(r); err != nil {
			return fmt.Errorf("decode Protocol failed: %w", err)
		}
	case 4:
		s.Misc = new(CauseMisc)
		if err = s.Misc.Decode(r); err != nil {
			return fmt.Errorf("decode Misc failed: %w", err)
		}
	case 5:
		s.ChoiceExtension = new(ProtocolIESingleContainer)
		if err = s.ChoiceExtension.Decode(r); err != nil {
			return fmt.Errorf("decode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("decode choice of Cause with unknown choice index %d", choice)
	}
	return nil
}
