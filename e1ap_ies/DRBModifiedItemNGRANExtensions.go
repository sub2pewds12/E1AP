package e1ap_ies

import (
	"bytes"
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DRBModifiedItemNGRANExtensions is a generated type-safe wrapper for extensions.
type DRBModifiedItemNGRANExtensions struct {
	EarlyForwardingCOUNTInfo         *EarlyForwardingCOUNTInfo
	OldQoSFlowMapULendmarkerexpected *QOSFlowList
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBModifiedItemNGRANExtensions) Encode(w *aper.AperWriter) error {
	var extensions []*ProtocolExtensionField

	if s.EarlyForwardingCOUNTInfo != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDEarlyForwardingCOUNTInfo},
			Criticality:    Criticality{Value: CriticalityReject},
			ExtensionValue: s.EarlyForwardingCOUNTInfo,
		})
	}

	if s.OldQoSFlowMapULendmarkerexpected != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			ID:             ProtocolIEID{Value: ProtocolIEIDOldQoSFlowMapULendmarkerexpected},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.OldQoSFlowMapULendmarkerexpected,
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
func (s *DRBModifiedItemNGRANExtensions) Decode(r *aper.AperReader) error {
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

		case ProtocolIEIDEarlyForwardingCOUNTInfo:
			s.EarlyForwardingCOUNTInfo = new(EarlyForwardingCOUNTInfo)
			if err := s.EarlyForwardingCOUNTInfo.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("decode extension EarlyForwardingCOUNTInfo failed: %w", err)
			}

		case ProtocolIEIDOldQoSFlowMapULendmarkerexpected:
			s.OldQoSFlowMapULendmarkerexpected = new(QOSFlowList)
			if err := s.OldQoSFlowMapULendmarkerexpected.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("decode extension OldQoSFlowMapULendmarkerexpected failed: %w", err)
			}
		default:
			// Unknown extension, ignore
		}
	}
	return nil
}
