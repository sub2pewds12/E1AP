package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
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
func (s *QOSCharacteristics) Encode(w *aper.AperWriter) (err error) {

	// 1. Write the choice index.
	if err = w.WriteChoice(uint64(s.Choice-1), 2, false); err != nil {
		return fmt.Errorf("Encode choice index failed: %w", err)
	}

	// 2. Encode the selected member.
	switch s.Choice {
	case QOSCharacteristicsPresentNonDynamic5QI:
		if err = s.NonDynamic5QI.Encode(w); err != nil {
			return fmt.Errorf("Encode NonDynamic5QI failed: %w", err)
		}
	case QOSCharacteristicsPresentDynamic5QI:
		if err = s.Dynamic5QI.Encode(w); err != nil {
			return fmt.Errorf("Encode Dynamic5QI failed: %w", err)
		}
	case QOSCharacteristicsPresentChoiceExtension:
		if err = s.ChoiceExtension.Encode(w); err != nil {
			return fmt.Errorf("Encode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of QOSCharacteristics with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *QOSCharacteristics) Decode(r *aper.AperReader) (err error) {

	// 1. Read the choice index.
	var choice uint64
	if choice, err = r.ReadChoice(2, false); err != nil {
		return fmt.Errorf("Read choice index failed: %w", err)
	}

	// 2. Decode the selected member.
	switch choice {
	case 0:
		s.NonDynamic5QI = new(NonDynamic5QIDescriptor)
		if err = s.NonDynamic5QI.Decode(r); err != nil {
			return fmt.Errorf("Decode NonDynamic5QI failed: %w", err)
		}
	case 1:
		s.Dynamic5QI = new(Dynamic5QIDescriptor)
		if err = s.Dynamic5QI.Decode(r); err != nil {
			return fmt.Errorf("Decode Dynamic5QI failed: %w", err)
		}
	case 2:
		s.ChoiceExtension = new(ProtocolIESingleContainer)
		if err = s.ChoiceExtension.Decode(r); err != nil {
			return fmt.Errorf("Decode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("Decode choice of QOSCharacteristics with unknown choice index %d", choice)
	}
	return nil
}
