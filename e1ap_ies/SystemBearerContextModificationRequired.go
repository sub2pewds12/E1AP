package e1ap_ies

import (
	"fmt"

	"asn1go/per"
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
func (s *SystemBearerContextModificationRequired) Encode(w *per.Encoder) (err error) {

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			per.AlternativeInfo{Name: "DRB-Required-To-Modify-List-EUTRAN", Tag: 0},
			per.AlternativeInfo{Name: "DRB-Required-To-Remove-List-EUTRAN", Tag: 0},
		},
	}
	choiceEncoder := w.NewChoiceEncoder(c)

	switch s.Choice {
	case SystemBearerContextModificationRequiredPresentDRBRequiredToModifyListEUTRAN:
		if err = choiceEncoder.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err = s.DRBRequiredToModifyListEUTRAN.Encode(w); err != nil {
			return fmt.Errorf("encode DRBRequiredToModifyListEUTRAN failed: %w", err)
		}
	case SystemBearerContextModificationRequiredPresentDRBRequiredToRemoveListEUTRAN:
		if err = choiceEncoder.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err = s.DRBRequiredToRemoveListEUTRAN.Encode(w); err != nil {
			return fmt.Errorf("encode DRBRequiredToRemoveListEUTRAN failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of SystemBearerContextModificationRequired with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *SystemBearerContextModificationRequired) Decode(r *per.Decoder) (err error) {

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			per.AlternativeInfo{Name: "DRB-Required-To-Modify-List-EUTRAN", Tag: 0},
			per.AlternativeInfo{Name: "DRB-Required-To-Remove-List-EUTRAN", Tag: 0},
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
		s.DRBRequiredToModifyListEUTRAN = new(DRBRequiredToModifyListEUTRAN)
		if err = s.DRBRequiredToModifyListEUTRAN.Decode(r); err != nil {
			return fmt.Errorf("decode DRBRequiredToModifyListEUTRAN failed: %w", err)
		}
	case 1:
		s.DRBRequiredToRemoveListEUTRAN = new(DRBRequiredToRemoveListEUTRAN)
		if err = s.DRBRequiredToRemoveListEUTRAN.Decode(r); err != nil {
			return fmt.Errorf("decode DRBRequiredToRemoveListEUTRAN failed: %w", err)
		}
	default:
		return fmt.Errorf("decode choice of SystemBearerContextModificationRequired with unknown choice index %d", choiceIndex)
	}
	return nil
}
