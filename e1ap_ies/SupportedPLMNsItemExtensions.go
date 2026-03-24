package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// SupportedPLMNsItemExtensions is a generated type-safe wrapper for extensions.
type SupportedPLMNsItemExtensions struct {
	NPNSupportInfo           *NPNSupportInfo
	ExtendedSliceSupportList *ExtendedSliceSupportList
	ExtendedNRCGISupportList *ExtendedNRCGISupportList
}

func (s *SupportedPLMNsItemExtensions) Encode(w *per.Encoder) error {
	var extensions []*ProtocolExtensionField

	if s.NPNSupportInfo != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDNPNSupportInfo},
			Criticality:    Criticality{Value: CriticalityReject},
			ExtensionValue: s.NPNSupportInfo,
		})
	}

	if s.ExtendedSliceSupportList != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDExtendedSliceSupportList},
			Criticality:    Criticality{Value: CriticalityReject},
			ExtensionValue: s.ExtendedSliceSupportList,
		})
	}

	if s.ExtendedNRCGISupportList != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDExtendedNRCGISupportList},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.ExtendedNRCGISupportList,
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

func (s *SupportedPLMNsItemExtensions) Decode(r *per.Decoder) error {
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

		case ProtocolIEIDNPNSupportInfo:
			s.NPNSupportInfo = new(NPNSupportInfo)
			if err := s.NPNSupportInfo.Decode(per.NewDecoder(ext.ValueBytes, per.APER)); err != nil {
				return fmt.Errorf("decode extension NPNSupportInfo failed: %w", err)
			}

		case ProtocolIEIDExtendedSliceSupportList:
			s.ExtendedSliceSupportList = new(ExtendedSliceSupportList)
			if err := s.ExtendedSliceSupportList.Decode(per.NewDecoder(ext.ValueBytes, per.APER)); err != nil {
				return fmt.Errorf("decode extension ExtendedSliceSupportList failed: %w", err)
			}

		case ProtocolIEIDExtendedNRCGISupportList:
			s.ExtendedNRCGISupportList = new(ExtendedNRCGISupportList)
			if err := s.ExtendedNRCGISupportList.Decode(per.NewDecoder(ext.ValueBytes, per.APER)); err != nil {
				return fmt.Errorf("decode extension ExtendedNRCGISupportList failed: %w", err)
			}
		default:
			// Unknown extension, ignore
		}
	}
	return nil
}
