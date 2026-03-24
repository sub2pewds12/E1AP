package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// QoSFlowLevelQoSParametersExtensions is a generated type-safe wrapper for extensions.
type QoSFlowLevelQoSParametersExtensions struct {
	QoSMonitoringRequest            *QosMonitoringRequest
	MCGOfferedGBRQoSFlowInfo        *GBRQoSFlowInformation
	QosMonitoringReportingFrequency *QosMonitoringReportingFrequency
	QoSMonitoringDisabled           *QosMonitoringDisabled
}

func (s *QoSFlowLevelQoSParametersExtensions) Encode(w *per.Encoder) error {
	var extensions []*ProtocolExtensionField

	if s.QoSMonitoringRequest != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDQoSMonitoringRequest},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.QoSMonitoringRequest,
		})
	}

	if s.MCGOfferedGBRQoSFlowInfo != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDMCGOfferedGBRQoSFlowInfo},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.MCGOfferedGBRQoSFlowInfo,
		})
	}

	if s.QosMonitoringReportingFrequency != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDQosMonitoringReportingFrequency},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.QosMonitoringReportingFrequency,
		})
	}

	if s.QoSMonitoringDisabled != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDQoSMonitoringDisabled},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.QoSMonitoringDisabled,
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

func (s *QoSFlowLevelQoSParametersExtensions) Decode(r *per.Decoder) error {
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

		case ProtocolIEIDQoSMonitoringRequest:
			s.QoSMonitoringRequest = new(QosMonitoringRequest)
			if err := s.QoSMonitoringRequest.Decode(per.NewDecoder(ext.ValueBytes, per.APER)); err != nil {
				return fmt.Errorf("decode extension QoSMonitoringRequest failed: %w", err)
			}

		case ProtocolIEIDMCGOfferedGBRQoSFlowInfo:
			s.MCGOfferedGBRQoSFlowInfo = new(GBRQoSFlowInformation)
			if err := s.MCGOfferedGBRQoSFlowInfo.Decode(per.NewDecoder(ext.ValueBytes, per.APER)); err != nil {
				return fmt.Errorf("decode extension MCGOfferedGBRQoSFlowInfo failed: %w", err)
			}

		case ProtocolIEIDQosMonitoringReportingFrequency:
			s.QosMonitoringReportingFrequency = new(QosMonitoringReportingFrequency)
			if err := s.QosMonitoringReportingFrequency.Decode(per.NewDecoder(ext.ValueBytes, per.APER)); err != nil {
				return fmt.Errorf("decode extension QosMonitoringReportingFrequency failed: %w", err)
			}

		case ProtocolIEIDQoSMonitoringDisabled:
			s.QoSMonitoringDisabled = new(QosMonitoringDisabled)
			if err := s.QoSMonitoringDisabled.Decode(per.NewDecoder(ext.ValueBytes, per.APER)); err != nil {
				return fmt.Errorf("decode extension QoSMonitoringDisabled failed: %w", err)
			}
		default:
			// Unknown extension, ignore
		}
	}
	return nil
}
