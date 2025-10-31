package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// SystemBearerContextModificationResponse is a generated CHOICE type.
type SystemBearerContextModificationResponse struct {
	Choice                        uint64 `json:"-"`
	DRBSetupModListEUTRAN         *[]DRBSetupModItemEUTRAN
	DRBFailedModListEUTRAN        *[]DRBFailedModItemEUTRAN
	DRBModifiedListEUTRAN         *[]DRBModifiedItemEUTRAN
	DRBFailedToModifyListEUTRAN   *[]DRBFailedToModifyItemEUTRAN
	RetainabilityMeasurementsInfo *[]DRBRemovedItem
}

const (
	SystemBearerContextModificationResponsePresentNothing uint64 = iota
	SystemBearerContextModificationResponsePresentDRBSetupModListEUTRAN
	SystemBearerContextModificationResponsePresentDRBFailedModListEUTRAN
	SystemBearerContextModificationResponsePresentDRBModifiedListEUTRAN
	SystemBearerContextModificationResponsePresentDRBFailedToModifyListEUTRAN
	SystemBearerContextModificationResponsePresentRetainabilityMeasurementsInfo
)

// Encode implements the aper.AperMarshaller interface.
func (s *SystemBearerContextModificationResponse) Encode(w *aper.AperWriter) (err error) {

	// 1. Write the choice index.
	if err = w.WriteChoice(uint64(s.Choice-1), 4, false); err != nil {
		return fmt.Errorf("Encode choice index failed: %w", err)
	}

	// 2. Encode the selected member.
	switch s.Choice {
	case SystemBearerContextModificationResponsePresentDRBSetupModListEUTRAN:
		tmp_wrapper := Sequence[*DRBSetupModItemEUTRAN]{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		for _, i := range *s.DRBSetupModListEUTRAN {
			tmp_wrapper.Value = append(tmp_wrapper.Value, &i)
		}
		if err = tmp_wrapper.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBSetupModListEUTRAN failed: %w", err)
		}
	case SystemBearerContextModificationResponsePresentDRBFailedModListEUTRAN:
		tmp_wrapper := Sequence[*DRBFailedModItemEUTRAN]{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		for _, i := range *s.DRBFailedModListEUTRAN {
			tmp_wrapper.Value = append(tmp_wrapper.Value, &i)
		}
		if err = tmp_wrapper.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBFailedModListEUTRAN failed: %w", err)
		}
	case SystemBearerContextModificationResponsePresentDRBModifiedListEUTRAN:
		tmp_wrapper := Sequence[*DRBModifiedItemEUTRAN]{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		for _, i := range *s.DRBModifiedListEUTRAN {
			tmp_wrapper.Value = append(tmp_wrapper.Value, &i)
		}
		if err = tmp_wrapper.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBModifiedListEUTRAN failed: %w", err)
		}
	case SystemBearerContextModificationResponsePresentDRBFailedToModifyListEUTRAN:
		tmp_wrapper := Sequence[*DRBFailedToModifyItemEUTRAN]{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		for _, i := range *s.DRBFailedToModifyListEUTRAN {
			tmp_wrapper.Value = append(tmp_wrapper.Value, &i)
		}
		if err = tmp_wrapper.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBFailedToModifyListEUTRAN failed: %w", err)
		}
	case SystemBearerContextModificationResponsePresentRetainabilityMeasurementsInfo:
		tmp_wrapper := Sequence[*DRBRemovedItem]{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		for _, i := range *s.RetainabilityMeasurementsInfo {
			tmp_wrapper.Value = append(tmp_wrapper.Value, &i)
		}
		if err = tmp_wrapper.Encode(w); err != nil {
			return fmt.Errorf("Encode RetainabilityMeasurementsInfo failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of SystemBearerContextModificationResponse with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *SystemBearerContextModificationResponse) Decode(r *aper.AperReader) (err error) {

	// 1. Read the choice index.
	var choice uint64
	if choice, err = r.ReadChoice(4, false); err != nil {
		return fmt.Errorf("Read choice index failed: %w", err)
	}

	// 2. Decode the selected member.
	switch choice {
	case 0:
		// 1. Create a DECODER function for the list item, as required by ReadSequenceOf.
		itemDecoder := func(r *aper.AperReader) (*DRBSetupModItemEUTRAN, error) {
			item := new(DRBSetupModItemEUTRAN)
			if err := item.Decode(r); err != nil {
				return nil, err
			}
			return item, nil
		}

		// 2. Decode the list using the aper library's generic function.
		var decodedItems []DRBSetupModItemEUTRAN
		if decodedItems, err = aper.ReadSequenceOf(itemDecoder, r, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return fmt.Errorf("Decode DRBSetupModListEUTRAN failed: %w", err)
		}

		// 3. Assign the decoded slice of values.
		s.DRBSetupModListEUTRAN = &decodedItems
	case 1:
		// 1. Create a DECODER function for the list item, as required by ReadSequenceOf.
		itemDecoder := func(r *aper.AperReader) (*DRBFailedModItemEUTRAN, error) {
			item := new(DRBFailedModItemEUTRAN)
			if err := item.Decode(r); err != nil {
				return nil, err
			}
			return item, nil
		}

		// 2. Decode the list using the aper library's generic function.
		var decodedItems []DRBFailedModItemEUTRAN
		if decodedItems, err = aper.ReadSequenceOf(itemDecoder, r, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return fmt.Errorf("Decode DRBFailedModListEUTRAN failed: %w", err)
		}

		// 3. Assign the decoded slice of values.
		s.DRBFailedModListEUTRAN = &decodedItems
	case 2:
		// 1. Create a DECODER function for the list item, as required by ReadSequenceOf.
		itemDecoder := func(r *aper.AperReader) (*DRBModifiedItemEUTRAN, error) {
			item := new(DRBModifiedItemEUTRAN)
			if err := item.Decode(r); err != nil {
				return nil, err
			}
			return item, nil
		}

		// 2. Decode the list using the aper library's generic function.
		var decodedItems []DRBModifiedItemEUTRAN
		if decodedItems, err = aper.ReadSequenceOf(itemDecoder, r, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return fmt.Errorf("Decode DRBModifiedListEUTRAN failed: %w", err)
		}

		// 3. Assign the decoded slice of values.
		s.DRBModifiedListEUTRAN = &decodedItems
	case 3:
		// 1. Create a DECODER function for the list item, as required by ReadSequenceOf.
		itemDecoder := func(r *aper.AperReader) (*DRBFailedToModifyItemEUTRAN, error) {
			item := new(DRBFailedToModifyItemEUTRAN)
			if err := item.Decode(r); err != nil {
				return nil, err
			}
			return item, nil
		}

		// 2. Decode the list using the aper library's generic function.
		var decodedItems []DRBFailedToModifyItemEUTRAN
		if decodedItems, err = aper.ReadSequenceOf(itemDecoder, r, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return fmt.Errorf("Decode DRBFailedToModifyListEUTRAN failed: %w", err)
		}

		// 3. Assign the decoded slice of values.
		s.DRBFailedToModifyListEUTRAN = &decodedItems
	case 4:
		// 1. Create a DECODER function for the list item, as required by ReadSequenceOf.
		itemDecoder := func(r *aper.AperReader) (*DRBRemovedItem, error) {
			item := new(DRBRemovedItem)
			if err := item.Decode(r); err != nil {
				return nil, err
			}
			return item, nil
		}

		// 2. Decode the list using the aper library's generic function.
		var decodedItems []DRBRemovedItem
		if decodedItems, err = aper.ReadSequenceOf(itemDecoder, r, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return fmt.Errorf("Decode RetainabilityMeasurementsInfo failed: %w", err)
		}

		// 3. Assign the decoded slice of values.
		s.RetainabilityMeasurementsInfo = &decodedItems
	default:
		return fmt.Errorf("Decode choice of SystemBearerContextModificationResponse with unknown choice index %d", choice)
	}
	return nil
}
