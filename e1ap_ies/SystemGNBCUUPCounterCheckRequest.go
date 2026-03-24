package e1ap_ies

import (
	"fmt"

	"asn1go/per"
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
func (s *SystemGNBCUUPCounterCheckRequest) Encode(w *per.Encoder) (err error) {

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			per.AlternativeInfo{Name: "DRBs-Subject-To-Counter-Check-List-EUTRAN", Tag: 0},
		},
	}
	choiceEncoder := w.NewChoiceEncoder(c)

	switch s.Choice {
	case SystemGNBCUUPCounterCheckRequestPresentDRBsSubjectToCounterCheckListEUTRAN:
		if err = choiceEncoder.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err = s.DRBsSubjectToCounterCheckListEUTRAN.Encode(w); err != nil {
			return fmt.Errorf("encode DRBsSubjectToCounterCheckListEUTRAN failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of SystemGNBCUUPCounterCheckRequest with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *SystemGNBCUUPCounterCheckRequest) Decode(r *per.Decoder) (err error) {

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			per.AlternativeInfo{Name: "DRBs-Subject-To-Counter-Check-List-EUTRAN", Tag: 0},
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
		s.DRBsSubjectToCounterCheckListEUTRAN = new(DRBsSubjectToCounterCheckListEUTRAN)
		if err = s.DRBsSubjectToCounterCheckListEUTRAN.Decode(r); err != nil {
			return fmt.Errorf("decode DRBsSubjectToCounterCheckListEUTRAN failed: %w", err)
		}
	default:
		return fmt.Errorf("decode choice of SystemGNBCUUPCounterCheckRequest with unknown choice index %d", choiceIndex)
	}
	return nil
}
