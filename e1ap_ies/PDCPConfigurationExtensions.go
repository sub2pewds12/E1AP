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
			ID:             ProtocolIEID{Value: ProtocolIEIDPDCPStatusReportIndication},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.PDCPStatusReportIndication,
		})
	}

	if s.AdditionalPDCPduplicationInformation != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDAdditionalPDCPduplicationInformation},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.AdditionalPDCPduplicationInformation,
		})
	}

	if s.EHCParameters != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDEHCParameters},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.EHCParameters,
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
func (s *PDCPConfigurationExtensions) Decode(r *aper.AperReader) error {
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

		case ProtocolIEIDPDCPStatusReportIndication:
			s.PDCPStatusReportIndication = new(PDCPStatusReportIndication)
			if err := s.PDCPStatusReportIndication.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("decode extension PDCPStatusReportIndication failed: %w", err)
			}

		case ProtocolIEIDAdditionalPDCPduplicationInformation:
			s.AdditionalPDCPduplicationInformation = new(AdditionalPDCPduplicationInformation)
			if err := s.AdditionalPDCPduplicationInformation.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("decode extension AdditionalPDCPduplicationInformation failed: %w", err)
			}

		case ProtocolIEIDEHCParameters:
			s.EHCParameters = new(EHCParameters)
			if err := s.EHCParameters.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("decode extension EHCParameters failed: %w", err)
			}
		default:
			// Unknown extension, ignore
		}
	}
	return nil
}
