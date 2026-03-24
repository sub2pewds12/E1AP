package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// Cause is a generated CHOICE type.
type Cause struct {
	Choice          uint64 `json:"-"`
	RadioNetwork    *CauseRadioNetwork
	Transport       *CauseTransport
	Protocol        *CauseProtocol
	Misc            *CauseMisc
	ChoiceExtension *ProtocolIESingleContainer
}

const (
	CausePresentNothing uint64 = iota
	CausePresentRadioNetwork
	CausePresentTransport
	CausePresentProtocol
	CausePresentMisc
	CausePresentChoiceExtension
)

// Encode implements the aper.AperMarshaller interface.
func (s *Cause) Encode(w *per.Encoder) (err error) {

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			per.AlternativeInfo{Name: "radioNetwork", Tag: 0},
			per.AlternativeInfo{Name: "transport", Tag: 0},
			per.AlternativeInfo{Name: "protocol", Tag: 0},
			per.AlternativeInfo{Name: "misc", Tag: 0},
			per.AlternativeInfo{Name: "choice-extension", Tag: 0},
		},
	}
	choiceEncoder := w.NewChoiceEncoder(c)

	switch s.Choice {
	case CausePresentRadioNetwork:
		if err = choiceEncoder.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err = s.RadioNetwork.Encode(w); err != nil {
			return fmt.Errorf("encode RadioNetwork failed: %w", err)
		}
	case CausePresentTransport:
		if err = choiceEncoder.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err = s.Transport.Encode(w); err != nil {
			return fmt.Errorf("encode Transport failed: %w", err)
		}
	case CausePresentProtocol:
		if err = choiceEncoder.EncodeChoice(2, false, nil); err != nil {
			return err
		}
		if err = s.Protocol.Encode(w); err != nil {
			return fmt.Errorf("encode Protocol failed: %w", err)
		}
	case CausePresentMisc:
		if err = choiceEncoder.EncodeChoice(3, false, nil); err != nil {
			return err
		}
		if err = s.Misc.Encode(w); err != nil {
			return fmt.Errorf("encode Misc failed: %w", err)
		}
	case CausePresentChoiceExtension:
		if err = choiceEncoder.EncodeChoice(4, false, nil); err != nil {
			return err
		}
		if err = s.ChoiceExtension.Encode(w); err != nil {
			return fmt.Errorf("encode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of Cause with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *Cause) Decode(r *per.Decoder) (err error) {

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			per.AlternativeInfo{Name: "radioNetwork", Tag: 0},
			per.AlternativeInfo{Name: "transport", Tag: 0},
			per.AlternativeInfo{Name: "protocol", Tag: 0},
			per.AlternativeInfo{Name: "misc", Tag: 0},
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
		s.RadioNetwork = new(CauseRadioNetwork)
		if err = s.RadioNetwork.Decode(r); err != nil {
			return fmt.Errorf("decode RadioNetwork failed: %w", err)
		}
	case 1:
		s.Transport = new(CauseTransport)
		if err = s.Transport.Decode(r); err != nil {
			return fmt.Errorf("decode Transport failed: %w", err)
		}
	case 2:
		s.Protocol = new(CauseProtocol)
		if err = s.Protocol.Decode(r); err != nil {
			return fmt.Errorf("decode Protocol failed: %w", err)
		}
	case 3:
		s.Misc = new(CauseMisc)
		if err = s.Misc.Decode(r); err != nil {
			return fmt.Errorf("decode Misc failed: %w", err)
		}
	case 4:
		s.ChoiceExtension = new(ProtocolIESingleContainer)
		if err = s.ChoiceExtension.Decode(r); err != nil {
			return fmt.Errorf("decode ChoiceExtension failed: %w", err)
		}
	default:
		return fmt.Errorf("decode choice of Cause with unknown choice index %d", choiceIndex)
	}
	return nil
}
