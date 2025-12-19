package e1ap_ies

import (
	"bytes"
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DRBToSetupModItemNGRANExtensions is a generated type-safe wrapper for extensions.
type DRBToSetupModItemNGRANExtensions struct {
	DRBQOS                      *QoSFlowLevelQoSParameters
	IgnoreMappingRuleIndication *IgnoreMappingRuleIndication
	DAPSRequestInfo             *DAPSRequestInfo
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBToSetupModItemNGRANExtensions) Encode(w *aper.AperWriter) error {
	var extensions []*ProtocolExtensionField

	if s.DRBQOS != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDDRBQOS},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.DRBQOS,
		})
	}

	if s.IgnoreMappingRuleIndication != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDIgnoreMappingRuleIndication},
			Criticality:    Criticality{Value: CriticalityReject},
			ExtensionValue: s.IgnoreMappingRuleIndication,
		})
	}

	if s.DAPSRequestInfo != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDDAPSRequestInfo},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.DAPSRequestInfo,
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
func (s *DRBToSetupModItemNGRANExtensions) Decode(r *aper.AperReader) error {
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

		case ProtocolIEIDDRBQOS:
			s.DRBQOS = new(QoSFlowLevelQoSParameters)
			if err := s.DRBQOS.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("decode extension DRBQOS failed: %w", err)
			}

		case ProtocolIEIDIgnoreMappingRuleIndication:
			s.IgnoreMappingRuleIndication = new(IgnoreMappingRuleIndication)
			if err := s.IgnoreMappingRuleIndication.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("decode extension IgnoreMappingRuleIndication failed: %w", err)
			}

		case ProtocolIEIDDAPSRequestInfo:
			s.DAPSRequestInfo = new(DAPSRequestInfo)
			if err := s.DAPSRequestInfo.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("decode extension DAPSRequestInfo failed: %w", err)
			}
		default:
			// Unknown extension, ignore
		}
	}
	return nil
}
