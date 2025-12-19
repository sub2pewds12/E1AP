package e1ap_ies

import (
	"bytes"
	"fmt"

	"github.com/lvdund/ngap/aper"
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

// Encode implements the aper.AperMarshaller interface.
func (s *DRBToModifyItemNGRANExtensions) Encode(w *aper.AperWriter) error {
	var extensions []*ProtocolExtensionField

	if s.OldQoSFlowMapULendmarkerexpected != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			Id:             ProtocolIEID{Value: ProtocolIEIDOldQoSFlowMapULendmarkerexpected},
			Criticality:    Criticality{Value: CriticalityReject},
			ExtensionValue: s.OldQoSFlowMapULendmarkerexpected,
		})
	}

	if s.DRBQOS != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			Id:             ProtocolIEID{Value: ProtocolIEIDDRBQOS},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.DRBQOS,
		})
	}

	if s.EarlyForwardingCOUNTReq != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			Id:             ProtocolIEID{Value: ProtocolIEIDEarlyForwardingCOUNTReq},
			Criticality:    Criticality{Value: CriticalityReject},
			ExtensionValue: s.EarlyForwardingCOUNTReq,
		})
	}

	if s.EarlyForwardingCOUNTInfo != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			Id:             ProtocolIEID{Value: ProtocolIEIDEarlyForwardingCOUNTInfo},
			Criticality:    Criticality{Value: CriticalityReject},
			ExtensionValue: s.EarlyForwardingCOUNTInfo,
		})
	}

	if s.DAPSRequestInfo != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			Id:             ProtocolIEID{Value: ProtocolIEIDDAPSRequestInfo},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.DAPSRequestInfo,
		})
	}

	if s.EarlyDataForwardingIndicator != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			Id:             ProtocolIEID{Value: ProtocolIEIDEarlyDataForwardingIndicator},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.EarlyDataForwardingIndicator,
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
func (s *DRBToModifyItemNGRANExtensions) Decode(r *aper.AperReader) error {
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

		case ProtocolIEIDOldQoSFlowMapULendmarkerexpected:
			s.OldQoSFlowMapULendmarkerexpected = new(QOSFlowList)
			if err := s.OldQoSFlowMapULendmarkerexpected.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("decode extension OldQoSFlowMapULendmarkerexpected failed: %w", err)
			}

		case ProtocolIEIDDRBQOS:
			s.DRBQOS = new(QoSFlowLevelQoSParameters)
			if err := s.DRBQOS.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("decode extension DRBQOS failed: %w", err)
			}

		case ProtocolIEIDEarlyForwardingCOUNTReq:
			s.EarlyForwardingCOUNTReq = new(EarlyForwardingCOUNTReq)
			if err := s.EarlyForwardingCOUNTReq.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("decode extension EarlyForwardingCOUNTReq failed: %w", err)
			}

		case ProtocolIEIDEarlyForwardingCOUNTInfo:
			s.EarlyForwardingCOUNTInfo = new(EarlyForwardingCOUNTInfo)
			if err := s.EarlyForwardingCOUNTInfo.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("decode extension EarlyForwardingCOUNTInfo failed: %w", err)
			}

		case ProtocolIEIDDAPSRequestInfo:
			s.DAPSRequestInfo = new(DAPSRequestInfo)
			if err := s.DAPSRequestInfo.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("decode extension DAPSRequestInfo failed: %w", err)
			}

		case ProtocolIEIDEarlyDataForwardingIndicator:
			s.EarlyDataForwardingIndicator = new(EarlyDataForwardingIndicator)
			if err := s.EarlyDataForwardingIndicator.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("decode extension EarlyDataForwardingIndicator failed: %w", err)
			}
		default:
			// Unknown extension, ignore
		}
	}
	return nil
}
