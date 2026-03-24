package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// PDUSessionResourceToSetupItemExtensions is a generated type-safe wrapper for extensions.
type PDUSessionResourceToSetupItemExtensions struct {
	CommonNetworkInstance          *CommonNetworkInstance
	RedundantNGULUPTNLInformation  *UPTNLInformation
	RedundantCommonNetworkInstance *CommonNetworkInstance
	RedundantPDUSessionInformation *RedundantPDUSessionInformation
}

func (s *PDUSessionResourceToSetupItemExtensions) Encode(w *per.Encoder) error {
	var extensions []*ProtocolExtensionField

	if s.CommonNetworkInstance != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDCommonNetworkInstance},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.CommonNetworkInstance,
		})
	}

	if s.RedundantNGULUPTNLInformation != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDRedundantNGULUPTNLInformation},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.RedundantNGULUPTNLInformation,
		})
	}

	if s.RedundantCommonNetworkInstance != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDRedundantCommonNetworkInstance},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.RedundantCommonNetworkInstance,
		})
	}

	if s.RedundantPDUSessionInformation != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDRedundantPDUSessionInformation},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.RedundantPDUSessionInformation,
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

func (s *PDUSessionResourceToSetupItemExtensions) Decode(r *per.Decoder) error {
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

		case ProtocolIEIDCommonNetworkInstance:
			s.CommonNetworkInstance = new(CommonNetworkInstance)
			if err := s.CommonNetworkInstance.Decode(per.NewDecoder(ext.ValueBytes, per.APER)); err != nil {
				return fmt.Errorf("decode extension CommonNetworkInstance failed: %w", err)
			}

		case ProtocolIEIDRedundantNGULUPTNLInformation:
			s.RedundantNGULUPTNLInformation = new(UPTNLInformation)
			if err := s.RedundantNGULUPTNLInformation.Decode(per.NewDecoder(ext.ValueBytes, per.APER)); err != nil {
				return fmt.Errorf("decode extension RedundantNGULUPTNLInformation failed: %w", err)
			}

		case ProtocolIEIDRedundantCommonNetworkInstance:
			s.RedundantCommonNetworkInstance = new(CommonNetworkInstance)
			if err := s.RedundantCommonNetworkInstance.Decode(per.NewDecoder(ext.ValueBytes, per.APER)); err != nil {
				return fmt.Errorf("decode extension RedundantCommonNetworkInstance failed: %w", err)
			}

		case ProtocolIEIDRedundantPDUSessionInformation:
			s.RedundantPDUSessionInformation = new(RedundantPDUSessionInformation)
			if err := s.RedundantPDUSessionInformation.Decode(per.NewDecoder(ext.ValueBytes, per.APER)); err != nil {
				return fmt.Errorf("decode extension RedundantPDUSessionInformation failed: %w", err)
			}
		default:
			// Unknown extension, ignore
		}
	}
	return nil
}
