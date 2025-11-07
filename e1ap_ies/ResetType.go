package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// ResetType is a generated CHOICE type.
type ResetType struct {
	Choice            uint64 `json:"-"`
	E1Interface       *ResetAll
	PartOfE1Interface *UEAssociatedLogicalE1ConnectionListRes
	ChoiceExtension   *ProtocolIESingleContainer
}

const (
	ResetTypePresentNothing uint64 = iota
	ResetTypePresentE1Interface
	ResetTypePresentPartOfE1Interface
	ResetTypePresentChoiceExtension
)

// Encode implements the aper.AperMarshaller interface.
func (s *ResetType) Encode(w *aper.AperWriter) (err error) {

	// 1. Write the choice index.
	if err = w.WriteChoice(uint64(s.Choice-1), 2, false); err != nil {
		return fmt.Errorf("Encode choice index failed: %w", err)
	}

	// 2. Encode the selected member.
	switch s.Choice {
	case ResetTypePresentE1Interface:
		if err = s.E1Interface.Encode(w); err != nil {
			return fmt.Errorf("Encode E1Interface failed: %w", err)
		}
	case ResetTypePresentPartOfE1Interface:
		if err = s.PartOfE1Interface.Encode(w); err != nil {
			return fmt.Errorf("Encode PartOfE1Interface failed: %w", err)
		}
	case ResetTypePresentChoiceExtension:
		if err = s.ChoiceExtension.Encode(w); err != nil {
			return fmt.Errorf("Encode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of ResetType with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ResetType) Decode(r *aper.AperReader) (err error) {

	// 1. Read the choice index (0-based) and assign it to the struct's Choice field.
	var choice uint64
	if choice, err = r.ReadChoice(2, false); err != nil {
		return fmt.Errorf("Read choice index failed: %w", err)
	}
	s.Choice = choice + 1 // Convert from 0-based wire format to 1-based constant format

	// 2. Decode the selected member.
	switch s.Choice {
	case 0:
		s.E1Interface = new(ResetAll)
		if err = s.E1Interface.Decode(r); err != nil {
			return fmt.Errorf("Decode E1Interface failed: %w", err)
		}
	case 1:
		s.PartOfE1Interface = new(UEAssociatedLogicalE1ConnectionListRes)
		if err = s.PartOfE1Interface.Decode(r); err != nil {
			return fmt.Errorf("Decode PartOfE1Interface failed: %w", err)
		}
	case 2:
		s.ChoiceExtension = new(ProtocolIESingleContainer)
		if err = s.ChoiceExtension.Decode(r); err != nil {
			return fmt.Errorf("Decode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("Decode choice of ResetType with unknown choice index %d", choice)
	}
	return nil
}
