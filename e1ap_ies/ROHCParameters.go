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
	// fmt.Printf("--- GO DEBUG: Encoding CHOICE ROHCParameters | Choice: %d, UpperBound: 2, Extensible: false\n", s.Choice-1) // UNCOMMENT FOR DEEP DEBUGGING
	if err = w.WriteChoice(uint64(s.Choice), 2, false); err != nil {
		return fmt.Errorf("Encode choice index failed for ROHCParameters: %w", err)
	}

	// 2. Encode the selected member.
	switch s.Choice {
	case ROHCParametersPresentROHC:
		if err = s.ROHC.Encode(w); err != nil {
			return fmt.Errorf("encode ROHC failed: %w", err)
		}
	case ROHCParametersPresentUPlinkOnlyROHC:
		if err = s.UPlinkOnlyROHC.Encode(w); err != nil {
			return fmt.Errorf("encode UPlinkOnlyROHC failed: %w", err)
		}
	case ROHCParametersPresentChoiceExtension:
		if err = s.ChoiceExtension.Encode(w); err != nil {
			return fmt.Errorf("encode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of ROHCParameters with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ROHCParameters) Decode(r *aper.AperReader) (err error) {

	// 1. Read the choice index (0-based) and assign it to the struct's Choice field.
	choice, err := r.ReadChoice(2, false)
	if err != nil {
		return fmt.Errorf("read choice index failed: %w", err)
	}
	s.Choice = choice // Choice is 1-based from ReadChoice

	// 2. Decode the selected member.
	switch choice {
	case 1:
		s.ROHC = new(ROHC)
		if err = s.ROHC.Decode(r); err != nil {
			return fmt.Errorf("decode ROHC failed: %w", err)
		}
	case 2:
		s.UPlinkOnlyROHC = new(UplinkOnlyROHC)
		if err = s.UPlinkOnlyROHC.Decode(r); err != nil {
			return fmt.Errorf("decode UPlinkOnlyROHC failed: %w", err)
		}
	case 3:
		s.ChoiceExtension = new(ProtocolIESingleContainer)
		if err = s.ChoiceExtension.Decode(r); err != nil {
			return fmt.Errorf("decode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("decode choice of ROHCParameters with unknown choice index %d", choice)
	}
	return nil
}
