package e1ap_ies

import (
	"bytes"
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// SupportedPLMNsItemExtensions is a generated type-safe wrapper for extensions.
type SupportedPLMNsItemExtensions struct {
	NPNSupportInfo           *NPNSupportInfo
	ExtendedSliceSupportList *ExtendedSliceSupportList
	ExtendedNRCGISupportList *ExtendedNRCGISupportList
}

// Encode implements the aper.AperMarshaller interface.
func (s *SupportedPLMNsItemExtensions) Encode(w *aper.AperWriter) error {
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
func (s *SupportedPLMNsItemExtensions) Decode(r *aper.AperReader) error {
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

		case ProtocolIEIDNPNSupportInfo:
			s.NPNSupportInfo = new(NPNSupportInfo)
			if err := s.NPNSupportInfo.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("decode extension NPNSupportInfo failed: %w", err)
			}

		case ProtocolIEIDExtendedSliceSupportList:
			s.ExtendedSliceSupportList = new(ExtendedSliceSupportList)
			if err := s.ExtendedSliceSupportList.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("decode extension ExtendedSliceSupportList failed: %w", err)
			}

		case ProtocolIEIDExtendedNRCGISupportList:
			s.ExtendedNRCGISupportList = new(ExtendedNRCGISupportList)
			if err := s.ExtendedNRCGISupportList.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("decode extension ExtendedNRCGISupportList failed: %w", err)
			}
		default:
			// Unknown extension, ignore
		}
	}
	return nil
}
