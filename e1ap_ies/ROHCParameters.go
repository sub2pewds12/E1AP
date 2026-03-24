package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// ROHCParameters is a generated CHOICE type.
type ROHCParameters struct {
	Choice          uint64 `json:"-"`
	ROHC            *ROHC
	UPlinkOnlyROHC  *UplinkOnlyROHC
	ChoiceExtension *ProtocolIESingleContainer
}

const (
	ROHCParametersPresentNothing uint64 = iota
	ROHCParametersPresentROHC
	ROHCParametersPresentUPlinkOnlyROHC
	ROHCParametersPresentChoiceExtension
)

// Encode implements the aper.AperMarshaller interface.
func (s *ROHCParameters) Encode(w *per.Encoder) (err error) {

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			per.AlternativeInfo{Name: "rOHC", Tag: 0},
			per.AlternativeInfo{Name: "uPlinkOnlyROHC", Tag: 0},
			per.AlternativeInfo{Name: "choice-Extension", Tag: 0},
		},
	}
	choiceEncoder := w.NewChoiceEncoder(c)

	switch s.Choice {
	case ROHCParametersPresentROHC:
		if err = choiceEncoder.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err = s.ROHC.Encode(w); err != nil {
			return fmt.Errorf("encode ROHC failed: %w", err)
		}
	case ROHCParametersPresentUPlinkOnlyROHC:
		if err = choiceEncoder.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err = s.UPlinkOnlyROHC.Encode(w); err != nil {
			return fmt.Errorf("encode UPlinkOnlyROHC failed: %w", err)
		}
	case ROHCParametersPresentChoiceExtension:
		if err = choiceEncoder.EncodeChoice(2, false, nil); err != nil {
			return err
		}
		if err = s.ChoiceExtension.Encode(w); err != nil {
			return fmt.Errorf("encode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of ROHCParameters with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ROHCParameters) Decode(r *per.Decoder) (err error) {

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			per.AlternativeInfo{Name: "rOHC", Tag: 0},
			per.AlternativeInfo{Name: "uPlinkOnlyROHC", Tag: 0},
			per.AlternativeInfo{Name: "choice-Extension", Tag: 0},
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
		s.ROHC = new(ROHC)
		if err = s.ROHC.Decode(r); err != nil {
			return fmt.Errorf("decode ROHC failed: %w", err)
		}
	case 1:
		s.UPlinkOnlyROHC = new(UplinkOnlyROHC)
		if err = s.UPlinkOnlyROHC.Decode(r); err != nil {
			return fmt.Errorf("decode UPlinkOnlyROHC failed: %w", err)
		}
	case 2:
		s.ChoiceExtension = new(ProtocolIESingleContainer)
		if err = s.ChoiceExtension.Decode(r); err != nil {
			return fmt.Errorf("decode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("decode choice of ROHCParameters with unknown choice index %d", choiceIndex)
	}
	return nil
}
