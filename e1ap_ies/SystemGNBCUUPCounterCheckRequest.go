package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// SystemGNBCUUPCounterCheckRequest is a generated CHOICE type.
type SystemGNBCUUPCounterCheckRequest struct {
	Choice                              uint64 `json:"-"`
	DRBsSubjectToCounterCheckListEUTRAN *DRBsSubjectToCounterCheckListEUTRAN
}

const (
	SystemGNBCUUPCounterCheckRequestPresentNothing uint64 = iota
	SystemGNBCUUPCounterCheckRequestPresentDRBsSubjectToCounterCheckListEUTRAN
)

// Encode implements the aper.AperMarshaller interface.
func (s *SystemGNBCUUPCounterCheckRequest) Encode(w *aper.AperWriter) (err error) {

	// 1. Write the choice index.
	// fmt.Printf("--- GO DEBUG: Encoding CHOICE SystemGNBCUUPCounterCheckRequest | Choice: %d, UpperBound: 0, Extensible: false\n", s.Choice-1) // UNCOMMENT FOR DEEP DEBUGGING
	if err = w.WriteChoice(uint64(s.Choice), 0, false); err != nil {
		return fmt.Errorf("Encode choice index failed for SystemGNBCUUPCounterCheckRequest: %w", err)
	}

	// 2. Encode the selected member.
	switch s.Choice {
	case SystemGNBCUUPCounterCheckRequestPresentDRBsSubjectToCounterCheckListEUTRAN:
		if err = s.DRBsSubjectToCounterCheckListEUTRAN.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBsSubjectToCounterCheckListEUTRAN failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of SystemGNBCUUPCounterCheckRequest with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *SystemGNBCUUPCounterCheckRequest) Decode(r *aper.AperReader) (err error) {

	// 1. Read the choice index (0-based) and assign it to the struct's Choice field.
	var choice uint64
	if choice, err = r.ReadChoice(0, false); err != nil {
		return fmt.Errorf("Read choice index failed: %w", err)
	}
	s.Choice = choice // Choice is 1-based from ReadChoice

	// 2. Decode the selected member.
	switch choice {
	case 1:
		s.DRBsSubjectToCounterCheckListEUTRAN = new(DRBsSubjectToCounterCheckListEUTRAN)
		if err = s.DRBsSubjectToCounterCheckListEUTRAN.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBsSubjectToCounterCheckListEUTRAN failed: %w", err)
		}
	default:
		return fmt.Errorf("Decode choice of SystemGNBCUUPCounterCheckRequest with unknown choice index %d", choice)
	}
	return nil
}
