package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// NGRANQOSSupportList is a generated LIST type.
type NGRANQOSSupportList struct {
	Value []NGRANQOSSupportItem
}

func (s *NGRANQOSSupportList) Encode(w *aper.AperWriter) (err error) {

	tmp := Sequence[aper.IE]{
		c:   aper.Constraint{Lb: 1, Ub: MaxnoofNGRANQOSParameters},
		ext: false,
	}

	for i := 0; i < len(s.Value); i++ {
		tmp.Value = append(tmp.Value, &s.Value[i])
	}

	if err = tmp.Encode(w); err != nil {
		err = fmt.Errorf("encode NGRANQOSSupportList failed: %w", err)
		return
	}
	return nil
}

func (s *NGRANQOSSupportList) Decode(r *aper.AperReader) (err error) {

	// 1. Create a decoder function for the item type.
	decoder := func(r *aper.AperReader) (*NGRANQOSSupportItem, error) {
		item := new(NGRANQOSSupportItem)
		if err := item.Decode(r); err != nil {
			return nil, err
		}
		return item, nil
	}

	// 2. Call the generic ReadSequenceOf helper.
	//    The variable type `[]AlternativeQoSParaSetItem` now matches the function's return type.
	var decodedItems []NGRANQOSSupportItem // <--- FIX: Removed the '*'
	if decodedItems, err = aper.ReadSequenceOf(decoder, r, &aper.Constraint{Lb: 1, Ub: MaxnoofNGRANQOSParameters}, false); err != nil {
		return fmt.Errorf("readSequenceOf for NGRANQOSSupportList failed: %w", err)
	}

	// 3. Assign the decoded slice of values directly.
	s.Value = decodedItems
	return nil
}
