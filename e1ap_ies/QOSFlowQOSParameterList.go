package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// QOSFlowQOSParameterList is a generated LIST type.
type QOSFlowQOSParameterList struct {
	Value []QOSFlowQOSParameterItem
}

func (s *QOSFlowQOSParameterList) Encode(w *aper.AperWriter) (err error) {

	tmp := Sequence[aper.IE]{
		c:   aper.Constraint{Lb: 1, Ub: MaxnoofQoSFlows},
		ext: false,
	}

	for i := 0; i < len(s.Value); i++ {
		tmp.Value = append(tmp.Value, &s.Value[i])
	}

	if err = tmp.Encode(w); err != nil {
		err = fmt.Errorf("encode QOSFlowQOSParameterList failed: %w", err)
		return
	}
	return nil
}

func (s *QOSFlowQOSParameterList) Decode(r *aper.AperReader) (err error) {

	// 1. Create a decoder function for the item type.
	decoder := func(r *aper.AperReader) (*QOSFlowQOSParameterItem, error) {
		item := new(QOSFlowQOSParameterItem)
		if err := item.Decode(r); err != nil {
			return nil, err
		}
		return item, nil
	}

	// 2. Call the generic ReadSequenceOf helper.
	//    The variable type `[]AlternativeQoSParaSetItem` now matches the function's return type.
	var decodedItems []QOSFlowQOSParameterItem // <--- FIX: Removed the '*'
	if decodedItems, err = aper.ReadSequenceOf(decoder, r, &aper.Constraint{Lb: 1, Ub: MaxnoofQoSFlows}, false); err != nil {
		return fmt.Errorf("readSequenceOf for QOSFlowQOSParameterList failed: %w", err)
	}

	// 3. Assign the decoded slice of values directly.
	s.Value = decodedItems
	return nil
}
