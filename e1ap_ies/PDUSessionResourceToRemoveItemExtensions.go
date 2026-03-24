package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// PDUSessionResourceToRemoveItemExtensions is a generated type-safe wrapper for extensions.
type PDUSessionResourceToRemoveItemExtensions struct {
	Cause *Cause
}

func (s *PDUSessionResourceToRemoveItemExtensions) Encode(w *per.Encoder) error {
	var extensions []*ProtocolExtensionField

	if s.Cause != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDCause},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.Cause,
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

func (s *PDUSessionResourceToRemoveItemExtensions) Decode(r *per.Decoder) error {
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

		case ProtocolIEIDCause:
			s.Cause = new(Cause)
			if err := s.Cause.Decode(per.NewDecoder(ext.ValueBytes, per.APER)); err != nil {
				return fmt.Errorf("decode extension Cause failed: %w", err)
			}
		default:
			// Unknown extension, ignore
		}
	}
	return nil
}
