package e1ap_ies

import (
	"bytes"
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// PDUSessionResourceSetupItemExtensions is a generated type-safe wrapper for extensions.
type PDUSessionResourceSetupItemExtensions struct {
	RedundantNGDLUPTNLInformation      *UPTNLInformation
	RedundantPDUSessionInformationUsed *RedundantPDUSessionInformation
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDUSessionResourceSetupItemExtensions) Encode(w *aper.AperWriter) error {
	var extensions []*ProtocolExtensionField

	if s.RedundantNGDLUPTNLInformation != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			Id:             ProtocolIEID{Value: ProtocolIEIDRedundantNGDLUPTNLInformation},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.RedundantNGDLUPTNLInformation,
		})
	}

	if s.RedundantPDUSessionInformationUsed != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			Id:             ProtocolIEID{Value: ProtocolIEIDRedundantPDUSessionInformationUsed},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.RedundantPDUSessionInformationUsed,
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
func (s *PDUSessionResourceSetupItemExtensions) Decode(r *aper.AperReader) error {
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

		case ProtocolIEIDRedundantNGDLUPTNLInformation:
			s.RedundantNGDLUPTNLInformation = new(UPTNLInformation)
			if err := s.RedundantNGDLUPTNLInformation.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("Decode extension RedundantNGDLUPTNLInformation failed: %w", err)
			}

		case ProtocolIEIDRedundantPDUSessionInformationUsed:
			s.RedundantPDUSessionInformationUsed = new(RedundantPDUSessionInformation)
			if err := s.RedundantPDUSessionInformationUsed.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("Decode extension RedundantPDUSessionInformationUsed failed: %w", err)
			}
		default:
			// Unknown extension, ignore
		}
	}
	return nil
}
