package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// SystemBearerContextModificationRequest is a generated CHOICE type.
type SystemBearerContextModificationRequest struct {
	Choice                     uint64 `json:"-"`
	DRBToSetupModListEUTRAN    *[]DRBToSetupModItemEUTRAN
	DRBToModifyListEUTRAN      *[]DRBToModifyItemEUTRAN
	DRBToRemoveListEUTRAN      *[]DRBToRemoveItemEUTRAN
	SubscriberProfileIDforRFP  *SubscriberProfileIDforRFP
	AdditionalRRMPriorityIndex *AdditionalRRMPriorityIndex
}

const (
	SystemBearerContextModificationRequestPresentNothing uint64 = iota
	SystemBearerContextModificationRequestPresentDRBToSetupModListEUTRAN
	SystemBearerContextModificationRequestPresentDRBToModifyListEUTRAN
	SystemBearerContextModificationRequestPresentDRBToRemoveListEUTRAN
	SystemBearerContextModificationRequestPresentSubscriberProfileIDforRFP
	SystemBearerContextModificationRequestPresentAdditionalRRMPriorityIndex
)

// Encode implements the aper.AperMarshaller interface.
func (s *SystemBearerContextModificationRequest) Encode(w *aper.AperWriter) (err error) {

	// 1. Write the choice index.
	if err = w.WriteChoice(uint64(s.Choice-1), 4, false); err != nil {
		return fmt.Errorf("Encode choice index failed: %w", err)
	}

	// 2. Encode the selected member.
	switch s.Choice {
	case SystemBearerContextModificationRequestPresentDRBToSetupModListEUTRAN:
		tmp_wrapper := Sequence[*DRBToSetupModItemEUTRAN]{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		for _, i := range *s.DRBToSetupModListEUTRAN {
			tmp_wrapper.Value = append(tmp_wrapper.Value, &i)
		}
		if err = tmp_wrapper.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBToSetupModListEUTRAN failed: %w", err)
		}
	case SystemBearerContextModificationRequestPresentDRBToModifyListEUTRAN:
		tmp_wrapper := Sequence[*DRBToModifyItemEUTRAN]{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		for _, i := range *s.DRBToModifyListEUTRAN {
			tmp_wrapper.Value = append(tmp_wrapper.Value, &i)
		}
		if err = tmp_wrapper.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBToModifyListEUTRAN failed: %w", err)
		}
	case SystemBearerContextModificationRequestPresentDRBToRemoveListEUTRAN:
		tmp_wrapper := Sequence[*DRBToRemoveItemEUTRAN]{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		for _, i := range *s.DRBToRemoveListEUTRAN {
			tmp_wrapper.Value = append(tmp_wrapper.Value, &i)
		}
		if err = tmp_wrapper.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBToRemoveListEUTRAN failed: %w", err)
		}
	case SystemBearerContextModificationRequestPresentSubscriberProfileIDforRFP:
		if err = s.SubscriberProfileIDforRFP.Encode(w); err != nil {
			return fmt.Errorf("Encode SubscriberProfileIDforRFP failed: %w", err)
		}
	case SystemBearerContextModificationRequestPresentAdditionalRRMPriorityIndex:
		if err = s.AdditionalRRMPriorityIndex.Encode(w); err != nil {
			return fmt.Errorf("Encode AdditionalRRMPriorityIndex failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of SystemBearerContextModificationRequest with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *SystemBearerContextModificationRequest) Decode(r *aper.AperReader) (err error) {

	// 1. Read the choice index.
	var choice uint64
	if choice, err = r.ReadChoice(4, false); err != nil {
		return fmt.Errorf("Read choice index failed: %w", err)
	}

	// 2. Decode the selected member.
	switch choice {
	case 0:
		// 1. Create a DECODER function for the list item, as required by ReadSequenceOf.
		itemDecoder := func(r *aper.AperReader) (*DRBToSetupModItemEUTRAN, error) {
			item := new(DRBToSetupModItemEUTRAN)
			if err := item.Decode(r); err != nil {
				return nil, err
			}
			return item, nil
		}

		// 2. Decode the list using the aper library's generic function.
		var decodedItems []DRBToSetupModItemEUTRAN
		if decodedItems, err = aper.ReadSequenceOf(itemDecoder, r, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return fmt.Errorf("Decode DRBToSetupModListEUTRAN failed: %w", err)
		}

		// 3. Assign the decoded slice of values.
		s.DRBToSetupModListEUTRAN = &decodedItems
	case 1:
		// 1. Create a DECODER function for the list item, as required by ReadSequenceOf.
		itemDecoder := func(r *aper.AperReader) (*DRBToModifyItemEUTRAN, error) {
			item := new(DRBToModifyItemEUTRAN)
			if err := item.Decode(r); err != nil {
				return nil, err
			}
			return item, nil
		}

		// 2. Decode the list using the aper library's generic function.
		var decodedItems []DRBToModifyItemEUTRAN
		if decodedItems, err = aper.ReadSequenceOf(itemDecoder, r, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return fmt.Errorf("Decode DRBToModifyListEUTRAN failed: %w", err)
		}

		// 3. Assign the decoded slice of values.
		s.DRBToModifyListEUTRAN = &decodedItems
	case 2:
		// 1. Create a DECODER function for the list item, as required by ReadSequenceOf.
		itemDecoder := func(r *aper.AperReader) (*DRBToRemoveItemEUTRAN, error) {
			item := new(DRBToRemoveItemEUTRAN)
			if err := item.Decode(r); err != nil {
				return nil, err
			}
			return item, nil
		}

		// 2. Decode the list using the aper library's generic function.
		var decodedItems []DRBToRemoveItemEUTRAN
		if decodedItems, err = aper.ReadSequenceOf(itemDecoder, r, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return fmt.Errorf("Decode DRBToRemoveListEUTRAN failed: %w", err)
		}

		// 3. Assign the decoded slice of values.
		s.DRBToRemoveListEUTRAN = &decodedItems
	case 3:
		s.SubscriberProfileIDforRFP = new(SubscriberProfileIDforRFP)
		if err = s.SubscriberProfileIDforRFP.Decode(r); err != nil {
			return fmt.Errorf("Decode SubscriberProfileIDforRFP failed: %w", err)
		}
	case 4:
		s.AdditionalRRMPriorityIndex = new(AdditionalRRMPriorityIndex)
		if err = s.AdditionalRRMPriorityIndex.Decode(r); err != nil {
			return fmt.Errorf("Decode AdditionalRRMPriorityIndex failed: %w", err)
		}
	default:
		return fmt.Errorf("Decode choice of SystemBearerContextModificationRequest with unknown choice index %d", choice)
	}
	return nil
}
