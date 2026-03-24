package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// DRBToSetupModItemNGRANExtensions is a generated type-safe wrapper for extensions.
type DRBToSetupModItemNGRANExtensions struct {
	DRBQOS                      *QoSFlowLevelQoSParameters
	IgnoreMappingRuleIndication *IgnoreMappingRuleIndication
	DAPSRequestInfo             *DAPSRequestInfo
}

func (s *DRBToSetupModItemNGRANExtensions) Encode(w *per.Encoder) error {
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

func (s *DRBToSetupModItemNGRANExtensions) Decode(r *per.Decoder) error {
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

		case ProtocolIEIDDRBQOS:
			s.DRBQOS = new(QoSFlowLevelQoSParameters)
			if err := s.DRBQOS.Decode(per.NewDecoder(ext.ValueBytes, per.APER)); err != nil {
				return fmt.Errorf("decode extension DRBQOS failed: %w", err)
			}

		case ProtocolIEIDIgnoreMappingRuleIndication:
			s.IgnoreMappingRuleIndication = new(IgnoreMappingRuleIndication)
			if err := s.IgnoreMappingRuleIndication.Decode(per.NewDecoder(ext.ValueBytes, per.APER)); err != nil {
				return fmt.Errorf("decode extension IgnoreMappingRuleIndication failed: %w", err)
			}

		case ProtocolIEIDDAPSRequestInfo:
			s.DAPSRequestInfo = new(DAPSRequestInfo)
			if err := s.DAPSRequestInfo.Decode(per.NewDecoder(ext.ValueBytes, per.APER)); err != nil {
				return fmt.Errorf("decode extension DAPSRequestInfo failed: %w", err)
			}
		default:
			// Unknown extension, ignore
		}
	}
	return nil
}
