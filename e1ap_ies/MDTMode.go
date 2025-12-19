package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// MDTMode is a generated CHOICE type.
type MDTMode struct {
	Choice          uint64 `json:"-"`
	ImmediateMDT    *ImmediateMDT
	ChoiceExtension *ProtocolIESingleContainer
}

const (
	MDTModePresentNothing uint64 = iota
	MDTModePresentImmediateMDT
	MDTModePresentChoiceExtension
)

// Encode implements the aper.AperMarshaller interface.
func (s *MDTMode) Encode(w *aper.AperWriter) (err error) {

	// 1. Write the choice index.
	// fmt.Printf("--- GO DEBUG: Encoding CHOICE MDTMode | Choice: %d, UpperBound: 1, Extensible: false\n", s.Choice-1) // UNCOMMENT FOR DEEP DEBUGGING
	if err = w.WriteChoice(uint64(s.Choice), 1, false); err != nil {
		return fmt.Errorf("Encode choice index failed for MDTMode: %w", err)
	}

	// 2. Encode the selected member.
	switch s.Choice {
	case MDTModePresentImmediateMDT:
		if err = s.ImmediateMDT.Encode(w); err != nil {
			return fmt.Errorf("Encode ImmediateMDT failed: %w", err)
		}
	case MDTModePresentChoiceExtension:
		if err = s.ChoiceExtension.Encode(w); err != nil {
			return fmt.Errorf("Encode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of MDTMode with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *MDTMode) Decode(r *aper.AperReader) (err error) {

	// 1. Read the choice index (0-based) and assign it to the struct's Choice field.
	var choice uint64
	if choice, err = r.ReadChoice(1, false); err != nil {
		return fmt.Errorf("Read choice index failed: %w", err)
	}
	s.Choice = choice // Choice is 1-based from ReadChoice

	// 2. Decode the selected member.
	switch choice {
	case 1:
		s.ImmediateMDT = new(ImmediateMDT)
		if err = s.ImmediateMDT.Decode(r); err != nil {
			return fmt.Errorf("Decode ImmediateMDT failed: %w", err)
		}
	case 2:
		s.ChoiceExtension = new(ProtocolIESingleContainer)
		if err = s.ChoiceExtension.Decode(r); err != nil {
			return fmt.Errorf("Decode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("Decode choice of MDTMode with unknown choice index %d", choice)
	}
	return nil
}
