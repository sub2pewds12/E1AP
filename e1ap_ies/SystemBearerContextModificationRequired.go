package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// SystemBearerContextModificationRequired is a generated CHOICE type.
type SystemBearerContextModificationRequired struct {
	Choice                        uint64 `json:"-"`
	DRBRequiredToModifyListEUTRAN *[]DRBRequiredToModifyItemEUTRAN
	DRBRequiredToRemoveListEUTRAN *[]DRBRequiredToRemoveItemEUTRAN
}

const (
	SystemBearerContextModificationRequiredPresentNothing uint64 = iota
	SystemBearerContextModificationRequiredPresentDRBRequiredToModifyListEUTRAN
	SystemBearerContextModificationRequiredPresentDRBRequiredToRemoveListEUTRAN
)

// Encode implements the aper.AperMarshaller interface.
func (s *SystemBearerContextModificationRequired) Encode(w *aper.AperWriter) (err error) {

	// 1. Write the choice index.
	if err = w.WriteChoice(uint64(s.Choice-1), 1, false); err != nil {
		return fmt.Errorf("Encode choice index failed: %w", err)
	}

	// 2. Encode the selected member.
	switch s.Choice {
	case SystemBearerContextModificationRequiredPresentDRBRequiredToModifyListEUTRAN:
		tmp_wrapper := Sequence[*DRBRequiredToModifyItemEUTRAN]{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		for _, i := range *s.DRBRequiredToModifyListEUTRAN {
			tmp_wrapper.Value = append(tmp_wrapper.Value, &i)
		}
		if err = tmp_wrapper.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBRequiredToModifyListEUTRAN failed: %w", err)
		}
	case SystemBearerContextModificationRequiredPresentDRBRequiredToRemoveListEUTRAN:
		tmp_wrapper := Sequence[*DRBRequiredToRemoveItemEUTRAN]{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		for _, i := range *s.DRBRequiredToRemoveListEUTRAN {
			tmp_wrapper.Value = append(tmp_wrapper.Value, &i)
		}
		if err = tmp_wrapper.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBRequiredToRemoveListEUTRAN failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of SystemBearerContextModificationRequired with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *SystemBearerContextModificationRequired) Decode(r *aper.AperReader) (err error) {

	// 1. Read the choice index.
	var choice uint64
	if choice, err = r.ReadChoice(1, false); err != nil {
		return fmt.Errorf("Read choice index failed: %w", err)
	}

	// 2. Decode the selected member.
	switch choice {
	case 0:
		// 1. Create a DECODER function for the list item, as required by ReadSequenceOf.
		itemDecoder := func(r *aper.AperReader) (*DRBRequiredToModifyItemEUTRAN, error) {
			item := new(DRBRequiredToModifyItemEUTRAN)
			if err := item.Decode(r); err != nil {
				return nil, err
			}
			return item, nil
		}

		// 2. Decode the list using the aper library's generic function.
		var decodedItems []DRBRequiredToModifyItemEUTRAN
		if decodedItems, err = aper.ReadSequenceOf(itemDecoder, r, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return fmt.Errorf("Decode DRBRequiredToModifyListEUTRAN failed: %w", err)
		}

		// 3. Assign the decoded slice of values.
		s.DRBRequiredToModifyListEUTRAN = &decodedItems
	case 1:
		// 1. Create a DECODER function for the list item, as required by ReadSequenceOf.
		itemDecoder := func(r *aper.AperReader) (*DRBRequiredToRemoveItemEUTRAN, error) {
			item := new(DRBRequiredToRemoveItemEUTRAN)
			if err := item.Decode(r); err != nil {
				return nil, err
			}
			return item, nil
		}

		// 2. Decode the list using the aper library's generic function.
		var decodedItems []DRBRequiredToRemoveItemEUTRAN
		if decodedItems, err = aper.ReadSequenceOf(itemDecoder, r, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return fmt.Errorf("Decode DRBRequiredToRemoveListEUTRAN failed: %w", err)
		}

		// 3. Assign the decoded slice of values.
		s.DRBRequiredToRemoveListEUTRAN = &decodedItems
	default:
		return fmt.Errorf("Decode choice of SystemBearerContextModificationRequired with unknown choice index %d", choice)
	}
	return nil
}
