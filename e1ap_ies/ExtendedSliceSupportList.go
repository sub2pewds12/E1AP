package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// ExtendedSliceSupportList is a generated LIST type.
type ExtendedSliceSupportList struct {
	Value []SliceSupportItem
}

func (s *ExtendedSliceSupportList) Encode(w *aper.AperWriter) (err error) {

	tmp := Sequence[aper.IE]{
		c:   aper.Constraint{Lb: 1, Ub: MaxnoofExtSliceItems},
		ext: false,
	}

	for i := 0; i < len(s.Value); i++ {
		tmp.Value = append(tmp.Value, &s.Value[i])
	}

	if err = tmp.Encode(w); err != nil {
		err = fmt.Errorf("encode ExtendedSliceSupportList failed: %w", err)
		return
	}
	return nil
}

func (s *ExtendedSliceSupportList) Decode(r *aper.AperReader) (err error) {

	// 1. Create a decoder function for the item type.
	decoder := func(r *aper.AperReader) (*SliceSupportItem, error) {
		item := new(SliceSupportItem)
		if err := item.Decode(r); err != nil {
			return nil, err
		}
		return item, nil
	}

	// 2. Call the generic ReadSequenceOf helper.
	//    The variable type `[]AlternativeQoSParaSetItem` now matches the function's return type.
	var decodedItems []SliceSupportItem // <--- FIX: Removed the '*'
	if decodedItems, err = aper.ReadSequenceOf(decoder, r, &aper.Constraint{Lb: 1, Ub: MaxnoofExtSliceItems}, false); err != nil {
		return fmt.Errorf("readSequenceOf for ExtendedSliceSupportList failed: %w", err)
	}

	// 3. Assign the decoded slice of values directly.
	s.Value = decodedItems
	return nil
}
