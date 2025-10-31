package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// SystemBearerContextSetupResponse is a generated CHOICE type.
type SystemBearerContextSetupResponse struct {
	Choice              uint64 `json:"-"`
	DRBSetupListEUTRAN  *[]DRBSetupItemEUTRAN
	DRBFailedListEUTRAN *[]DRBFailedItemEUTRAN
}

const (
	SystemBearerContextSetupResponsePresentNothing uint64 = iota
	SystemBearerContextSetupResponsePresentDRBSetupListEUTRAN
	SystemBearerContextSetupResponsePresentDRBFailedListEUTRAN
)

// Encode implements the aper.AperMarshaller interface.
func (s *SystemBearerContextSetupResponse) Encode(w *aper.AperWriter) (err error) {

	// 1. Write the choice index.
	if err = w.WriteChoice(uint64(s.Choice-1), 1, false); err != nil {
		return fmt.Errorf("Encode choice index failed: %w", err)
	}

	// 2. Encode the selected member.
	switch s.Choice {
	case SystemBearerContextSetupResponsePresentDRBSetupListEUTRAN:
		tmp_wrapper := Sequence[*DRBSetupItemEUTRAN]{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		for _, i := range *s.DRBSetupListEUTRAN {
			tmp_wrapper.Value = append(tmp_wrapper.Value, &i)
		}
		if err = tmp_wrapper.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBSetupListEUTRAN failed: %w", err)
		}
	case SystemBearerContextSetupResponsePresentDRBFailedListEUTRAN:
		tmp_wrapper := Sequence[*DRBFailedItemEUTRAN]{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		for _, i := range *s.DRBFailedListEUTRAN {
			tmp_wrapper.Value = append(tmp_wrapper.Value, &i)
		}
		if err = tmp_wrapper.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBFailedListEUTRAN failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of SystemBearerContextSetupResponse with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *SystemBearerContextSetupResponse) Decode(r *aper.AperReader) (err error) {

	// 1. Read the choice index.
	var choice uint64
	if choice, err = r.ReadChoice(1, false); err != nil {
		return fmt.Errorf("Read choice index failed: %w", err)
	}

	// 2. Decode the selected member.
	switch choice {
	case 0:
		// 1. Create a DECODER function for the list item, as required by ReadSequenceOf.
		itemDecoder := func(r *aper.AperReader) (*DRBSetupItemEUTRAN, error) {
			item := new(DRBSetupItemEUTRAN)
			if err := item.Decode(r); err != nil {
				return nil, err
			}
			return item, nil
		}

		// 2. Decode the list using the aper library's generic function.
		var decodedItems []DRBSetupItemEUTRAN
		if decodedItems, err = aper.ReadSequenceOf(itemDecoder, r, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return fmt.Errorf("Decode DRBSetupListEUTRAN failed: %w", err)
		}

		// 3. Assign the decoded slice of values.
		s.DRBSetupListEUTRAN = &decodedItems
	case 1:
		// 1. Create a DECODER function for the list item, as required by ReadSequenceOf.
		itemDecoder := func(r *aper.AperReader) (*DRBFailedItemEUTRAN, error) {
			item := new(DRBFailedItemEUTRAN)
			if err := item.Decode(r); err != nil {
				return nil, err
			}
			return item, nil
		}

		// 2. Decode the list using the aper library's generic function.
		var decodedItems []DRBFailedItemEUTRAN
		if decodedItems, err = aper.ReadSequenceOf(itemDecoder, r, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return fmt.Errorf("Decode DRBFailedListEUTRAN failed: %w", err)
		}

		// 3. Assign the decoded slice of values.
		s.DRBFailedListEUTRAN = &decodedItems
	default:
		return fmt.Errorf("Decode choice of SystemBearerContextSetupResponse with unknown choice index %d", choice)
	}
	return nil
}
