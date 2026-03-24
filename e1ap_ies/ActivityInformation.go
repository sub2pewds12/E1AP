package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// ActivityInformation is a generated CHOICE type.
type ActivityInformation struct {
	Choice                         uint64 `json:"-"`
	DRBActivityList                *DRBActivityList
	PDUSessionResourceActivityList *PDUSessionResourceActivityList
	UEActivity                     *UEActivity
	ChoiceExtension                *ProtocolIESingleContainer
}

const (
	ActivityInformationPresentNothing uint64 = iota
	ActivityInformationPresentDRBActivityList
	ActivityInformationPresentPDUSessionResourceActivityList
	ActivityInformationPresentUEActivity
	ActivityInformationPresentChoiceExtension
)

// Encode implements the aper.AperMarshaller interface.
func (s *ActivityInformation) Encode(w *per.Encoder) (err error) {

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			per.AlternativeInfo{Name: "dRB-Activity-List", Tag: 0},
			per.AlternativeInfo{Name: "pDU-Session-Resource-Activity-List", Tag: 0},
			per.AlternativeInfo{Name: "uE-Activity", Tag: 0},
			per.AlternativeInfo{Name: "choice-extension", Tag: 0},
		},
	}
	choiceEncoder := w.NewChoiceEncoder(c)

	switch s.Choice {
	case ActivityInformationPresentDRBActivityList:
		if err = choiceEncoder.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err = s.DRBActivityList.Encode(w); err != nil {
			return fmt.Errorf("encode DRBActivityList failed: %w", err)
		}
	case ActivityInformationPresentPDUSessionResourceActivityList:
		if err = choiceEncoder.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err = s.PDUSessionResourceActivityList.Encode(w); err != nil {
			return fmt.Errorf("encode PDUSessionResourceActivityList failed: %w", err)
		}
	case ActivityInformationPresentUEActivity:
		if err = choiceEncoder.EncodeChoice(2, false, nil); err != nil {
			return err
		}
		if err = s.UEActivity.Encode(w); err != nil {
			return fmt.Errorf("encode UEActivity failed: %w", err)
		}
	case ActivityInformationPresentChoiceExtension:
		if err = choiceEncoder.EncodeChoice(3, false, nil); err != nil {
			return err
		}
		if err = s.ChoiceExtension.Encode(w); err != nil {
			return fmt.Errorf("encode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of ActivityInformation with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ActivityInformation) Decode(r *per.Decoder) (err error) {

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			per.AlternativeInfo{Name: "dRB-Activity-List", Tag: 0},
			per.AlternativeInfo{Name: "pDU-Session-Resource-Activity-List", Tag: 0},
			per.AlternativeInfo{Name: "uE-Activity", Tag: 0},
			per.AlternativeInfo{Name: "choice-extension", Tag: 0},
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
		s.DRBActivityList = new(DRBActivityList)
		if err = s.DRBActivityList.Decode(r); err != nil {
			return fmt.Errorf("decode DRBActivityList failed: %w", err)
		}
	case 1:
		s.PDUSessionResourceActivityList = new(PDUSessionResourceActivityList)
		if err = s.PDUSessionResourceActivityList.Decode(r); err != nil {
			return fmt.Errorf("decode PDUSessionResourceActivityList failed: %w", err)
		}
	case 2:
		s.UEActivity = new(UEActivity)
		if err = s.UEActivity.Decode(r); err != nil {
			return fmt.Errorf("decode UEActivity failed: %w", err)
		}
	case 3:
		s.ChoiceExtension = new(ProtocolIESingleContainer)
		if err = s.ChoiceExtension.Decode(r); err != nil {
			return fmt.Errorf("decode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("decode choice of ActivityInformation with unknown choice index %d", choiceIndex)
	}
	return nil
}
