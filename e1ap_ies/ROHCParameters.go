package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// ROHCParameters is a generated CHOICE type.
type ROHCParameters struct {
	Choice          uint64 `json:"-"`
	ROHC            *ROHC
	UPlinkOnlyROHC  *UplinkOnlyROHC
	ChoiceExtension *ProtocolIESingleContainer
}

const (
	ROHCParametersPresentNothing uint64 = iota
	ROHCParametersPresentROHC
	ROHCParametersPresentUPlinkOnlyROHC
	ROHCParametersPresentChoiceExtension
)

// Encode implements the aper.AperMarshaller interface.
func (s *ROHCParameters) Encode(w *aper.AperWriter) (err error) {

	// 1. Write the choice index.
	if err = w.WriteChoice(uint64(s.Choice-1), 2, false); err != nil {
		return fmt.Errorf("Encode choice index failed: %w", err)
	}

	// 2. Encode the selected member.
	switch s.Choice {
	case ROHCParametersPresentROHC:
		if err = s.ROHC.Encode(w); err != nil {
			return fmt.Errorf("Encode ROHC failed: %w", err)
		}
	case ROHCParametersPresentUPlinkOnlyROHC:
		if err = s.UPlinkOnlyROHC.Encode(w); err != nil {
			return fmt.Errorf("Encode UPlinkOnlyROHC failed: %w", err)
		}
	case ROHCParametersPresentChoiceExtension:
		if err = s.ChoiceExtension.Encode(w); err != nil {
			return fmt.Errorf("Encode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of ROHCParameters with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ROHCParameters) Decode(r *aper.AperReader) (err error) {

	// 1. Read the choice index (0-based) and assign it to the struct's Choice field.
	var choice uint64
	if choice, err = r.ReadChoice(2, false); err != nil {
		return fmt.Errorf("Read choice index failed: %w", err)
	}
	s.Choice = choice + 1 // Convert from 0-based wire format to 1-based constant format

	// 2. Decode the selected member.
	switch s.Choice {
	case 0:
		s.ROHC = new(ROHC)
		if err = s.ROHC.Decode(r); err != nil {
			return fmt.Errorf("Decode ROHC failed: %w", err)
		}
	case 1:
		s.UPlinkOnlyROHC = new(UplinkOnlyROHC)
		if err = s.UPlinkOnlyROHC.Decode(r); err != nil {
			return fmt.Errorf("Decode UPlinkOnlyROHC failed: %w", err)
		}
	case 2:
		s.ChoiceExtension = new(ProtocolIESingleContainer)
		if err = s.ChoiceExtension.Decode(r); err != nil {
			return fmt.Errorf("Decode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("Decode choice of ROHCParameters with unknown choice index %d", choice)
	}
	return nil
}
