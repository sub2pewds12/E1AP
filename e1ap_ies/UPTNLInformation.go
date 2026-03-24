package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// UPTNLInformation is a generated CHOICE type.
type UPTNLInformation struct {
	Choice          uint64 `json:"-"`
	GTPTunnel       *GTPTunnel
	ChoiceExtension *ProtocolIESingleContainer
}

const (
	UPTNLInformationPresentNothing uint64 = iota
	UPTNLInformationPresentGTPTunnel
	UPTNLInformationPresentChoiceExtension
)

// Encode implements the aper.AperMarshaller interface.
func (s *UPTNLInformation) Encode(w *per.Encoder) (err error) {

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			per.AlternativeInfo{Name: "gTPTunnel", Tag: 0},
			per.AlternativeInfo{Name: "choice-extension", Tag: 0},
		},
	}
	choiceEncoder := w.NewChoiceEncoder(c)

	switch s.Choice {
	case UPTNLInformationPresentGTPTunnel:
		if err = choiceEncoder.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err = s.GTPTunnel.Encode(w); err != nil {
			return fmt.Errorf("encode GTPTunnel failed: %w", err)
		}
	case UPTNLInformationPresentChoiceExtension:
		if err = choiceEncoder.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err = s.ChoiceExtension.Encode(w); err != nil {
			return fmt.Errorf("encode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of UPTNLInformation with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *UPTNLInformation) Decode(r *per.Decoder) (err error) {

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			per.AlternativeInfo{Name: "gTPTunnel", Tag: 0},
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
		s.GTPTunnel = new(GTPTunnel)
		if err = s.GTPTunnel.Decode(r); err != nil {
			return fmt.Errorf("decode GTPTunnel failed: %w", err)
		}
	case 1:
		s.ChoiceExtension = new(ProtocolIESingleContainer)
		if err = s.ChoiceExtension.Decode(r); err != nil {
			return fmt.Errorf("decode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("decode choice of UPTNLInformation with unknown choice index %d", choiceIndex)
	}
	return nil
}
