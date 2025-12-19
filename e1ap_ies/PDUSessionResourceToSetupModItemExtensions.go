package e1ap_ies

import (
	"bytes"
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// PDUSessionResourceToSetupModItemExtensions is a generated type-safe wrapper for extensions.
type PDUSessionResourceToSetupModItemExtensions struct {
	NetworkInstance                *NetworkInstance
	CommonNetworkInstance          *CommonNetworkInstance
	RedundantNGULUPTNLInformation  *UPTNLInformation
	RedundantCommonNetworkInstance *CommonNetworkInstance
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDUSessionResourceToSetupModItemExtensions) Encode(w *aper.AperWriter) error {
	var extensions []*ProtocolExtensionField

	if s.NetworkInstance != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			Id:             ProtocolIEID{Value: ProtocolIEIDNetworkInstance},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.NetworkInstance,
		})
	}

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
func (s *PDUSessionResourceToSetupModItemExtensions) Decode(r *aper.AperReader) error {
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

		case ProtocolIEIDNetworkInstance:
			s.NetworkInstance = new(NetworkInstance)
			if err := s.NetworkInstance.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("Decode extension NetworkInstance failed: %w", err)
			}

		case ProtocolIEIDCommonNetworkInstance:
			s.CommonNetworkInstance = new(CommonNetworkInstance)
			if err := s.CommonNetworkInstance.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("Decode extension CommonNetworkInstance failed: %w", err)
			}

		case ProtocolIEIDRedundantNGULUPTNLInformation:
			s.RedundantNGULUPTNLInformation = new(UPTNLInformation)
			if err := s.RedundantNGULUPTNLInformation.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("Decode extension RedundantNGULUPTNLInformation failed: %w", err)
			}

		case ProtocolIEIDRedundantCommonNetworkInstance:
			s.RedundantCommonNetworkInstance = new(CommonNetworkInstance)
			if err := s.RedundantCommonNetworkInstance.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("Decode extension RedundantCommonNetworkInstance failed: %w", err)
			}
		default:
			// Unknown extension, ignore
		}
	}
	return nil
}
