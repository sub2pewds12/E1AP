package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// GNBCUCPTNLAToAddList is a generated LIST type.
type GNBCUCPTNLAToAddList struct {
	Value []GNBCUCPTNLAToAddItem
}

func (s *GNBCUCPTNLAToAddList) Encode(w *aper.AperWriter) (err error) {

	tmp := Sequence[aper.IE]{
		c:   aper.Constraint{Lb: 0, Ub: MaxnoofTNLAssociations},
		ext: false,
	}

	for i := 0; i < len(s.Value); i++ {
		tmp.Value = append(tmp.Value, &s.Value[i])
	}

	if err = tmp.Encode(w); err != nil {
		err = fmt.Errorf("encode GNBCUCPTNLAToAddList failed: %w", err)
		return
	}
	return nil
}

func (s *GNBCUCPTNLAToAddList) Decode(r *aper.AperReader) (err error) {

	// 1. Create a decoder function for the item type.
	decoder := func(r *aper.AperReader) (*GNBCUCPTNLAToAddItem, error) {
		item := new(GNBCUCPTNLAToAddItem)
		if err := item.Decode(r); err != nil {
			return nil, err
		}
		return item, nil
	}

	// 2. Call the generic ReadSequenceOf helper.
	//    The variable type `[]AlternativeQoSParaSetItem` now matches the function's return type.
	var decodedItems []GNBCUCPTNLAToAddItem // <--- FIX: Removed the '*'
	if decodedItems, err = aper.ReadSequenceOf(decoder, r, &aper.Constraint{Lb: 0, Ub: MaxnoofTNLAssociations}, false); err != nil {
		return fmt.Errorf("readSequenceOf for GNBCUCPTNLAToAddList failed: %w", err)
	}

	// 3. Assign the decoded slice of values directly.
	s.Value = decodedItems
	return nil
}
