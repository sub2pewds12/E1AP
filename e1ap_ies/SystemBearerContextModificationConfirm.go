package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
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
func (s *SystemBearerContextModificationConfirm) Encode(w *per.Encoder) (err error) {

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			per.AlternativeInfo{Name: "DRB-Confirm-Modified-List-EUTRAN", Tag: 0},
		},
	}
	choiceEncoder := w.NewChoiceEncoder(c)

	switch s.Choice {
	case SystemBearerContextModificationConfirmPresentDRBConfirmModifiedListEUTRAN:
		if err = choiceEncoder.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err = s.DRBConfirmModifiedListEUTRAN.Encode(w); err != nil {
			return fmt.Errorf("encode DRBConfirmModifiedListEUTRAN failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of SystemBearerContextModificationConfirm with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *SystemBearerContextModificationConfirm) Decode(r *per.Decoder) (err error) {

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			per.AlternativeInfo{Name: "DRB-Confirm-Modified-List-EUTRAN", Tag: 0},
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
		s.DRBConfirmModifiedListEUTRAN = new(DRBConfirmModifiedListEUTRAN)
		if err = s.DRBConfirmModifiedListEUTRAN.Decode(r); err != nil {
			return fmt.Errorf("decode DRBConfirmModifiedListEUTRAN failed: %w", err)
		}
	default:
		return fmt.Errorf("decode choice of SystemBearerContextModificationConfirm with unknown choice index %d", choiceIndex)
	}
	return nil
}
