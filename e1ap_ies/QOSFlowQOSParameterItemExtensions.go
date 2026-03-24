package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// QOSFlowQOSParameterItemExtensions is a generated type-safe wrapper for extensions.
type QOSFlowQOSParameterItemExtensions struct {
	RedundantQosFlowIndicator *RedundantQoSFlowIndicator
	TSCTrafficCharacteristics *TSCTrafficCharacteristics
}

func (s *QOSFlowQOSParameterItemExtensions) Encode(w *per.Encoder) error {
	var extensions []*ProtocolExtensionField

	if s.RedundantQosFlowIndicator != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDRedundantQosFlowIndicator},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.RedundantQosFlowIndicator,
		})
	}

	if s.TSCTrafficCharacteristics != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDTSCTrafficCharacteristics},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.TSCTrafficCharacteristics,
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

func (s *QOSFlowQOSParameterItemExtensions) Decode(r *per.Decoder) error {
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

		case ProtocolIEIDRedundantQosFlowIndicator:
			s.RedundantQosFlowIndicator = new(RedundantQoSFlowIndicator)
			if err := s.RedundantQosFlowIndicator.Decode(per.NewDecoder(ext.ValueBytes, per.APER)); err != nil {
				return fmt.Errorf("decode extension RedundantQosFlowIndicator failed: %w", err)
			}

		case ProtocolIEIDTSCTrafficCharacteristics:
			s.TSCTrafficCharacteristics = new(TSCTrafficCharacteristics)
			if err := s.TSCTrafficCharacteristics.Decode(per.NewDecoder(ext.ValueBytes, per.APER)); err != nil {
				return fmt.Errorf("decode extension TSCTrafficCharacteristics failed: %w", err)
			}
		default:
			// Unknown extension, ignore
		}
	}
	return nil
}
