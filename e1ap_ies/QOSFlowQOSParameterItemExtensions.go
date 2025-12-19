package e1ap_ies

import (
	"bytes"
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// QOSFlowQOSParameterItemExtensions is a generated type-safe wrapper for extensions.
type QOSFlowQOSParameterItemExtensions struct {
	RedundantQosFlowIndicator *RedundantQoSFlowIndicator
	TSCTrafficCharacteristics *TSCTrafficCharacteristics
}

// Encode implements the aper.AperMarshaller interface.
func (s *QOSFlowQOSParameterItemExtensions) Encode(w *aper.AperWriter) error {
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
func (s *QOSFlowQOSParameterItemExtensions) Decode(r *aper.AperReader) error {
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

		case ProtocolIEIDRedundantQosFlowIndicator:
			s.RedundantQosFlowIndicator = new(RedundantQoSFlowIndicator)
			if err := s.RedundantQosFlowIndicator.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("decode extension RedundantQosFlowIndicator failed: %w", err)
			}

		case ProtocolIEIDTSCTrafficCharacteristics:
			s.TSCTrafficCharacteristics = new(TSCTrafficCharacteristics)
			if err := s.TSCTrafficCharacteristics.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("decode extension TSCTrafficCharacteristics failed: %w", err)
			}
		default:
			// Unknown extension, ignore
		}
	}
	return nil
}
