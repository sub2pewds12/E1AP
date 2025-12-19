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
			Id:             ProtocolIEID{Value: ProtocolIEIDRedundantQosFlowIndicator},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.RedundantQosFlowIndicator,
		})
	}

	if s.TSCTrafficCharacteristics != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			Id:             ProtocolIEID{Value: ProtocolIEIDTSCTrafficCharacteristics},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.TSCTrafficCharacteristics,
		})
	}

	if len(extensions) > 0 {
		itemPointers := make([]aper.AperMarshaller, len(extensions))
		for i := 0; i < len(extensions); i++ {
			itemPointers[i] = extensions[i]
		}
		if err := aper.WriteSequenceOf(itemPointers, w, &aper.Constraint{Lb: 1, Ub: MaxProtocolExtensions}, false); err != nil {
			return fmt.Errorf("Encode extension container failed: %w", err)
		}
	} else {
		if err := aper.WriteSequenceOf([]aper.AperMarshaller(nil), w, &aper.Constraint{Lb: 1, Ub: MaxProtocolExtensions}, false); err != nil {
			return fmt.Errorf("Encode empty extension container failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *QOSFlowQOSParameterItemExtensions) Decode(r *aper.AperReader) error {
	var decoder func(*aper.AperReader) (**ProtocolExtensionField, error)
	decoder = func(r *aper.AperReader) (**ProtocolExtensionField, error) {
		var item ProtocolExtensionField
		if err := item.Decode(r); err != nil {
			return nil, err
		}
		ptr := &item
		return &ptr, nil
	}

	var extensions []*ProtocolExtensionField
	var err error
	if extensions, err = aper.ReadSequenceOf[*ProtocolExtensionField](decoder, r, &aper.Constraint{Lb: 1, Ub: MaxProtocolExtensions}, false); err != nil {
		return fmt.Errorf("Decode extension container failed: %w", err)
	}

	for _, ext := range extensions {
		switch ext.Id.Value {

		case ProtocolIEIDRedundantQosFlowIndicator:
			s.RedundantQosFlowIndicator = new(RedundantQoSFlowIndicator)
			if err := s.RedundantQosFlowIndicator.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("Decode extension RedundantQosFlowIndicator failed: %w", err)
			}

		case ProtocolIEIDTSCTrafficCharacteristics:
			s.TSCTrafficCharacteristics = new(TSCTrafficCharacteristics)
			if err := s.TSCTrafficCharacteristics.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("Decode extension TSCTrafficCharacteristics failed: %w", err)
			}
		default:
			// Unknown extension, ignore
		}
	}
	return nil
}
