package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// PDUSessionResourceFailedToModifyList is a generated LIST type.
type PDUSessionResourceFailedToModifyList struct {
	Value []PDUSessionResourceFailedToModifyItem
}

func (s *PDUSessionResourceFailedToModifyList) Encode(w *aper.AperWriter) (err error) {

	tmp := Sequence[aper.IE]{
		c:   aper.Constraint{Lb: 1, Ub: MaxnoofPDUSessionResource},
		ext: false,
	}

	for i := 0; i < len(s.Value); i++ {
		tmp.Value = append(tmp.Value, &s.Value[i])
	}

	if err = tmp.Encode(w); err != nil {
		err = fmt.Errorf("encode PDUSessionResourceFailedToModifyList failed: %w", err)
		return
	}
	return nil
}

func (s *PDUSessionResourceFailedToModifyList) Decode(r *aper.AperReader) (err error) {

	// 1. Create a decoder function for the item type.
	decoder := func(r *aper.AperReader) (*PDUSessionResourceFailedToModifyItem, error) {
		item := new(PDUSessionResourceFailedToModifyItem)
		if err := item.Decode(r); err != nil {
			return nil, err
		}
		return item, nil
	}

	// 2. Call the generic ReadSequenceOf helper.
	//    The variable type `[]AlternativeQoSParaSetItem` now matches the function's return type.
	var decodedItems []PDUSessionResourceFailedToModifyItem // <--- FIX: Removed the '*'
	if decodedItems, err = aper.ReadSequenceOf(decoder, r, &aper.Constraint{Lb: 1, Ub: MaxnoofPDUSessionResource}, false); err != nil {
		return fmt.Errorf("readSequenceOf for PDUSessionResourceFailedToModifyList failed: %w", err)
	}

	// 3. Assign the decoded slice of values directly.
	s.Value = decodedItems
	return nil
}
