package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// SystemBearerContextSetupRequest is a generated CHOICE type.
type SystemBearerContextSetupRequest struct {
	Choice                     uint64 `json:"-"`
	DRBToSetupListEUTRAN       *[]DRBToSetupItemEUTRAN
	SubscriberProfileIDforRFP  *SubscriberProfileIDforRFP
	AdditionalRRMPriorityIndex *AdditionalRRMPriorityIndex
}

const (
	SystemBearerContextSetupRequestPresentNothing uint64 = iota
	SystemBearerContextSetupRequestPresentDRBToSetupListEUTRAN
	SystemBearerContextSetupRequestPresentSubscriberProfileIDforRFP
	SystemBearerContextSetupRequestPresentAdditionalRRMPriorityIndex
)

// Encode implements the aper.AperMarshaller interface.
func (s *SystemBearerContextSetupRequest) Encode(w *aper.AperWriter) (err error) {

	// 1. Write the choice index.
	if err = w.WriteChoice(uint64(s.Choice-1), 2, false); err != nil {
		return fmt.Errorf("Encode choice index failed: %w", err)
	}

	// 2. Encode the selected member.
	switch s.Choice {
	case SystemBearerContextSetupRequestPresentDRBToSetupListEUTRAN:
		tmp_wrapper := Sequence[*DRBToSetupItemEUTRAN]{
			c:   aper.Constraint{Lb: 0, Ub: 0},
			ext: false,
		}
		for _, i := range *s.DRBToSetupListEUTRAN {
			tmp_wrapper.Value = append(tmp_wrapper.Value, &i)
		}
		if err = tmp_wrapper.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBToSetupListEUTRAN failed: %w", err)
		}
	case SystemBearerContextSetupRequestPresentSubscriberProfileIDforRFP:
		if err = s.SubscriberProfileIDforRFP.Encode(w); err != nil {
			return fmt.Errorf("Encode SubscriberProfileIDforRFP failed: %w", err)
		}
	case SystemBearerContextSetupRequestPresentAdditionalRRMPriorityIndex:
		if err = s.AdditionalRRMPriorityIndex.Encode(w); err != nil {
			return fmt.Errorf("Encode AdditionalRRMPriorityIndex failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of SystemBearerContextSetupRequest with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *SystemBearerContextSetupRequest) Decode(r *aper.AperReader) (err error) {

	// 1. Read the choice index.
	var choice uint64
	if choice, err = r.ReadChoice(2, false); err != nil {
		return fmt.Errorf("Read choice index failed: %w", err)
	}

	// 2. Decode the selected member.
	switch choice {
	case 0:
		// 1. Create a DECODER function for the list item, as required by ReadSequenceOf.
		itemDecoder := func(r *aper.AperReader) (*DRBToSetupItemEUTRAN, error) {
			item := new(DRBToSetupItemEUTRAN)
			if err := item.Decode(r); err != nil {
				return nil, err
			}
			return item, nil
		}

		// 2. Decode the list using the aper library's generic function.
		var decodedItems []DRBToSetupItemEUTRAN
		if decodedItems, err = aper.ReadSequenceOf(itemDecoder, r, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return fmt.Errorf("Decode DRBToSetupListEUTRAN failed: %w", err)
		}

		// 3. Assign the decoded slice of values.
		s.DRBToSetupListEUTRAN = &decodedItems
	case 1:
		s.SubscriberProfileIDforRFP = new(SubscriberProfileIDforRFP)
		if err = s.SubscriberProfileIDforRFP.Decode(r); err != nil {
			return fmt.Errorf("Decode SubscriberProfileIDforRFP failed: %w", err)
		}
	case 2:
		s.AdditionalRRMPriorityIndex = new(AdditionalRRMPriorityIndex)
		if err = s.AdditionalRRMPriorityIndex.Decode(r); err != nil {
			return fmt.Errorf("Decode AdditionalRRMPriorityIndex failed: %w", err)
		}
	default:
		return fmt.Errorf("Decode choice of SystemBearerContextSetupRequest with unknown choice index %d", choice)
	}
	return nil
}
