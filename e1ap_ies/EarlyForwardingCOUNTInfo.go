package e1ap_ies

import (
	"fmt"

	"asn1go/per"
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
func (s *EarlyForwardingCOUNTInfo) Encode(w *per.Encoder) (err error) {

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			per.AlternativeInfo{Name: "firstDLCount", Tag: 0},
			per.AlternativeInfo{Name: "dLDiscardingCount", Tag: 0},
			per.AlternativeInfo{Name: "choice-Extension", Tag: 0},
		},
	}
	choiceEncoder := w.NewChoiceEncoder(c)

	switch s.Choice {
	case EarlyForwardingCOUNTInfoPresentFirstDLCount:
		if err = choiceEncoder.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err = s.FirstDLCount.Encode(w); err != nil {
			return fmt.Errorf("encode FirstDLCount failed: %w", err)
		}
	case EarlyForwardingCOUNTInfoPresentDLDiscardingCount:
		if err = choiceEncoder.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err = s.DLDiscardingCount.Encode(w); err != nil {
			return fmt.Errorf("encode DLDiscardingCount failed: %w", err)
		}
	case EarlyForwardingCOUNTInfoPresentChoiceExtension:
		if err = choiceEncoder.EncodeChoice(2, false, nil); err != nil {
			return err
		}
		if err = s.ChoiceExtension.Encode(w); err != nil {
			return fmt.Errorf("encode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of EarlyForwardingCOUNTInfo with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *EarlyForwardingCOUNTInfo) Decode(r *per.Decoder) (err error) {

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			per.AlternativeInfo{Name: "firstDLCount", Tag: 0},
			per.AlternativeInfo{Name: "dLDiscardingCount", Tag: 0},
			per.AlternativeInfo{Name: "choice-Extension", Tag: 0},
		},
	}
	choiceDecoder := r.NewChoiceDecoder(c)

	choiceIndex, isExtension, _, err := choiceDecoder.DecodeChoice()
	if err != nil {
		return fmt.Errorf("decode choice index failed: %w", err)
	}

	if isExtension {
		return fmt.Errorf("extension choices are not fully supported yet")
	}

	s.Choice = uint64(choiceIndex + 1) // 1-based internal Choice enum

	switch choiceIndex {
	case 0:
		s.FirstDLCount = new(FirstDLCount)
		if err = s.FirstDLCount.Decode(r); err != nil {
			return fmt.Errorf("decode FirstDLCount failed: %w", err)
		}
	case 1:
		s.DLDiscardingCount = new(DLDiscarding)
		if err = s.DLDiscardingCount.Decode(r); err != nil {
			return fmt.Errorf("decode DLDiscardingCount failed: %w", err)
		}
	case 2:
		s.ChoiceExtension = new(ProtocolIESingleContainer)
		if err = s.ChoiceExtension.Decode(r); err != nil {
			return fmt.Errorf("decode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("decode choice of EarlyForwardingCOUNTInfo with unknown choice index %d", choiceIndex)
	}
	return nil
}
