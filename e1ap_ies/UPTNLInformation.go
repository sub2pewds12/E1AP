package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// UPTNLInformation is a generated CHOICE type.
type UPTNLInformation struct {
	Choice          uint64 `json:"-"`
	GTPTunnel       *GTPTunnel
	ChoiceExtension *ProtocolIESingleContainer
}

const (
	UPTNLInformationPresentNothing uint64 = iota
	UPTNLInformationPresentGTPTunnel
	UPTNLInformationPresentChoiceExtension
)

// Encode implements the aper.AperMarshaller interface.
func (s *UPTNLInformation) Encode(w *aper.AperWriter) (err error) {

	// 1. Write the choice index.
	// fmt.Printf("--- GO DEBUG: Encoding CHOICE UPTNLInformation | Choice: %d, UpperBound: 1, Extensible: false\n", s.Choice-1) // UNCOMMENT FOR DEEP DEBUGGING
	if err = w.WriteChoice(uint64(s.Choice), 1, false); err != nil {
		return fmt.Errorf("Encode choice index failed for UPTNLInformation: %w", err)
	}

	// 2. Encode the selected member.
	switch s.Choice {
	case UPTNLInformationPresentGTPTunnel:
		if err = s.GTPTunnel.Encode(w); err != nil {
			return fmt.Errorf("encode GTPTunnel failed: %w", err)
		}
	case UPTNLInformationPresentChoiceExtension:
		if err = s.ChoiceExtension.Encode(w); err != nil {
			return fmt.Errorf("encode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of UPTNLInformation with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *UPTNLInformation) Decode(r *aper.AperReader) (err error) {

	// 1. Read the choice index (0-based) and assign it to the struct's Choice field.
	choice, err := r.ReadChoice(1, false)
	if err != nil {
		return fmt.Errorf("read choice index failed: %w", err)
	}
	s.Choice = choice // Choice is 1-based from ReadChoice

	// 2. Decode the selected member.
	switch choice {
	case 1:
		s.GTPTunnel = new(GTPTunnel)
		if err = s.GTPTunnel.Decode(r); err != nil {
			return fmt.Errorf("decode GTPTunnel failed: %w", err)
		}
	case 2:
		s.ChoiceExtension = new(ProtocolIESingleContainer)
		if err = s.ChoiceExtension.Decode(r); err != nil {
			return fmt.Errorf("decode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("decode choice of UPTNLInformation with unknown choice index %d", choice)
	}
	return nil
}
