package e1ap_ies

import (
	"bytes"
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// QoSFlowLevelQoSParametersExtensions is a generated type-safe wrapper for extensions.
type QoSFlowLevelQoSParametersExtensions struct {
	QoSMonitoringRequest            *QosMonitoringRequest
	MCGOfferedGBRQoSFlowInfo        *GBRQoSFlowInformation
	QosMonitoringReportingFrequency *QosMonitoringReportingFrequency
	QoSMonitoringDisabled           *QosMonitoringDisabled
}

// Encode implements the aper.AperMarshaller interface.
func (s *QoSFlowLevelQoSParametersExtensions) Encode(w *aper.AperWriter) error {
	var extensions []*ProtocolExtensionField

	if s.QoSMonitoringRequest != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			Id:             ProtocolIEID{Value: ProtocolIEIDQoSMonitoringRequest},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.QoSMonitoringRequest,
		})
	}

	if s.MCGOfferedGBRQoSFlowInfo != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			Id:             ProtocolIEID{Value: ProtocolIEIDMCGOfferedGBRQoSFlowInfo},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.MCGOfferedGBRQoSFlowInfo,
		})
	}

	if s.QosMonitoringReportingFrequency != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			Id:             ProtocolIEID{Value: ProtocolIEIDQosMonitoringReportingFrequency},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.QosMonitoringReportingFrequency,
		})
	}

	if s.QoSMonitoringDisabled != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			Id:             ProtocolIEID{Value: ProtocolIEIDQoSMonitoringDisabled},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.QoSMonitoringDisabled,
		})
	}

	if len(extensions) > 0 {
		itemPointers := make([]aper.AperMarshaller, len(extensions))
		for i := 0; i < len(extensions); i++ {
			itemPointers[i] = extensions[i]
		}
		if err := aper.WriteSequenceOf(itemPointers, w, &aper.Constraint{Lb: 1, Ub: MaxProtocolExtensions}, false); err != nil {
			return fmt.Errorf("encode extension container failed: %w", err)
		}
	} else {
		if err := aper.WriteSequenceOf([]aper.AperMarshaller(nil), w, &aper.Constraint{Lb: 1, Ub: MaxProtocolExtensions}, false); err != nil {
			return fmt.Errorf("encode empty extension container failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *QoSFlowLevelQoSParametersExtensions) Decode(r *aper.AperReader) error {
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
		switch ext.Id.Value {

		case ProtocolIEIDQoSMonitoringRequest:
			s.QoSMonitoringRequest = new(QosMonitoringRequest)
			if err := s.QoSMonitoringRequest.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("decode extension QoSMonitoringRequest failed: %w", err)
			}

		case ProtocolIEIDMCGOfferedGBRQoSFlowInfo:
			s.MCGOfferedGBRQoSFlowInfo = new(GBRQoSFlowInformation)
			if err := s.MCGOfferedGBRQoSFlowInfo.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("decode extension MCGOfferedGBRQoSFlowInfo failed: %w", err)
			}

		case ProtocolIEIDQosMonitoringReportingFrequency:
			s.QosMonitoringReportingFrequency = new(QosMonitoringReportingFrequency)
			if err := s.QosMonitoringReportingFrequency.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("decode extension QosMonitoringReportingFrequency failed: %w", err)
			}

		case ProtocolIEIDQoSMonitoringDisabled:
			s.QoSMonitoringDisabled = new(QosMonitoringDisabled)
			if err := s.QoSMonitoringDisabled.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("decode extension QoSMonitoringDisabled failed: %w", err)
			}
		default:
			// Unknown extension, ignore
		}
	}
	return nil
}
