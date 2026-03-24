package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
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
func (s *MDTMode) Encode(w *per.Encoder) (err error) {

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			per.AlternativeInfo{Name: "immediateMDT", Tag: 0},
			per.AlternativeInfo{Name: "choice-extension", Tag: 0},
		},
	}
	choiceEncoder := w.NewChoiceEncoder(c)

	switch s.Choice {
	case MDTModePresentImmediateMDT:
		if err = choiceEncoder.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err = s.ImmediateMDT.Encode(w); err != nil {
			return fmt.Errorf("encode ImmediateMDT failed: %w", err)
		}
	case MDTModePresentChoiceExtension:
		if err = choiceEncoder.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err = s.ChoiceExtension.Encode(w); err != nil {
			return fmt.Errorf("encode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of MDTMode with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *MDTMode) Decode(r *per.Decoder) (err error) {

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			per.AlternativeInfo{Name: "immediateMDT", Tag: 0},
			per.AlternativeInfo{Name: "choice-extension", Tag: 0},
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
		s.ImmediateMDT = new(ImmediateMDT)
		if err = s.ImmediateMDT.Decode(r); err != nil {
			return fmt.Errorf("decode ImmediateMDT failed: %w", err)
		}
	case 1:
		s.ChoiceExtension = new(ProtocolIESingleContainer)
		if err = s.ChoiceExtension.Decode(r); err != nil {
			return fmt.Errorf("decode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("decode choice of MDTMode with unknown choice index %d", choiceIndex)
	}
	return nil
}
