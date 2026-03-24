package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// CellGroupInformationItemExtensions is a generated type-safe wrapper for extensions.
type CellGroupInformationItemExtensions struct {
	NumberOfTunnels *NumberOfTunnels
}

func (s *CellGroupInformationItemExtensions) Encode(w *per.Encoder) error {
	var extensions []*ProtocolExtensionField

	if s.NumberOfTunnels != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDNumberOfTunnels},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.NumberOfTunnels,
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

func (s *CellGroupInformationItemExtensions) Decode(r *per.Decoder) error {
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

		case ProtocolIEIDNumberOfTunnels:
			s.NumberOfTunnels = new(NumberOfTunnels)
			if err := s.NumberOfTunnels.Decode(per.NewDecoder(ext.ValueBytes, per.APER)); err != nil {
				return fmt.Errorf("decode extension NumberOfTunnels failed: %w", err)
			}
		default:
			// Unknown extension, ignore
		}
	}
	return nil
}
