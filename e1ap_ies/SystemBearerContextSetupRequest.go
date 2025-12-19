package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// SystemBearerContextSetupRequest is a generated CHOICE type.
type SystemBearerContextSetupRequest struct {
	Choice                     uint64 `json:"-"`
	DRBToSetupListEUTRAN       *DRBToSetupListEUTRAN
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
	// fmt.Printf("--- GO DEBUG: Encoding CHOICE SystemBearerContextSetupRequest | Choice: %d, UpperBound: 2, Extensible: false\n", s.Choice-1) // UNCOMMENT FOR DEEP DEBUGGING
	if err = w.WriteChoice(uint64(s.Choice), 2, false); err != nil {
		return fmt.Errorf("Encode choice index failed for SystemBearerContextSetupRequest: %w", err)
	}

	// 2. Encode the selected member.
	switch s.Choice {
	case SystemBearerContextSetupRequestPresentDRBToSetupListEUTRAN:
		if err = s.DRBToSetupListEUTRAN.Encode(w); err != nil {
			return fmt.Errorf("encode DRBToSetupListEUTRAN failed: %w", err)
		}
	case SystemBearerContextSetupRequestPresentSubscriberProfileIDforRFP:
		if err = w.WriteInteger(int64(s.SubscriberProfileIDforRFP.Value), &aper.Constraint{Lb: 1, Ub: 256}, true); err != nil {
			return fmt.Errorf("encode SubscriberProfileIDforRFP failed: %w", err)
		}
	case SystemBearerContextSetupRequestPresentAdditionalRRMPriorityIndex:
		if err = s.AdditionalRRMPriorityIndex.Encode(w); err != nil {
			return fmt.Errorf("encode AdditionalRRMPriorityIndex failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of SystemBearerContextSetupRequest with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *SystemBearerContextSetupRequest) Decode(r *aper.AperReader) (err error) {

	// 1. Read the choice index (0-based) and assign it to the struct's Choice field.
	choice, err := r.ReadChoice(2, true)
	if err != nil {
		return fmt.Errorf("read choice index failed: %w", err)
	}
	s.Choice = choice // Choice is 1-based from ReadChoice

	// 2. Decode the selected member.
	switch choice {
	case 1:
		s.DRBToSetupListEUTRAN = new(DRBToSetupListEUTRAN)
		if err = s.DRBToSetupListEUTRAN.Decode(r); err != nil {
			return fmt.Errorf("decode DRBToSetupListEUTRAN failed: %w", err)
		}
	case 2:
		s.SubscriberProfileIDforRFP = new(SubscriberProfileIDforRFP)
		val, err := r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 256}, true)
		if err != nil {
			return fmt.Errorf("decode SubscriberProfileIDforRFP failed: %w", err)
		}
		s.SubscriberProfileIDforRFP.Value = aper.Integer(val)
	case 3:
		s.AdditionalRRMPriorityIndex = new(AdditionalRRMPriorityIndex)
		if err = s.AdditionalRRMPriorityIndex.Decode(r); err != nil {
			return fmt.Errorf("decode AdditionalRRMPriorityIndex failed: %w", err)
		}
	default:
		return fmt.Errorf("decode choice of SystemBearerContextSetupRequest with unknown choice index %d", choice)
	}
	return nil
}
