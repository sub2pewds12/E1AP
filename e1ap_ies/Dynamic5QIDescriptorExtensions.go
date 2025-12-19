package e1ap_ies

import (
	"bytes"
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// Dynamic5QIDescriptorExtensions is a generated type-safe wrapper for extensions.
type Dynamic5QIDescriptorExtensions struct {
	ExtendedPacketDelayBudget   *ExtendedPacketDelayBudget
	CNPacketDelayBudgetDownlink *ExtendedPacketDelayBudget
	CNPacketDelayBudgetUplink   *ExtendedPacketDelayBudget
}

// Encode implements the aper.AperMarshaller interface.
func (s *Dynamic5QIDescriptorExtensions) Encode(w *aper.AperWriter) error {
	var extensions []*ProtocolExtensionField

	if s.ExtendedPacketDelayBudget != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			Id:             ProtocolIEID{Value: ProtocolIEIDExtendedPacketDelayBudget},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.ExtendedPacketDelayBudget,
		})
	}

	if s.CNPacketDelayBudgetDownlink != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			Id:             ProtocolIEID{Value: ProtocolIEIDCNPacketDelayBudgetDownlink},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.CNPacketDelayBudgetDownlink,
		})
	}

	if s.CNPacketDelayBudgetUplink != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			Id:             ProtocolIEID{Value: ProtocolIEIDCNPacketDelayBudgetUplink},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.CNPacketDelayBudgetUplink,
		})
	}

	if len(extensions) > 0 {
		itemPointers := make([]aper.AperMarshaller, len(extensions))
		for i := 0; i < len(extensions); i++ {
			itemPointers[i] = extensions[i]
		}
		if err := aper.WriteSequenceOf(itemPointers, w, &aper.Constraint{Lb: 1, Ub: MaxProtocolExtensions}, false); err != nil {
			return fmt.Errorf("Encode extension container failed: %w", err)
		}
	} else {
		if err := aper.WriteSequenceOf([]aper.AperMarshaller(nil), w, &aper.Constraint{Lb: 1, Ub: MaxProtocolExtensions}, false); err != nil {
			return fmt.Errorf("Encode empty extension container failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *Dynamic5QIDescriptorExtensions) Decode(r *aper.AperReader) error {
	var decoder func(*aper.AperReader) (**ProtocolExtensionField, error)
	decoder = func(r *aper.AperReader) (**ProtocolExtensionField, error) {
		var item ProtocolExtensionField
		if err := item.Decode(r); err != nil {
			return nil, err
		}
		ptr := &item
		return &ptr, nil
	}

	var extensions []*ProtocolExtensionField
	var err error
	if extensions, err = aper.ReadSequenceOf[*ProtocolExtensionField](decoder, r, &aper.Constraint{Lb: 1, Ub: MaxProtocolExtensions}, false); err != nil {
		return fmt.Errorf("Decode extension container failed: %w", err)
	}

	for _, ext := range extensions {
		switch ext.Id.Value {

		case ProtocolIEIDExtendedPacketDelayBudget:
			s.ExtendedPacketDelayBudget = new(ExtendedPacketDelayBudget)
			if err := s.ExtendedPacketDelayBudget.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("Decode extension ExtendedPacketDelayBudget failed: %w", err)
			}

		case ProtocolIEIDCNPacketDelayBudgetDownlink:
			s.CNPacketDelayBudgetDownlink = new(ExtendedPacketDelayBudget)
			if err := s.CNPacketDelayBudgetDownlink.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("Decode extension CNPacketDelayBudgetDownlink failed: %w", err)
			}

		case ProtocolIEIDCNPacketDelayBudgetUplink:
			s.CNPacketDelayBudgetUplink = new(ExtendedPacketDelayBudget)
			if err := s.CNPacketDelayBudgetUplink.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("Decode extension CNPacketDelayBudgetUplink failed: %w", err)
			}
		default:
			// Unknown extension, ignore
		}
	}
	return nil
}
