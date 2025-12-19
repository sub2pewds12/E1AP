package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// CellGroupInformation is a generated LIST type.
type CellGroupInformation struct {
	Value []CellGroupInformationItem
}

func (s *CellGroupInformation) Encode(w *aper.AperWriter) (err error) {

	tmp := Sequence[aper.IE]{
		c:   aper.Constraint{Lb: 1, Ub: MaxnoofCellGroups},
		ext: false,
	}

	for i := 0; i < len(s.Value); i++ {
		tmp.Value = append(tmp.Value, &s.Value[i])
	}

	if err = tmp.Encode(w); err != nil {
		err = fmt.Errorf("encode CellGroupInformation failed: %w", err)
		return
	}
	return nil
}

func (s *CellGroupInformation) Decode(r *aper.AperReader) (err error) {

	// 1. Create a decoder function for the item type.
	decoder := func(r *aper.AperReader) (*CellGroupInformationItem, error) {
		item := new(CellGroupInformationItem)
		if err := item.Decode(r); err != nil {
			return nil, err
		}
		return item, nil
	}

	// 2. Call the generic ReadSequenceOf helper.
	//    The variable type `[]AlternativeQoSParaSetItem` now matches the function's return type.
	var decodedItems []CellGroupInformationItem // <--- FIX: Removed the '*'
	if decodedItems, err = aper.ReadSequenceOf(decoder, r, &aper.Constraint{Lb: 1, Ub: MaxnoofCellGroups}, false); err != nil {
		return fmt.Errorf("readSequenceOf for CellGroupInformation failed: %w", err)
	}

	// 3. Assign the decoded slice of values directly.
	s.Value = decodedItems
	return nil
}
