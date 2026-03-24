package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// QOSCharacteristics is a generated CHOICE type.
type QOSCharacteristics struct {
	Choice          uint64 `json:"-"`
	NonDynamic5QI   *NonDynamic5QIDescriptor
	Dynamic5QI      *Dynamic5QIDescriptor
	ChoiceExtension *ProtocolIESingleContainer
}

const (
	QOSCharacteristicsPresentNothing uint64 = iota
	QOSCharacteristicsPresentNonDynamic5QI
	QOSCharacteristicsPresentDynamic5QI
	QOSCharacteristicsPresentChoiceExtension
)

// Encode implements the aper.AperMarshaller interface.
func (s *QOSCharacteristics) Encode(w *per.Encoder) (err error) {

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			per.AlternativeInfo{Name: "non-Dynamic-5QI", Tag: 0},
			per.AlternativeInfo{Name: "dynamic-5QI", Tag: 0},
			per.AlternativeInfo{Name: "choice-extension", Tag: 0},
		},
	}
	choiceEncoder := w.NewChoiceEncoder(c)

	switch s.Choice {
	case QOSCharacteristicsPresentNonDynamic5QI:
		if err = choiceEncoder.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err = s.NonDynamic5QI.Encode(w); err != nil {
			return fmt.Errorf("encode NonDynamic5QI failed: %w", err)
		}
	case QOSCharacteristicsPresentDynamic5QI:
		if err = choiceEncoder.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err = s.Dynamic5QI.Encode(w); err != nil {
			return fmt.Errorf("encode Dynamic5QI failed: %w", err)
		}
	case QOSCharacteristicsPresentChoiceExtension:
		if err = choiceEncoder.EncodeChoice(2, false, nil); err != nil {
			return err
		}
		if err = s.ChoiceExtension.Encode(w); err != nil {
			return fmt.Errorf("encode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of QOSCharacteristics with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *QOSCharacteristics) Decode(r *per.Decoder) (err error) {

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			per.AlternativeInfo{Name: "non-Dynamic-5QI", Tag: 0},
			per.AlternativeInfo{Name: "dynamic-5QI", Tag: 0},
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
		s.NonDynamic5QI = new(NonDynamic5QIDescriptor)
		if err = s.NonDynamic5QI.Decode(r); err != nil {
			return fmt.Errorf("decode NonDynamic5QI failed: %w", err)
		}
	case 1:
		s.Dynamic5QI = new(Dynamic5QIDescriptor)
		if err = s.Dynamic5QI.Decode(r); err != nil {
			return fmt.Errorf("decode Dynamic5QI failed: %w", err)
		}
	case 2:
		s.ChoiceExtension = new(ProtocolIESingleContainer)
		if err = s.ChoiceExtension.Decode(r); err != nil {
			return fmt.Errorf("decode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("decode choice of QOSCharacteristics with unknown choice index %d", choiceIndex)
	}
	return nil
}
