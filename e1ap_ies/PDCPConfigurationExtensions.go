package e1ap_ies

import (
	"bytes"
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// PDCPConfigurationExtensions is a generated type-safe wrapper for extensions.
type PDCPConfigurationExtensions struct {
	PDCPStatusReportIndication           *PDCPStatusReportIndication
	AdditionalPDCPduplicationInformation *AdditionalPDCPduplicationInformation
	EHCParameters                        *EHCParameters
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDCPConfigurationExtensions) Encode(w *aper.AperWriter) error {
	var extensions []*ProtocolExtensionField

	if s.PDCPStatusReportIndication != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			Id:             ProtocolIEID{Value: ProtocolIEIDPDCPStatusReportIndication},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.PDCPStatusReportIndication,
		})
	}

	if s.AdditionalPDCPduplicationInformation != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			Id:             ProtocolIEID{Value: ProtocolIEIDAdditionalPDCPduplicationInformation},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.AdditionalPDCPduplicationInformation,
		})
	}

	if s.EHCParameters != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			Id:             ProtocolIEID{Value: ProtocolIEIDEHCParameters},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.EHCParameters,
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
func (s *PDCPConfigurationExtensions) Decode(r *aper.AperReader) error {
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

		case ProtocolIEIDPDCPStatusReportIndication:
			s.PDCPStatusReportIndication = new(PDCPStatusReportIndication)
			if err := s.PDCPStatusReportIndication.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("Decode extension PDCPStatusReportIndication failed: %w", err)
			}

		case ProtocolIEIDAdditionalPDCPduplicationInformation:
			s.AdditionalPDCPduplicationInformation = new(AdditionalPDCPduplicationInformation)
			if err := s.AdditionalPDCPduplicationInformation.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("Decode extension AdditionalPDCPduplicationInformation failed: %w", err)
			}

		case ProtocolIEIDEHCParameters:
			s.EHCParameters = new(EHCParameters)
			if err := s.EHCParameters.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("Decode extension EHCParameters failed: %w", err)
			}
		default:
			// Unknown extension, ignore
		}
	}
	return nil
}
