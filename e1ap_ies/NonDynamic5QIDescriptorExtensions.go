package e1ap_ies

import (
	"bytes"
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// NonDynamic5QIDescriptorExtensions is a generated type-safe wrapper for extensions.
type NonDynamic5QIDescriptorExtensions struct {
	CNPacketDelayBudgetDownlink *ExtendedPacketDelayBudget
	CNPacketDelayBudgetUplink   *ExtendedPacketDelayBudget
}

// Encode implements the aper.AperMarshaller interface.
func (s *NonDynamic5QIDescriptorExtensions) Encode(w *aper.AperWriter) error {
	var extensions []*ProtocolExtensionField

	if s.CNPacketDelayBudgetDownlink != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDCNPacketDelayBudgetDownlink},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.CNPacketDelayBudgetDownlink,
		})
	}

	if s.CNPacketDelayBudgetUplink != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDCNPacketDelayBudgetUplink},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.CNPacketDelayBudgetUplink,
		})
	}

	if len(extensions) > 0 {
		tmp := Sequence[*ProtocolExtensionField]{
			c:   aper.Constraint{Lb: 1, Ub: MaxProtocolExtensions},
			ext: false,
		}
		for i := 0; i < len(extensions); i++ {
			tmp.Value = append(tmp.Value, extensions[i])
		}
		if err := tmp.Encode(w); err != nil {
			return fmt.Errorf("encode extension container failed: %w", err)
		}
	} else {
		tmp := Sequence[*ProtocolExtensionField]{
			c:   aper.Constraint{Lb: 1, Ub: MaxProtocolExtensions},
			ext: false,
		}
		if err := tmp.Encode(w); err != nil {
			return fmt.Errorf("encode empty extension container failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *NonDynamic5QIDescriptorExtensions) Decode(r *aper.AperReader) error {
	decoder := func(r *aper.AperReader) (**ProtocolExtensionField, error) {
		item := new(ProtocolExtensionField)
		if err := item.Decode(r); err != nil {
			return nil, err
		}
		return &item, nil
	}

	extensions, err := aper.ReadSequenceOf(decoder, r, &aper.Constraint{Lb: 1, Ub: MaxProtocolExtensions}, false)
	if err != nil {
		return fmt.Errorf("decode extension container failed: %w", err)
	}

	for _, ext := range extensions {
		switch ext.ID.Value {

		case ProtocolIEIDCNPacketDelayBudgetDownlink:
			s.CNPacketDelayBudgetDownlink = new(ExtendedPacketDelayBudget)
			if err := s.CNPacketDelayBudgetDownlink.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("decode extension CNPacketDelayBudgetDownlink failed: %w", err)
			}

		case ProtocolIEIDCNPacketDelayBudgetUplink:
			s.CNPacketDelayBudgetUplink = new(ExtendedPacketDelayBudget)
			if err := s.CNPacketDelayBudgetUplink.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("decode extension CNPacketDelayBudgetUplink failed: %w", err)
			}
		default:
			// Unknown extension, ignore
		}
	}
	return nil
}
