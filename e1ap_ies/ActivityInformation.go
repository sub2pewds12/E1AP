package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// ActivityInformation is a generated CHOICE type.
type ActivityInformation struct {
	Choice                         uint64 `json:"-"`
	DRBActivityList                *[]DRBActivityItem
	PDUSessionResourceActivityList *[]PDUSessionResourceActivityItem
	UEActivity                     *UEActivity
	ChoiceExtension                *ProtocolIESingleContainer
}

const (
	ActivityInformationPresentNothing uint64 = iota
	ActivityInformationPresentDRBActivityList
	ActivityInformationPresentPDUSessionResourceActivityList
	ActivityInformationPresentUEActivity
	ActivityInformationPresentChoiceExtension
)

// Encode implements the aper.AperMarshaller interface.
func (s *ActivityInformation) Encode(w *aper.AperWriter) (err error) {

	// 1. Write the choice index.
	if err = w.WriteChoice(uint64(s.Choice-1), 3, false); err != nil {
		return fmt.Errorf("Encode choice index failed: %w", err)
	}

	// 2. Encode the selected member.
	switch s.Choice {
	case ActivityInformationPresentDRBActivityList:
		tmp_wrapper := Sequence[*DRBActivityItem]{
			c:   aper.Constraint{Lb: 1, Ub: MaxnoofDRBs},
			ext: false,
		}
		for _, i := range *s.DRBActivityList {
			tmp_wrapper.Value = append(tmp_wrapper.Value, &i)
		}
		if err = tmp_wrapper.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBActivityList failed: %w", err)
		}
	case ActivityInformationPresentPDUSessionResourceActivityList:
		tmp_wrapper := Sequence[*PDUSessionResourceActivityItem]{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		for _, i := range *s.PDUSessionResourceActivityList {
			tmp_wrapper.Value = append(tmp_wrapper.Value, &i)
		}
		if err = tmp_wrapper.Encode(w); err != nil {
			return fmt.Errorf("Encode PDUSessionResourceActivityList failed: %w", err)
		}
	case ActivityInformationPresentUEActivity:
		if err = s.UEActivity.Encode(w); err != nil {
			return fmt.Errorf("Encode UEActivity failed: %w", err)
		}
	case ActivityInformationPresentChoiceExtension:
		if err = s.ChoiceExtension.Encode(w); err != nil {
			return fmt.Errorf("Encode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of ActivityInformation with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ActivityInformation) Decode(r *aper.AperReader) (err error) {

	// 1. Read the choice index.
	var choice uint64
	if choice, err = r.ReadChoice(3, false); err != nil {
		return fmt.Errorf("Read choice index failed: %w", err)
	}

	// 2. Decode the selected member.
	switch choice {
	case 0:
		// 1. Create a DECODER function for the list item, as required by ReadSequenceOf.
		itemDecoder := func(r *aper.AperReader) (*DRBActivityItem, error) {
			item := new(DRBActivityItem)
			if err := item.Decode(r); err != nil {
				return nil, err
			}
			return item, nil
		}

		// 2. Decode the list using the aper library's generic function.
		var decodedItems []DRBActivityItem
		if decodedItems, err = aper.ReadSequenceOf(itemDecoder, r, &aper.Constraint{Lb: 1, Ub: MaxnoofDRBs}, false); err != nil {
			return fmt.Errorf("Decode DRBActivityList failed: %w", err)
		}

		// 3. Assign the decoded slice of values.
		s.DRBActivityList = &decodedItems
	case 1:
		// 1. Create a DECODER function for the list item, as required by ReadSequenceOf.
		itemDecoder := func(r *aper.AperReader) (*PDUSessionResourceActivityItem, error) {
			item := new(PDUSessionResourceActivityItem)
			if err := item.Decode(r); err != nil {
				return nil, err
			}
			return item, nil
		}

		// 2. Decode the list using the aper library's generic function.
		var decodedItems []PDUSessionResourceActivityItem
		if decodedItems, err = aper.ReadSequenceOf(itemDecoder, r, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return fmt.Errorf("Decode PDUSessionResourceActivityList failed: %w", err)
		}

		// 3. Assign the decoded slice of values.
		s.PDUSessionResourceActivityList = &decodedItems
	case 2:
		s.UEActivity = new(UEActivity)
		if err = s.UEActivity.Decode(r); err != nil {
			return fmt.Errorf("Decode UEActivity failed: %w", err)
		}
	case 3:
		s.ChoiceExtension = new(ProtocolIESingleContainer)
		if err = s.ChoiceExtension.Decode(r); err != nil {
			return fmt.Errorf("Decode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("Decode choice of ActivityInformation with unknown choice index %d", choice)
	}
	return nil
}
