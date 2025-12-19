package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DataForwardingtoNGRANQoSFlowInformationList is a generated LIST type.
type DataForwardingtoNGRANQoSFlowInformationList struct {
	Value []DataForwardingtoNGRANQoSFlowInformationListItem
}

func (s *DataForwardingtoNGRANQoSFlowInformationList) Encode(w *aper.AperWriter) (err error) {

	tmp := Sequence[aper.IE]{
		c:   aper.Constraint{Lb: 1, Ub: MaxnoofQoSFlows},
		ext: false,
	}

	for i := 0; i < len(s.Value); i++ {
		tmp.Value = append(tmp.Value, &s.Value[i])
	}

	if err = tmp.Encode(w); err != nil {
		err = fmt.Errorf("encode DataForwardingtoNGRANQoSFlowInformationList failed: %w", err)
		return
	}
	return nil
}

func (s *DataForwardingtoNGRANQoSFlowInformationList) Decode(r *aper.AperReader) (err error) {

	// 1. Create a decoder function for the item type.
	decoder := func(r *aper.AperReader) (*DataForwardingtoNGRANQoSFlowInformationListItem, error) {
		item := new(DataForwardingtoNGRANQoSFlowInformationListItem)
		if err := item.Decode(r); err != nil {
			return nil, err
		}
		return item, nil
	}

	// 2. Call the generic ReadSequenceOf helper.
	//    The variable type `[]AlternativeQoSParaSetItem` now matches the function's return type.
	var decodedItems []DataForwardingtoNGRANQoSFlowInformationListItem // <--- FIX: Removed the '*'
	if decodedItems, err = aper.ReadSequenceOf(decoder, r, &aper.Constraint{Lb: 1, Ub: MaxnoofQoSFlows}, false); err != nil {
		return fmt.Errorf("readSequenceOf for DataForwardingtoNGRANQoSFlowInformationList failed: %w", err)
	}

	// 3. Assign the decoded slice of values directly.
	s.Value = decodedItems
	return nil
}
