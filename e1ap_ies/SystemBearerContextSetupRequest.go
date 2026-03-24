package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// SystemBearerContextSetupRequest is a generated CHOICE type.
type SystemBearerContextSetupRequest struct {
	Choice                     uint64 `json:"-"`
	DRBToSetupListEUTRAN       *DRBToSetupListEUTRAN
	SubscriberProfileIDforRFP  *SubscriberProfileIDforRFP
	AdditionalRRMPriorityIndex *AdditionalRRMPriorityIndex
}

const (
	SystemBearerContextSetupRequestPresentNothing uint64 = iota
	SystemBearerContextSetupRequestPresentDRBToSetupListEUTRAN
	SystemBearerContextSetupRequestPresentSubscriberProfileIDforRFP
	SystemBearerContextSetupRequestPresentAdditionalRRMPriorityIndex
)

// Encode implements the aper.AperMarshaller interface.
func (s *SystemBearerContextSetupRequest) Encode(w *per.Encoder) (err error) {

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			per.AlternativeInfo{Name: "DRB-To-Setup-List-EUTRAN", Tag: 0},
			per.AlternativeInfo{Name: "SubscriberProfileIDforRFP", Tag: 0},
			per.AlternativeInfo{Name: "AdditionalRRMPriorityIndex", Tag: 0},
		},
	}
	choiceEncoder := w.NewChoiceEncoder(c)

	switch s.Choice {
	case SystemBearerContextSetupRequestPresentDRBToSetupListEUTRAN:
		if err = choiceEncoder.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err = s.DRBToSetupListEUTRAN.Encode(w); err != nil {
			return fmt.Errorf("encode DRBToSetupListEUTRAN failed: %w", err)
		}
	case SystemBearerContextSetupRequestPresentSubscriberProfileIDforRFP:
		if err = choiceEncoder.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err = w.EncodeInteger(int64(s.SubscriberProfileIDforRFP.Value), per.ConstrainedExtensible(1, 256)); err != nil {
			return fmt.Errorf("encode SubscriberProfileIDforRFP failed: %w", err)
		}
	case SystemBearerContextSetupRequestPresentAdditionalRRMPriorityIndex:
		if err = choiceEncoder.EncodeChoice(2, false, nil); err != nil {
			return err
		}
		if err = s.AdditionalRRMPriorityIndex.Encode(w); err != nil {
			return fmt.Errorf("encode AdditionalRRMPriorityIndex failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of SystemBearerContextSetupRequest with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *SystemBearerContextSetupRequest) Decode(r *per.Decoder) (err error) {

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			per.AlternativeInfo{Name: "DRB-To-Setup-List-EUTRAN", Tag: 0},
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
		s.DRBToSetupListEUTRAN = new(DRBToSetupListEUTRAN)
		if err = s.DRBToSetupListEUTRAN.Decode(r); err != nil {
			return fmt.Errorf("decode DRBToSetupListEUTRAN failed: %w", err)
		}
	case 1:
		s.SubscriberProfileIDforRFP = new(SubscriberProfileIDforRFP)
		val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 256))
		if err != nil {
			return fmt.Errorf("decode SubscriberProfileIDforRFP failed: %w", err)
		}
		s.SubscriberProfileIDforRFP.Value = val
	case 2:
		s.AdditionalRRMPriorityIndex = new(AdditionalRRMPriorityIndex)
		if err = s.AdditionalRRMPriorityIndex.Decode(r); err != nil {
			return fmt.Errorf("decode AdditionalRRMPriorityIndex failed: %w", err)
		}
	default:
		return fmt.Errorf("decode choice of SystemBearerContextSetupRequest with unknown choice index %d", choiceIndex)
	}
	return nil
}
