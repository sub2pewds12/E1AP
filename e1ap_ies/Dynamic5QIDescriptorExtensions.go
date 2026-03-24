package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// Dynamic5QIDescriptorExtensions is a generated type-safe wrapper for extensions.
type Dynamic5QIDescriptorExtensions struct {
	ExtendedPacketDelayBudget   *ExtendedPacketDelayBudget
	CNPacketDelayBudgetDownlink *ExtendedPacketDelayBudget
	CNPacketDelayBudgetUplink   *ExtendedPacketDelayBudget
}

func (s *Dynamic5QIDescriptorExtensions) Encode(w *per.Encoder) error {
	var extensions []*ProtocolExtensionField

	if s.ExtendedPacketDelayBudget != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDExtendedPacketDelayBudget},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.ExtendedPacketDelayBudget,
		})
	}

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
		c := per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(MaxProtocolExtensions)}
		if err := w.EncodeLengthDeterminant(int64(len(extensions)), c); err != nil {
			return fmt.Errorf("encode extension container length failed: %w", err)
		}
		for _, ext := range extensions {
			if err := ext.Encode(w); err != nil {
				return fmt.Errorf("encode extension failed: %w", err)
			}
		}
	} else {
		// empty extension container
		c := per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(MaxProtocolExtensions)}
		if err := w.EncodeLengthDeterminant(0, c); err != nil {
			return err
		}
	}
	return nil
}

func (s *Dynamic5QIDescriptorExtensions) Decode(r *per.Decoder) error {
	c := per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(MaxProtocolExtensions)}
	length, err := r.DecodeLengthDeterminant(c)
	if err != nil {
		return fmt.Errorf("decode extension container length failed: %w", err)
	}

	extensions := make([]*ProtocolExtensionField, length)
	for i := int64(0); i < length; i++ {
		ext := new(ProtocolExtensionField)
		if err := ext.Decode(r); err != nil {
			return fmt.Errorf("decode extension failed: %w", err)
		}
		extensions[i] = ext
	}

	for _, ext := range extensions {
		switch ext.ID.Value {

		case ProtocolIEIDExtendedPacketDelayBudget:
			s.ExtendedPacketDelayBudget = new(ExtendedPacketDelayBudget)
			if err := s.ExtendedPacketDelayBudget.Decode(per.NewDecoder(ext.ValueBytes, per.APER)); err != nil {
				return fmt.Errorf("decode extension ExtendedPacketDelayBudget failed: %w", err)
			}

		case ProtocolIEIDCNPacketDelayBudgetDownlink:
			s.CNPacketDelayBudgetDownlink = new(ExtendedPacketDelayBudget)
			if err := s.CNPacketDelayBudgetDownlink.Decode(per.NewDecoder(ext.ValueBytes, per.APER)); err != nil {
				return fmt.Errorf("decode extension CNPacketDelayBudgetDownlink failed: %w", err)
			}

		case ProtocolIEIDCNPacketDelayBudgetUplink:
			s.CNPacketDelayBudgetUplink = new(ExtendedPacketDelayBudget)
			if err := s.CNPacketDelayBudgetUplink.Decode(per.NewDecoder(ext.ValueBytes, per.APER)); err != nil {
				return fmt.Errorf("decode extension CNPacketDelayBudgetUplink failed: %w", err)
			}
		default:
			// Unknown extension, ignore
		}
	}
	return nil
}
