package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// SystemGNBCUUPCounterCheckRequest is a generated CHOICE type.
type SystemGNBCUUPCounterCheckRequest struct {
	Choice                              uint64 `json:"-"`
	DRBsSubjectToCounterCheckListEUTRAN *[]DRBsSubjectToCounterCheckItemEUTRAN
}

const (
	SystemGNBCUUPCounterCheckRequestPresentNothing uint64 = iota
	SystemGNBCUUPCounterCheckRequestPresentDRBsSubjectToCounterCheckListEUTRAN
)

// Encode implements the aper.AperMarshaller interface.
func (s *SystemGNBCUUPCounterCheckRequest) Encode(w *aper.AperWriter) (err error) {

	// 1. Write the choice index.
	if err = w.WriteChoice(uint64(s.Choice-1), 0, false); err != nil {
		return fmt.Errorf("Encode choice index failed: %w", err)
	}

	// 2. Encode the selected member.
	switch s.Choice {
	case SystemGNBCUUPCounterCheckRequestPresentDRBsSubjectToCounterCheckListEUTRAN:
		tmp_wrapper := Sequence[*DRBsSubjectToCounterCheckItemEUTRAN]{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		for _, i := range *s.DRBsSubjectToCounterCheckListEUTRAN {
			tmp_wrapper.Value = append(tmp_wrapper.Value, &i)
		}
		if err = tmp_wrapper.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBsSubjectToCounterCheckListEUTRAN failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of SystemGNBCUUPCounterCheckRequest with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *SystemGNBCUUPCounterCheckRequest) Decode(r *aper.AperReader) (err error) {

	// 1. Read the choice index.
	var choice uint64
	if choice, err = r.ReadChoice(0, false); err != nil {
		return fmt.Errorf("Read choice index failed: %w", err)
	}

	// 2. Decode the selected member.
	switch choice {
	case 0:
		// 1. Create a DECODER function for the list item, as required by ReadSequenceOf.
		itemDecoder := func(r *aper.AperReader) (*DRBsSubjectToCounterCheckItemEUTRAN, error) {
			item := new(DRBsSubjectToCounterCheckItemEUTRAN)
			if err := item.Decode(r); err != nil {
				return nil, err
			}
			return item, nil
		}

		// 2. Decode the list using the aper library's generic function.
		var decodedItems []DRBsSubjectToCounterCheckItemEUTRAN
		if decodedItems, err = aper.ReadSequenceOf(itemDecoder, r, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return fmt.Errorf("Decode DRBsSubjectToCounterCheckListEUTRAN failed: %w", err)
		}

		// 3. Assign the decoded slice of values.
		s.DRBsSubjectToCounterCheckListEUTRAN = &decodedItems
	default:
		return fmt.Errorf("Decode choice of SystemGNBCUUPCounterCheckRequest with unknown choice index %d", choice)
	}
	return nil
}
