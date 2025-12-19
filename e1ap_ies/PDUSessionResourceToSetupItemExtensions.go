package e1ap_ies

import (
	"bytes"
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// PDUSessionResourceToSetupItemExtensions is a generated type-safe wrapper for extensions.
type PDUSessionResourceToSetupItemExtensions struct {
	CommonNetworkInstance          *CommonNetworkInstance
	RedundantNGULUPTNLInformation  *UPTNLInformation
	RedundantCommonNetworkInstance *CommonNetworkInstance
	RedundantPDUSessionInformation *RedundantPDUSessionInformation
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDUSessionResourceToSetupItemExtensions) Encode(w *aper.AperWriter) error {
	var extensions []*ProtocolExtensionField

	if s.CommonNetworkInstance != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			Id:             ProtocolIEID{Value: ProtocolIEIDCommonNetworkInstance},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.CommonNetworkInstance,
		})
	}

	if s.RedundantNGULUPTNLInformation != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			Id:             ProtocolIEID{Value: ProtocolIEIDRedundantNGULUPTNLInformation},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.RedundantNGULUPTNLInformation,
		})
	}

	if s.RedundantCommonNetworkInstance != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			Id:             ProtocolIEID{Value: ProtocolIEIDRedundantCommonNetworkInstance},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.RedundantCommonNetworkInstance,
		})
	}

	if s.RedundantPDUSessionInformation != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			Id:             ProtocolIEID{Value: ProtocolIEIDRedundantPDUSessionInformation},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.RedundantPDUSessionInformation,
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
func (s *PDUSessionResourceToSetupItemExtensions) Decode(r *aper.AperReader) error {
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

		case ProtocolIEIDCommonNetworkInstance:
			s.CommonNetworkInstance = new(CommonNetworkInstance)
			if err := s.CommonNetworkInstance.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("decode extension CommonNetworkInstance failed: %w", err)
			}

		case ProtocolIEIDRedundantNGULUPTNLInformation:
			s.RedundantNGULUPTNLInformation = new(UPTNLInformation)
			if err := s.RedundantNGULUPTNLInformation.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("decode extension RedundantNGULUPTNLInformation failed: %w", err)
			}

		case ProtocolIEIDRedundantCommonNetworkInstance:
			s.RedundantCommonNetworkInstance = new(CommonNetworkInstance)
			if err := s.RedundantCommonNetworkInstance.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("decode extension RedundantCommonNetworkInstance failed: %w", err)
			}

		case ProtocolIEIDRedundantPDUSessionInformation:
			s.RedundantPDUSessionInformation = new(RedundantPDUSessionInformation)
			if err := s.RedundantPDUSessionInformation.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("decode extension RedundantPDUSessionInformation failed: %w", err)
			}
		default:
			// Unknown extension, ignore
		}
	}
	return nil
}
