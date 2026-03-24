package e1ap_ies

import (
	"fmt"

	"asn1go/per"
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
func (s *ResetType) Encode(w *per.Encoder) (err error) {

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			per.AlternativeInfo{Name: "e1-Interface", Tag: 0},
			per.AlternativeInfo{Name: "partOfE1-Interface", Tag: 0},
			per.AlternativeInfo{Name: "choice-extension", Tag: 0},
		},
	}
	choiceEncoder := w.NewChoiceEncoder(c)

	switch s.Choice {
	case ResetTypePresentE1Interface:
		if err = choiceEncoder.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err = s.E1Interface.Encode(w); err != nil {
			return fmt.Errorf("encode E1Interface failed: %w", err)
		}
	case ResetTypePresentPartOfE1Interface:
		if err = choiceEncoder.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err = s.PartOfE1Interface.Encode(w); err != nil {
			return fmt.Errorf("encode PartOfE1Interface failed: %w", err)
		}
	case ResetTypePresentChoiceExtension:
		if err = choiceEncoder.EncodeChoice(2, false, nil); err != nil {
			return err
		}
		if err = s.ChoiceExtension.Encode(w); err != nil {
			return fmt.Errorf("encode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of ResetType with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ResetType) Decode(r *per.Decoder) (err error) {

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			per.AlternativeInfo{Name: "e1-Interface", Tag: 0},
			per.AlternativeInfo{Name: "partOfE1-Interface", Tag: 0},
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
		s.E1Interface = new(ResetAll)
		if err = s.E1Interface.Decode(r); err != nil {
			return fmt.Errorf("decode E1Interface failed: %w", err)
		}
	case 1:
		s.PartOfE1Interface = new(UEAssociatedLogicalE1ConnectionListRes)
		if err = s.PartOfE1Interface.Decode(r); err != nil {
			return fmt.Errorf("decode PartOfE1Interface failed: %w", err)
		}
	case 2:
		s.ChoiceExtension = new(ProtocolIESingleContainer)
		if err = s.ChoiceExtension.Decode(r); err != nil {
			return fmt.Errorf("decode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("decode choice of ResetType with unknown choice index %d", choiceIndex)
	}
	return nil
}
