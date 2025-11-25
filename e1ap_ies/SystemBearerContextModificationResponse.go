package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// SystemBearerContextModificationResponse is a generated CHOICE type.
type SystemBearerContextModificationResponse struct {
	Choice                        uint64 `json:"-"`
	DRBSetupModListEUTRAN         *DRBSetupModListEUTRAN
	DRBFailedModListEUTRAN        *DRBFailedModListEUTRAN
	DRBModifiedListEUTRAN         *DRBModifiedListEUTRAN
	DRBFailedToModifyListEUTRAN   *DRBFailedToModifyListEUTRAN
	RetainabilityMeasurementsInfo *RetainabilityMeasurementsInfo
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
	// fmt.Printf("--- GO DEBUG: Encoding CHOICE SystemBearerContextModificationResponse | Choice: %d, UpperBound: 4, Extensible: false\n", s.Choice-1) // UNCOMMENT FOR DEEP DEBUGGING
	if err = w.WriteChoice(uint64(s.Choice), 4, false); err != nil {
		return fmt.Errorf("Encode choice index failed for SystemBearerContextModificationResponse: %w", err)
	}

	// 2. Encode the selected member.
	switch s.Choice {
	case SystemBearerContextModificationResponsePresentDRBSetupModListEUTRAN:
		if err = s.DRBSetupModListEUTRAN.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBSetupModListEUTRAN failed: %w", err)
		}
	case SystemBearerContextModificationResponsePresentDRBFailedModListEUTRAN:
		if err = s.DRBFailedModListEUTRAN.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBFailedModListEUTRAN failed: %w", err)
		}
	case SystemBearerContextModificationResponsePresentDRBModifiedListEUTRAN:
		if err = s.DRBModifiedListEUTRAN.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBModifiedListEUTRAN failed: %w", err)
		}
	case SystemBearerContextModificationResponsePresentDRBFailedToModifyListEUTRAN:
		if err = s.DRBFailedToModifyListEUTRAN.Encode(w); err != nil {
			return fmt.Errorf("Encode DRBFailedToModifyListEUTRAN failed: %w", err)
		}
	case SystemBearerContextModificationResponsePresentRetainabilityMeasurementsInfo:
		if err = s.RetainabilityMeasurementsInfo.Encode(w); err != nil {
			return fmt.Errorf("Encode RetainabilityMeasurementsInfo failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of SystemBearerContextModificationResponse with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *SystemBearerContextModificationResponse) Decode(r *aper.AperReader) (err error) {

	// 1. Read the choice index (0-based) and assign it to the struct's Choice field.
	var choice uint64
	if choice, err = r.ReadChoice(4, false); err != nil {
		return fmt.Errorf("Read choice index failed: %w", err)
	}
	s.Choice = choice + 1 // Convert from 0-based wire format to 1-based constant format

	// 2. Decode the selected member.
	switch s.Choice {
	case 0:
		s.DRBSetupModListEUTRAN = new(DRBSetupModListEUTRAN)
		if err = s.DRBSetupModListEUTRAN.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBSetupModListEUTRAN failed: %w", err)
		}
	case 1:
		s.DRBFailedModListEUTRAN = new(DRBFailedModListEUTRAN)
		if err = s.DRBFailedModListEUTRAN.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBFailedModListEUTRAN failed: %w", err)
		}
	case 2:
		s.DRBModifiedListEUTRAN = new(DRBModifiedListEUTRAN)
		if err = s.DRBModifiedListEUTRAN.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBModifiedListEUTRAN failed: %w", err)
		}
	case 3:
		s.DRBFailedToModifyListEUTRAN = new(DRBFailedToModifyListEUTRAN)
		if err = s.DRBFailedToModifyListEUTRAN.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBFailedToModifyListEUTRAN failed: %w", err)
		}
	case 4:
		s.RetainabilityMeasurementsInfo = new(RetainabilityMeasurementsInfo)
		if err = s.RetainabilityMeasurementsInfo.Decode(r); err != nil {
			return fmt.Errorf("Decode RetainabilityMeasurementsInfo failed: %w", err)
		}
	default:
		return fmt.Errorf("Decode choice of SystemBearerContextModificationResponse with unknown choice index %d", choice)
	}
	return nil
}
