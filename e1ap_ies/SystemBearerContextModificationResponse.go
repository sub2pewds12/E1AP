package e1ap_ies

import (
	"fmt"

	"asn1go/per"
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
func (s *SystemBearerContextModificationResponse) Encode(w *per.Encoder) (err error) {

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			per.AlternativeInfo{Name: "DRB-Setup-Mod-List-EUTRAN", Tag: 0},
			per.AlternativeInfo{Name: "DRB-Failed-Mod-List-EUTRAN", Tag: 0},
			per.AlternativeInfo{Name: "DRB-Modified-List-EUTRAN", Tag: 0},
			per.AlternativeInfo{Name: "DRB-Failed-To-Modify-List-EUTRAN", Tag: 0},
			per.AlternativeInfo{Name: "RetainabilityMeasurementsInfo", Tag: 0},
		},
	}
	choiceEncoder := w.NewChoiceEncoder(c)

	switch s.Choice {
	case SystemBearerContextModificationResponsePresentDRBSetupModListEUTRAN:
		if err = choiceEncoder.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err = s.DRBSetupModListEUTRAN.Encode(w); err != nil {
			return fmt.Errorf("encode DRBSetupModListEUTRAN failed: %w", err)
		}
	case SystemBearerContextModificationResponsePresentDRBFailedModListEUTRAN:
		if err = choiceEncoder.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err = s.DRBFailedModListEUTRAN.Encode(w); err != nil {
			return fmt.Errorf("encode DRBFailedModListEUTRAN failed: %w", err)
		}
	case SystemBearerContextModificationResponsePresentDRBModifiedListEUTRAN:
		if err = choiceEncoder.EncodeChoice(2, false, nil); err != nil {
			return err
		}
		if err = s.DRBModifiedListEUTRAN.Encode(w); err != nil {
			return fmt.Errorf("encode DRBModifiedListEUTRAN failed: %w", err)
		}
	case SystemBearerContextModificationResponsePresentDRBFailedToModifyListEUTRAN:
		if err = choiceEncoder.EncodeChoice(3, false, nil); err != nil {
			return err
		}
		if err = s.DRBFailedToModifyListEUTRAN.Encode(w); err != nil {
			return fmt.Errorf("encode DRBFailedToModifyListEUTRAN failed: %w", err)
		}
	case SystemBearerContextModificationResponsePresentRetainabilityMeasurementsInfo:
		if err = choiceEncoder.EncodeChoice(4, false, nil); err != nil {
			return err
		}
		if err = s.RetainabilityMeasurementsInfo.Encode(w); err != nil {
			return fmt.Errorf("encode RetainabilityMeasurementsInfo failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of SystemBearerContextModificationResponse with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *SystemBearerContextModificationResponse) Decode(r *per.Decoder) (err error) {

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			per.AlternativeInfo{Name: "DRB-Setup-Mod-List-EUTRAN", Tag: 0},
			per.AlternativeInfo{Name: "DRB-Failed-Mod-List-EUTRAN", Tag: 0},
			per.AlternativeInfo{Name: "DRB-Modified-List-EUTRAN", Tag: 0},
			per.AlternativeInfo{Name: "DRB-Failed-To-Modify-List-EUTRAN", Tag: 0},
			per.AlternativeInfo{Name: "RetainabilityMeasurementsInfo", Tag: 0},
		},
	}
	choiceDecoder := r.NewChoiceDecoder(c)

	choiceIndex, isExtension, _, err := choiceDecoder.DecodeChoice()
	if err != nil {
		return fmt.Errorf("decode choice index failed: %w", err)
	}

	if isExtension {
		return fmt.Errorf("extension choices are not fully supported yet")
	}

	s.Choice = uint64(choiceIndex + 1) // 1-based internal Choice enum

	switch choiceIndex {
	case 0:
		s.DRBSetupModListEUTRAN = new(DRBSetupModListEUTRAN)
		if err = s.DRBSetupModListEUTRAN.Decode(r); err != nil {
			return fmt.Errorf("decode DRBSetupModListEUTRAN failed: %w", err)
		}
	case 1:
		s.DRBFailedModListEUTRAN = new(DRBFailedModListEUTRAN)
		if err = s.DRBFailedModListEUTRAN.Decode(r); err != nil {
			return fmt.Errorf("decode DRBFailedModListEUTRAN failed: %w", err)
		}
	case 2:
		s.DRBModifiedListEUTRAN = new(DRBModifiedListEUTRAN)
		if err = s.DRBModifiedListEUTRAN.Decode(r); err != nil {
			return fmt.Errorf("decode DRBModifiedListEUTRAN failed: %w", err)
		}
	case 3:
		s.DRBFailedToModifyListEUTRAN = new(DRBFailedToModifyListEUTRAN)
		if err = s.DRBFailedToModifyListEUTRAN.Decode(r); err != nil {
			return fmt.Errorf("decode DRBFailedToModifyListEUTRAN failed: %w", err)
		}
	case 4:
		s.RetainabilityMeasurementsInfo = new(RetainabilityMeasurementsInfo)
		if err = s.RetainabilityMeasurementsInfo.Decode(r); err != nil {
			return fmt.Errorf("decode RetainabilityMeasurementsInfo failed: %w", err)
		}
	default:
		return fmt.Errorf("decode choice of SystemBearerContextModificationResponse with unknown choice index %d", choiceIndex)
	}
	return nil
}
