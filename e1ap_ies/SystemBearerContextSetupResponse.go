package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
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
func (s *SystemBearerContextSetupResponse) Encode(w *per.Encoder) (err error) {

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			per.AlternativeInfo{Name: "DRB-Setup-List-EUTRAN", Tag: 0},
			per.AlternativeInfo{Name: "DRB-Failed-List-EUTRAN", Tag: 0},
		},
	}
	choiceEncoder := w.NewChoiceEncoder(c)

	switch s.Choice {
	case SystemBearerContextSetupResponsePresentDRBSetupListEUTRAN:
		if err = choiceEncoder.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err = s.DRBSetupListEUTRAN.Encode(w); err != nil {
			return fmt.Errorf("encode DRBSetupListEUTRAN failed: %w", err)
		}
	case SystemBearerContextSetupResponsePresentDRBFailedListEUTRAN:
		if err = choiceEncoder.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err = s.DRBFailedListEUTRAN.Encode(w); err != nil {
			return fmt.Errorf("encode DRBFailedListEUTRAN failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of SystemBearerContextSetupResponse with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *SystemBearerContextSetupResponse) Decode(r *per.Decoder) (err error) {

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			per.AlternativeInfo{Name: "DRB-Setup-List-EUTRAN", Tag: 0},
			per.AlternativeInfo{Name: "DRB-Failed-List-EUTRAN", Tag: 0},
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
		s.DRBSetupListEUTRAN = new(DRBSetupListEUTRAN)
		if err = s.DRBSetupListEUTRAN.Decode(r); err != nil {
			return fmt.Errorf("decode DRBSetupListEUTRAN failed: %w", err)
		}
	case 1:
		s.DRBFailedListEUTRAN = new(DRBFailedListEUTRAN)
		if err = s.DRBFailedListEUTRAN.Decode(r); err != nil {
			return fmt.Errorf("decode DRBFailedListEUTRAN failed: %w", err)
		}
	default:
		return fmt.Errorf("decode choice of SystemBearerContextSetupResponse with unknown choice index %d", choiceIndex)
	}
	return nil
}
