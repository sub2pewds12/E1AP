package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// DRBToModifyItemNGRANExtensions is a generated type-safe wrapper for extensions.
type DRBToModifyItemNGRANExtensions struct {
	OldQoSFlowMapULendmarkerexpected *QOSFlowList
	DRBQOS                           *QoSFlowLevelQoSParameters
	EarlyForwardingCOUNTReq          *EarlyForwardingCOUNTReq
	EarlyForwardingCOUNTInfo         *EarlyForwardingCOUNTInfo
	DAPSRequestInfo                  *DAPSRequestInfo
	EarlyDataForwardingIndicator     *EarlyDataForwardingIndicator
}

func (s *DRBToModifyItemNGRANExtensions) Encode(w *per.Encoder) error {
	var extensions []*ProtocolExtensionField

	if s.OldQoSFlowMapULendmarkerexpected != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDOldQoSFlowMapULendmarkerexpected},
			Criticality:    Criticality{Value: CriticalityReject},
			ExtensionValue: s.OldQoSFlowMapULendmarkerexpected,
		})
	}

	if s.DRBQOS != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDDRBQOS},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.DRBQOS,
		})
	}

	if s.EarlyForwardingCOUNTReq != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDEarlyForwardingCOUNTReq},
			Criticality:    Criticality{Value: CriticalityReject},
			ExtensionValue: s.EarlyForwardingCOUNTReq,
		})
	}

	if s.EarlyForwardingCOUNTInfo != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDEarlyForwardingCOUNTInfo},
			Criticality:    Criticality{Value: CriticalityReject},
			ExtensionValue: s.EarlyForwardingCOUNTInfo,
		})
	}

	if s.DAPSRequestInfo != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDDAPSRequestInfo},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.DAPSRequestInfo,
		})
	}

	if s.EarlyDataForwardingIndicator != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDEarlyDataForwardingIndicator},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.EarlyDataForwardingIndicator,
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

func (s *DRBToModifyItemNGRANExtensions) Decode(r *per.Decoder) error {
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

		case ProtocolIEIDOldQoSFlowMapULendmarkerexpected:
			s.OldQoSFlowMapULendmarkerexpected = new(QOSFlowList)
			if err := s.OldQoSFlowMapULendmarkerexpected.Decode(per.NewDecoder(ext.ValueBytes, per.APER)); err != nil {
				return fmt.Errorf("decode extension OldQoSFlowMapULendmarkerexpected failed: %w", err)
			}

		case ProtocolIEIDDRBQOS:
			s.DRBQOS = new(QoSFlowLevelQoSParameters)
			if err := s.DRBQOS.Decode(per.NewDecoder(ext.ValueBytes, per.APER)); err != nil {
				return fmt.Errorf("decode extension DRBQOS failed: %w", err)
			}

		case ProtocolIEIDEarlyForwardingCOUNTReq:
			s.EarlyForwardingCOUNTReq = new(EarlyForwardingCOUNTReq)
			if err := s.EarlyForwardingCOUNTReq.Decode(per.NewDecoder(ext.ValueBytes, per.APER)); err != nil {
				return fmt.Errorf("decode extension EarlyForwardingCOUNTReq failed: %w", err)
			}

		case ProtocolIEIDEarlyForwardingCOUNTInfo:
			s.EarlyForwardingCOUNTInfo = new(EarlyForwardingCOUNTInfo)
			if err := s.EarlyForwardingCOUNTInfo.Decode(per.NewDecoder(ext.ValueBytes, per.APER)); err != nil {
				return fmt.Errorf("decode extension EarlyForwardingCOUNTInfo failed: %w", err)
			}

		case ProtocolIEIDDAPSRequestInfo:
			s.DAPSRequestInfo = new(DAPSRequestInfo)
			if err := s.DAPSRequestInfo.Decode(per.NewDecoder(ext.ValueBytes, per.APER)); err != nil {
				return fmt.Errorf("decode extension DAPSRequestInfo failed: %w", err)
			}

		case ProtocolIEIDEarlyDataForwardingIndicator:
			s.EarlyDataForwardingIndicator = new(EarlyDataForwardingIndicator)
			if err := s.EarlyDataForwardingIndicator.Decode(per.NewDecoder(ext.ValueBytes, per.APER)); err != nil {
				return fmt.Errorf("decode extension EarlyDataForwardingIndicator failed: %w", err)
			}
		default:
			// Unknown extension, ignore
		}
	}
	return nil
}
