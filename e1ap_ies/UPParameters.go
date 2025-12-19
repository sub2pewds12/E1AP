package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// UPParameters is a generated LIST type.
type UPParameters struct {
	Value []UPParametersItem
}

func (s *UPParameters) Encode(w *aper.AperWriter) (err error) {

	tmp := Sequence[aper.IE]{
		c:   aper.Constraint{Lb: 1, Ub: MaxnoofUPParameters},
		ext: false,
	}

	for i := 0; i < len(s.Value); i++ {
		tmp.Value = append(tmp.Value, &s.Value[i])
	}

	if err = tmp.Encode(w); err != nil {
		err = fmt.Errorf("encode UPParameters failed: %w", err)
		return
	}
	return nil
}

func (s *UPParameters) Decode(r *aper.AperReader) (err error) {

	// 1. Create a decoder function for the item type.
	decoder := func(r *aper.AperReader) (*UPParametersItem, error) {
		item := new(UPParametersItem)
		if err := item.Decode(r); err != nil {
			return nil, err
		}
		return item, nil
	}

	// 2. Call the generic ReadSequenceOf helper.
	//    The variable type `[]AlternativeQoSParaSetItem` now matches the function's return type.
	var decodedItems []UPParametersItem // <--- FIX: Removed the '*'
	if decodedItems, err = aper.ReadSequenceOf(decoder, r, &aper.Constraint{Lb: 1, Ub: MaxnoofUPParameters}, false); err != nil {
		return fmt.Errorf("readSequenceOf for UPParameters failed: %w", err)
	}

	// 3. Assign the decoded slice of values directly.
	s.Value = decodedItems
	return nil
}
