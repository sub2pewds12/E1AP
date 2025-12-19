package e1ap_ies

import (
	"bytes"
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DRBToSetupItemNGRANExtensions is a generated type-safe wrapper for extensions.
type DRBToSetupItemNGRANExtensions struct {
	DRBQOS                      *QoSFlowLevelQoSParameters
	DAPSRequestInfo             *DAPSRequestInfo
	IgnoreMappingRuleIndication *IgnoreMappingRuleIndication
	QoSFlowsDRBRemapping        *QOSFlowsDRBRemapping
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBToSetupItemNGRANExtensions) Encode(w *aper.AperWriter) error {
	var extensions []*ProtocolExtensionField

	if s.DRBQOS != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			Id:             ProtocolIEID{Value: ProtocolIEIDDRBQOS},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.DRBQOS,
		})
	}

	if s.DAPSRequestInfo != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			Id:             ProtocolIEID{Value: ProtocolIEIDDAPSRequestInfo},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.DAPSRequestInfo,
		})
	}

	if s.IgnoreMappingRuleIndication != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			Id:             ProtocolIEID{Value: ProtocolIEIDIgnoreMappingRuleIndication},
			Criticality:    Criticality{Value: CriticalityReject},
			ExtensionValue: s.IgnoreMappingRuleIndication,
		})
	}

	if s.QoSFlowsDRBRemapping != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			Id:             ProtocolIEID{Value: ProtocolIEIDQoSFlowsDRBRemapping},
			Criticality:    Criticality{Value: CriticalityReject},
			ExtensionValue: s.QoSFlowsDRBRemapping,
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
func (s *DRBToSetupItemNGRANExtensions) Decode(r *aper.AperReader) error {
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

		case ProtocolIEIDDRBQOS:
			s.DRBQOS = new(QoSFlowLevelQoSParameters)
			if err := s.DRBQOS.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("Decode extension DRBQOS failed: %w", err)
			}

		case ProtocolIEIDDAPSRequestInfo:
			s.DAPSRequestInfo = new(DAPSRequestInfo)
			if err := s.DAPSRequestInfo.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("Decode extension DAPSRequestInfo failed: %w", err)
			}

		case ProtocolIEIDIgnoreMappingRuleIndication:
			s.IgnoreMappingRuleIndication = new(IgnoreMappingRuleIndication)
			if err := s.IgnoreMappingRuleIndication.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("Decode extension IgnoreMappingRuleIndication failed: %w", err)
			}

		case ProtocolIEIDQoSFlowsDRBRemapping:
			s.QoSFlowsDRBRemapping = new(QOSFlowsDRBRemapping)
			if err := s.QoSFlowsDRBRemapping.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("Decode extension QoSFlowsDRBRemapping failed: %w", err)
			}
		default:
			// Unknown extension, ignore
		}
	}
	return nil
}
