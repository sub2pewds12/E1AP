package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// ActivityInformation is a generated CHOICE type.
type ActivityInformation struct {
	Choice                         uint64 `json:"-"`
	DRBActivityList                *DRBActivityList
	PDUSessionResourceActivityList *PDUSessionResourceActivityList
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
	// fmt.Printf("--- GO DEBUG: Encoding CHOICE ActivityInformation | Choice: %d, UpperBound: 3, Extensible: false\n", s.Choice-1) // UNCOMMENT FOR DEEP DEBUGGING
	if err = w.WriteChoice(uint64(s.Choice), 3, false); err != nil {
		return fmt.Errorf("Encode choice index failed for ActivityInformation: %w", err)
	}

	// 2. Encode the selected member.
	switch s.Choice {
	case ActivityInformationPresentDRBActivityList:
		if err = s.DRBActivityList.Encode(w); err != nil {
			return fmt.Errorf("encode DRBActivityList failed: %w", err)
		}
	case ActivityInformationPresentPDUSessionResourceActivityList:
		if err = s.PDUSessionResourceActivityList.Encode(w); err != nil {
			return fmt.Errorf("encode PDUSessionResourceActivityList failed: %w", err)
		}
	case ActivityInformationPresentUEActivity:
		if err = s.UEActivity.Encode(w); err != nil {
			return fmt.Errorf("encode UEActivity failed: %w", err)
		}
	case ActivityInformationPresentChoiceExtension:
		if err = s.ChoiceExtension.Encode(w); err != nil {
			return fmt.Errorf("encode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of ActivityInformation with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ActivityInformation) Decode(r *aper.AperReader) (err error) {

	// 1. Read the choice index (0-based) and assign it to the struct's Choice field.
	choice, err := r.ReadChoice(3, false)
	if err != nil {
		return fmt.Errorf("read choice index failed: %w", err)
	}
	s.Choice = choice // Choice is 1-based from ReadChoice

	// 2. Decode the selected member.
	switch choice {
	case 1:
		s.DRBActivityList = new(DRBActivityList)
		if err = s.DRBActivityList.Decode(r); err != nil {
			return fmt.Errorf("decode DRBActivityList failed: %w", err)
		}
	case 2:
		s.PDUSessionResourceActivityList = new(PDUSessionResourceActivityList)
		if err = s.PDUSessionResourceActivityList.Decode(r); err != nil {
			return fmt.Errorf("decode PDUSessionResourceActivityList failed: %w", err)
		}
	case 3:
		s.UEActivity = new(UEActivity)
		if err = s.UEActivity.Decode(r); err != nil {
			return fmt.Errorf("decode UEActivity failed: %w", err)
		}
	case 4:
		s.ChoiceExtension = new(ProtocolIESingleContainer)
		if err = s.ChoiceExtension.Decode(r); err != nil {
			return fmt.Errorf("decode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("decode choice of ActivityInformation with unknown choice index %d", choice)
	}
	return nil
}
