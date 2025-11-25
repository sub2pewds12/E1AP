package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// EarlyForwardingCOUNTInfo is a generated CHOICE type.
type EarlyForwardingCOUNTInfo struct {
	Choice            uint64 `json:"-"`
	FirstDLCount      *FirstDLCount
	DLDiscardingCount *DLDiscarding
	ChoiceExtension   *ProtocolIESingleContainer
}

const (
	EarlyForwardingCOUNTInfoPresentNothing uint64 = iota
	EarlyForwardingCOUNTInfoPresentFirstDLCount
	EarlyForwardingCOUNTInfoPresentDLDiscardingCount
	EarlyForwardingCOUNTInfoPresentChoiceExtension
)

// Encode implements the aper.AperMarshaller interface.
func (s *EarlyForwardingCOUNTInfo) Encode(w *aper.AperWriter) (err error) {

	// 1. Write the choice index.
	// fmt.Printf("--- GO DEBUG: Encoding CHOICE EarlyForwardingCOUNTInfo | Choice: %d, UpperBound: 2, Extensible: false\n", s.Choice-1) // UNCOMMENT FOR DEEP DEBUGGING
	if err = w.WriteChoice(uint64(s.Choice), 2, false); err != nil {
		return fmt.Errorf("Encode choice index failed for EarlyForwardingCOUNTInfo: %w", err)
	}

	// 2. Encode the selected member.
	switch s.Choice {
	case EarlyForwardingCOUNTInfoPresentFirstDLCount:
		if err = s.FirstDLCount.Encode(w); err != nil {
			return fmt.Errorf("Encode FirstDLCount failed: %w", err)
		}
	case EarlyForwardingCOUNTInfoPresentDLDiscardingCount:
		if err = s.DLDiscardingCount.Encode(w); err != nil {
			return fmt.Errorf("Encode DLDiscardingCount failed: %w", err)
		}
	case EarlyForwardingCOUNTInfoPresentChoiceExtension:
		if err = s.ChoiceExtension.Encode(w); err != nil {
			return fmt.Errorf("Encode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of EarlyForwardingCOUNTInfo with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *EarlyForwardingCOUNTInfo) Decode(r *aper.AperReader) (err error) {

	// 1. Read the choice index (0-based) and assign it to the struct's Choice field.
	var choice uint64
	if choice, err = r.ReadChoice(2, false); err != nil {
		return fmt.Errorf("Read choice index failed: %w", err)
	}
	s.Choice = choice + 1 // Convert from 0-based wire format to 1-based constant format

	// 2. Decode the selected member.
	switch s.Choice {
	case 0:
		s.FirstDLCount = new(FirstDLCount)
		if err = s.FirstDLCount.Decode(r); err != nil {
			return fmt.Errorf("Decode FirstDLCount failed: %w", err)
		}
	case 1:
		s.DLDiscardingCount = new(DLDiscarding)
		if err = s.DLDiscardingCount.Decode(r); err != nil {
			return fmt.Errorf("Decode DLDiscardingCount failed: %w", err)
		}
	case 2:
		s.ChoiceExtension = new(ProtocolIESingleContainer)
		if err = s.ChoiceExtension.Decode(r); err != nil {
			return fmt.Errorf("Decode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("Decode choice of EarlyForwardingCOUNTInfo with unknown choice index %d", choice)
	}
	return nil
}
