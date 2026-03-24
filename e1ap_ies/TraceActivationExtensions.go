package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// TraceActivationExtensions is a generated type-safe wrapper for extensions.
type TraceActivationExtensions struct {
	MDTConfiguration         *MDTConfiguration
	TraceCollectionEntityURI *URIaddress
}

func (s *TraceActivationExtensions) Encode(w *per.Encoder) error {
	var extensions []*ProtocolExtensionField

	if s.MDTConfiguration != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDMDTConfiguration},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.MDTConfiguration,
		})
	}

	if s.TraceCollectionEntityURI != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDTraceCollectionEntityURI},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.TraceCollectionEntityURI,
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

func (s *TraceActivationExtensions) Decode(r *per.Decoder) error {
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

		case ProtocolIEIDMDTConfiguration:
			s.MDTConfiguration = new(MDTConfiguration)
			if err := s.MDTConfiguration.Decode(per.NewDecoder(ext.ValueBytes, per.APER)); err != nil {
				return fmt.Errorf("decode extension MDTConfiguration failed: %w", err)
			}

		case ProtocolIEIDTraceCollectionEntityURI:
			s.TraceCollectionEntityURI = new(URIaddress)
			if err := s.TraceCollectionEntityURI.Decode(per.NewDecoder(ext.ValueBytes, per.APER)); err != nil {
				return fmt.Errorf("decode extension TraceCollectionEntityURI failed: %w", err)
			}
		default:
			// Unknown extension, ignore
		}
	}
	return nil
}
