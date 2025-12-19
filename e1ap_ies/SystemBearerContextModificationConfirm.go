package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// SystemBearerContextModificationConfirm is a generated CHOICE type.
type SystemBearerContextModificationConfirm struct {
	Choice                       uint64 `json:"-"`
	DRBConfirmModifiedListEUTRAN *DRBConfirmModifiedListEUTRAN
}

const (
	SystemBearerContextModificationConfirmPresentNothing uint64 = iota
	SystemBearerContextModificationConfirmPresentDRBConfirmModifiedListEUTRAN
)

// Encode implements the aper.AperMarshaller interface.
func (s *SystemBearerContextModificationConfirm) Encode(w *aper.AperWriter) (err error) {

	// 1. Write the choice index.
	// fmt.Printf("--- GO DEBUG: Encoding CHOICE SystemBearerContextModificationConfirm | Choice: %d, UpperBound: 0, Extensible: false\n", s.Choice-1) // UNCOMMENT FOR DEEP DEBUGGING
	if err = w.WriteChoice(uint64(s.Choice), 0, false); err != nil {
		return fmt.Errorf("Encode choice index failed for SystemBearerContextModificationConfirm: %w", err)
	}

	// 2. Encode the selected member.
	switch s.Choice {
	case SystemBearerContextModificationConfirmPresentDRBConfirmModifiedListEUTRAN:
		if err = s.DRBConfirmModifiedListEUTRAN.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBConfirmModifiedListEUTRAN failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of SystemBearerContextModificationConfirm with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *SystemBearerContextModificationConfirm) Decode(r *aper.AperReader) (err error) {

	// 1. Read the choice index (0-based) and assign it to the struct's Choice field.
	var choice uint64
	if choice, err = r.ReadChoice(0, false); err != nil {
		return fmt.Errorf("Read choice index failed: %w", err)
	}
	s.Choice = choice // Choice is 1-based from ReadChoice

	// 2. Decode the selected member.
	switch choice {
	case 1:
		s.DRBConfirmModifiedListEUTRAN = new(DRBConfirmModifiedListEUTRAN)
		if err = s.DRBConfirmModifiedListEUTRAN.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBConfirmModifiedListEUTRAN failed: %w", err)
		}
	default:
		return fmt.Errorf("Decode choice of SystemBearerContextModificationConfirm with unknown choice index %d", choice)
	}
	return nil
}
