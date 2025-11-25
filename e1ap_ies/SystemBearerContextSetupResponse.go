package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// SystemBearerContextSetupResponse is a generated CHOICE type.
type SystemBearerContextSetupResponse struct {
	Choice              uint64 `json:"-"`
	DRBSetupListEUTRAN  *DRBSetupListEUTRAN
	DRBFailedListEUTRAN *DRBFailedListEUTRAN
}

const (
	SystemBearerContextSetupResponsePresentNothing uint64 = iota
	SystemBearerContextSetupResponsePresentDRBSetupListEUTRAN
	SystemBearerContextSetupResponsePresentDRBFailedListEUTRAN
)

// Encode implements the aper.AperMarshaller interface.
func (s *SystemBearerContextSetupResponse) Encode(w *aper.AperWriter) (err error) {

	// 1. Write the choice index.
	// fmt.Printf("--- GO DEBUG: Encoding CHOICE SystemBearerContextSetupResponse | Choice: %d, UpperBound: 1, Extensible: false\n", s.Choice-1) // UNCOMMENT FOR DEEP DEBUGGING
	if err = w.WriteChoice(uint64(s.Choice), 1, false); err != nil {
		return fmt.Errorf("Encode choice index failed for SystemBearerContextSetupResponse: %w", err)
	}

	// 2. Encode the selected member.
	switch s.Choice {
	case SystemBearerContextSetupResponsePresentDRBSetupListEUTRAN:
		if err = s.DRBSetupListEUTRAN.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBSetupListEUTRAN failed: %w", err)
		}
	case SystemBearerContextSetupResponsePresentDRBFailedListEUTRAN:
		if err = s.DRBFailedListEUTRAN.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBFailedListEUTRAN failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of SystemBearerContextSetupResponse with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *SystemBearerContextSetupResponse) Decode(r *aper.AperReader) (err error) {

	// 1. Read the choice index (0-based) and assign it to the struct's Choice field.
	var choice uint64
	if choice, err = r.ReadChoice(1, false); err != nil {
		return fmt.Errorf("Read choice index failed: %w", err)
	}
	s.Choice = choice + 1 // Convert from 0-based wire format to 1-based constant format

	// 2. Decode the selected member.
	switch s.Choice {
	case 0:
		s.DRBSetupListEUTRAN = new(DRBSetupListEUTRAN)
		if err = s.DRBSetupListEUTRAN.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBSetupListEUTRAN failed: %w", err)
		}
	case 1:
		s.DRBFailedListEUTRAN = new(DRBFailedListEUTRAN)
		if err = s.DRBFailedListEUTRAN.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBFailedListEUTRAN failed: %w", err)
		}
	default:
		return fmt.Errorf("Decode choice of SystemBearerContextSetupResponse with unknown choice index %d", choice)
	}
	return nil
}
