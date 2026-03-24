package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
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
func (s *SystemBearerContextModificationRequest) Encode(w *per.Encoder) (err error) {

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			per.AlternativeInfo{Name: "DRB-To-Setup-Mod-List-EUTRAN", Tag: 0},
			per.AlternativeInfo{Name: "DRB-To-Modify-List-EUTRAN", Tag: 0},
			per.AlternativeInfo{Name: "DRB-To-Remove-List-EUTRAN", Tag: 0},
			per.AlternativeInfo{Name: "SubscriberProfileIDforRFP", Tag: 0},
			per.AlternativeInfo{Name: "AdditionalRRMPriorityIndex", Tag: 0},
		},
	}
	choiceEncoder := w.NewChoiceEncoder(c)

	switch s.Choice {
	case SystemBearerContextModificationRequestPresentDRBToSetupModListEUTRAN:
		if err = choiceEncoder.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err = s.DRBToSetupModListEUTRAN.Encode(w); err != nil {
			return fmt.Errorf("encode DRBToSetupModListEUTRAN failed: %w", err)
		}
	case SystemBearerContextModificationRequestPresentDRBToModifyListEUTRAN:
		if err = choiceEncoder.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err = s.DRBToModifyListEUTRAN.Encode(w); err != nil {
			return fmt.Errorf("encode DRBToModifyListEUTRAN failed: %w", err)
		}
	case SystemBearerContextModificationRequestPresentDRBToRemoveListEUTRAN:
		if err = choiceEncoder.EncodeChoice(2, false, nil); err != nil {
			return err
		}
		if err = s.DRBToRemoveListEUTRAN.Encode(w); err != nil {
			return fmt.Errorf("encode DRBToRemoveListEUTRAN failed: %w", err)
		}
	case SystemBearerContextModificationRequestPresentSubscriberProfileIDforRFP:
		if err = choiceEncoder.EncodeChoice(3, false, nil); err != nil {
			return err
		}
		if err = w.EncodeInteger(int64(s.SubscriberProfileIDforRFP.Value), per.ConstrainedExtensible(1, 256)); err != nil {
			return fmt.Errorf("encode SubscriberProfileIDforRFP failed: %w", err)
		}
	case SystemBearerContextModificationRequestPresentAdditionalRRMPriorityIndex:
		if err = choiceEncoder.EncodeChoice(4, false, nil); err != nil {
			return err
		}
		if err = s.AdditionalRRMPriorityIndex.Encode(w); err != nil {
			return fmt.Errorf("encode AdditionalRRMPriorityIndex failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of SystemBearerContextModificationRequest with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *SystemBearerContextModificationRequest) Decode(r *per.Decoder) (err error) {

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			per.AlternativeInfo{Name: "DRB-To-Setup-Mod-List-EUTRAN", Tag: 0},
			per.AlternativeInfo{Name: "DRB-To-Modify-List-EUTRAN", Tag: 0},
			per.AlternativeInfo{Name: "DRB-To-Remove-List-EUTRAN", Tag: 0},
			per.AlternativeInfo{Name: "SubscriberProfileIDforRFP", Tag: 0},
			per.AlternativeInfo{Name: "AdditionalRRMPriorityIndex", Tag: 0},
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
		s.DRBToSetupModListEUTRAN = new(DRBToSetupModListEUTRAN)
		if err = s.DRBToSetupModListEUTRAN.Decode(r); err != nil {
			return fmt.Errorf("decode DRBToSetupModListEUTRAN failed: %w", err)
		}
	case 1:
		s.DRBToModifyListEUTRAN = new(DRBToModifyListEUTRAN)
		if err = s.DRBToModifyListEUTRAN.Decode(r); err != nil {
			return fmt.Errorf("decode DRBToModifyListEUTRAN failed: %w", err)
		}
	case 2:
		s.DRBToRemoveListEUTRAN = new(DRBToRemoveListEUTRAN)
		if err = s.DRBToRemoveListEUTRAN.Decode(r); err != nil {
			return fmt.Errorf("decode DRBToRemoveListEUTRAN failed: %w", err)
		}
	case 3:
		s.SubscriberProfileIDforRFP = new(SubscriberProfileIDforRFP)
		val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 256))
		if err != nil {
			return fmt.Errorf("decode SubscriberProfileIDforRFP failed: %w", err)
		}
		s.SubscriberProfileIDforRFP.Value = val
	case 4:
		s.AdditionalRRMPriorityIndex = new(AdditionalRRMPriorityIndex)
		if err = s.AdditionalRRMPriorityIndex.Decode(r); err != nil {
			return fmt.Errorf("decode AdditionalRRMPriorityIndex failed: %w", err)
		}
	default:
		return fmt.Errorf("decode choice of SystemBearerContextModificationRequest with unknown choice index %d", choiceIndex)
	}
	return nil
}
