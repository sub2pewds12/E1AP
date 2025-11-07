package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// SystemBearerContextModificationRequest is a generated CHOICE type.
type SystemBearerContextModificationRequest struct {
	Choice                     uint64 `json:"-"`
	DRBToSetupModListEUTRAN    *DRBToSetupModListEUTRAN
	DRBToModifyListEUTRAN      *DRBToModifyListEUTRAN
	DRBToRemoveListEUTRAN      *DRBToRemoveListEUTRAN
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
	if err = w.WriteChoice(uint64(s.Choice-1), 4, true); err != nil {
		return fmt.Errorf("Encode choice index failed: %w", err)
	}

	// 2. Encode the selected member.
	switch s.Choice {
	case SystemBearerContextModificationRequestPresentDRBToSetupModListEUTRAN:
		if err = s.DRBToSetupModListEUTRAN.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBToSetupModListEUTRAN failed: %w", err)
		}
	case SystemBearerContextModificationRequestPresentDRBToModifyListEUTRAN:
		if err = s.DRBToModifyListEUTRAN.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBToModifyListEUTRAN failed: %w", err)
		}
	case SystemBearerContextModificationRequestPresentDRBToRemoveListEUTRAN:
		if err = s.DRBToRemoveListEUTRAN.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBToRemoveListEUTRAN failed: %w", err)
		}
	case SystemBearerContextModificationRequestPresentSubscriberProfileIDforRFP:
		if err = w.WriteInteger(int64(s.SubscriberProfileIDforRFP.Value), &aper.Constraint{Lb: 1, Ub: 256}, true); err != nil {
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

	// 1. Read the choice index (0-based) and assign it to the struct's Choice field.
	var choice uint64
	if choice, err = r.ReadChoice(4, true); err != nil {
		return fmt.Errorf("Read choice index failed: %w", err)
	}
	s.Choice = choice + 1 // Convert from 0-based wire format to 1-based constant format

	// 2. Decode the selected member.
	switch s.Choice {
	case 0:
		s.DRBToSetupModListEUTRAN = new(DRBToSetupModListEUTRAN)
		if err = s.DRBToSetupModListEUTRAN.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBToSetupModListEUTRAN failed: %w", err)
		}
	case 1:
		s.DRBToModifyListEUTRAN = new(DRBToModifyListEUTRAN)
		if err = s.DRBToModifyListEUTRAN.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBToModifyListEUTRAN failed: %w", err)
		}
	case 2:
		s.DRBToRemoveListEUTRAN = new(DRBToRemoveListEUTRAN)
		if err = s.DRBToRemoveListEUTRAN.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBToRemoveListEUTRAN failed: %w", err)
		}
	case 3:
		s.SubscriberProfileIDforRFP = new(SubscriberProfileIDforRFP)
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 256}, true); err != nil {
			return fmt.Errorf("Decode SubscriberProfileIDforRFP failed: %w", err)
		}
		s.SubscriberProfileIDforRFP.Value = aper.Integer(val)
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
