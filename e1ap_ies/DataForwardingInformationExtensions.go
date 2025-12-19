package e1ap_ies

import (
	"bytes"
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DataForwardingInformationExtensions is a generated type-safe wrapper for extensions.
type DataForwardingInformationExtensions struct {
	DataForwardingtoNGRANQoSFlowInformationList *DataForwardingtoNGRANQoSFlowInformationList
}

// Encode implements the aper.AperMarshaller interface.
func (s *DataForwardingInformationExtensions) Encode(w *aper.AperWriter) error {
	var extensions []*ProtocolExtensionField

	if s.DataForwardingtoNGRANQoSFlowInformationList != nil {
		extensions = append(extensions, &ProtocolExtensionField{
			Id:             ProtocolIEID{Value: ProtocolIEIDDataForwardingtoNGRANQoSFlowInformationList},
			Criticality:    Criticality{Value: CriticalityIgnore},
			ExtensionValue: s.DataForwardingtoNGRANQoSFlowInformationList,
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
func (s *DataForwardingInformationExtensions) Decode(r *aper.AperReader) error {
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

		case ProtocolIEIDDataForwardingtoNGRANQoSFlowInformationList:
			s.DataForwardingtoNGRANQoSFlowInformationList = new(DataForwardingtoNGRANQoSFlowInformationList)
			if err := s.DataForwardingtoNGRANQoSFlowInformationList.Decode(aper.NewReader(bytes.NewReader(ext.ValueBytes))); err != nil {
				return fmt.Errorf("decode extension DataForwardingtoNGRANQoSFlowInformationList failed: %w", err)
			}
		default:
			// Unknown extension, ignore
		}
	}
	return nil
}
