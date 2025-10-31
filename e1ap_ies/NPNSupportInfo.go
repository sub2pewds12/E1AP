package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// NPNSupportInfo is a generated CHOICE type.
type NPNSupportInfo struct {
	Choice          uint64 `json:"-"`
	SNPN            *NPNSupportInfoSNPN
	ChoiceExtension *ProtocolIESingleContainer
}

const (
	NPNSupportInfoPresentNothing uint64 = iota
	NPNSupportInfoPresentSNPN
	NPNSupportInfoPresentChoiceExtension
)

// Encode implements the aper.AperMarshaller interface.
func (s *NPNSupportInfo) Encode(w *aper.AperWriter) (err error) {

	// 1. Write the choice index.
	if err = w.WriteChoice(uint64(s.Choice-1), 1, false); err != nil {
		return fmt.Errorf("Encode choice index failed: %w", err)
	}

	// 2. Encode the selected member.
	switch s.Choice {
	case NPNSupportInfoPresentSNPN:
		if err = s.SNPN.Encode(w); err != nil {
			return fmt.Errorf("Encode SNPN failed: %w", err)
		}
	case NPNSupportInfoPresentChoiceExtension:
		if err = s.ChoiceExtension.Encode(w); err != nil {
			return fmt.Errorf("Encode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of NPNSupportInfo with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *NPNSupportInfo) Decode(r *aper.AperReader) (err error) {

	// 1. Read the choice index.
	var choice uint64
	if choice, err = r.ReadChoice(1, false); err != nil {
		return fmt.Errorf("Read choice index failed: %w", err)
	}

	// 2. Decode the selected member.
	switch choice {
	case 0:
		s.SNPN = new(NPNSupportInfoSNPN)
		if err = s.SNPN.Decode(r); err != nil {
			return fmt.Errorf("Decode SNPN failed: %w", err)
		}
	case 1:
		s.ChoiceExtension = new(ProtocolIESingleContainer)
		if err = s.ChoiceExtension.Decode(r); err != nil {
			return fmt.Errorf("Decode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("Decode choice of NPNSupportInfo with unknown choice index %d", choice)
	}
	return nil
}
