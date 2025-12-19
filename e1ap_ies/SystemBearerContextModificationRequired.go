package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// SystemBearerContextModificationRequired is a generated CHOICE type.
type SystemBearerContextModificationRequired struct {
	Choice                        uint64 `json:"-"`
	DRBRequiredToModifyListEUTRAN *DRBRequiredToModifyListEUTRAN
	DRBRequiredToRemoveListEUTRAN *DRBRequiredToRemoveListEUTRAN
}

const (
	SystemBearerContextModificationRequiredPresentNothing uint64 = iota
	SystemBearerContextModificationRequiredPresentDRBRequiredToModifyListEUTRAN
	SystemBearerContextModificationRequiredPresentDRBRequiredToRemoveListEUTRAN
)

// Encode implements the aper.AperMarshaller interface.
func (s *SystemBearerContextModificationRequired) Encode(w *aper.AperWriter) (err error) {

	// 1. Write the choice index.
	// fmt.Printf("--- GO DEBUG: Encoding CHOICE SystemBearerContextModificationRequired | Choice: %d, UpperBound: 1, Extensible: false\n", s.Choice-1) // UNCOMMENT FOR DEEP DEBUGGING
	if err = w.WriteChoice(uint64(s.Choice), 1, false); err != nil {
		return fmt.Errorf("Encode choice index failed for SystemBearerContextModificationRequired: %w", err)
	}

	// 2. Encode the selected member.
	switch s.Choice {
	case SystemBearerContextModificationRequiredPresentDRBRequiredToModifyListEUTRAN:
		if err = s.DRBRequiredToModifyListEUTRAN.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBRequiredToModifyListEUTRAN failed: %w", err)
		}
	case SystemBearerContextModificationRequiredPresentDRBRequiredToRemoveListEUTRAN:
		if err = s.DRBRequiredToRemoveListEUTRAN.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBRequiredToRemoveListEUTRAN failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of SystemBearerContextModificationRequired with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *SystemBearerContextModificationRequired) Decode(r *aper.AperReader) (err error) {

	// 1. Read the choice index (0-based) and assign it to the struct's Choice field.
	var choice uint64
	if choice, err = r.ReadChoice(1, false); err != nil {
		return fmt.Errorf("Read choice index failed: %w", err)
	}
	s.Choice = choice // Choice is 1-based from ReadChoice

	// 2. Decode the selected member.
	switch choice {
	case 1:
		s.DRBRequiredToModifyListEUTRAN = new(DRBRequiredToModifyListEUTRAN)
		if err = s.DRBRequiredToModifyListEUTRAN.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBRequiredToModifyListEUTRAN failed: %w", err)
		}
	case 2:
		s.DRBRequiredToRemoveListEUTRAN = new(DRBRequiredToRemoveListEUTRAN)
		if err = s.DRBRequiredToRemoveListEUTRAN.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBRequiredToRemoveListEUTRAN failed: %w", err)
		}
	default:
		return fmt.Errorf("Decode choice of SystemBearerContextModificationRequired with unknown choice index %d", choice)
	}
	return nil
}
