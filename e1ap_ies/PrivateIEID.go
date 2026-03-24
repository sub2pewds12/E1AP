package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// PrivateIEID is a generated CHOICE type.
type PrivateIEID struct {
	Choice uint64 `json:"-"`
	Local  *INTEGER
	Global *string
}

const (
	PrivateIEIDPresentNothing uint64 = iota
	PrivateIEIDPresentLocal
	PrivateIEIDPresentGlobal
)

// Encode implements the aper.AperMarshaller interface.
func (s *PrivateIEID) Encode(w *per.Encoder) (err error) {

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			per.AlternativeInfo{Name: "local", Tag: 0},
			per.AlternativeInfo{Name: "global", Tag: 0},
		},
	}
	choiceEncoder := w.NewChoiceEncoder(c)

	switch s.Choice {
	case PrivateIEIDPresentLocal:
		if err = choiceEncoder.EncodeChoice(0, false, nil); err != nil {
			return err
		}
		if err = s.Local.Encode(w); err != nil {
			return fmt.Errorf("encode Local failed: %w", err)
		}
	case PrivateIEIDPresentGlobal:
		if err = choiceEncoder.EncodeChoice(1, false, nil); err != nil {
			return err
		}
		if err = w.EncodeOctetString([]byte(*s.Global), per.SizeConstraints{Extensible: false, Min: nil, Max: nil}); err != nil {
			return fmt.Errorf("encode Global failed: %w", err)
		}
	default:
		return fmt.Errorf("Encode choice of PrivateIEID with unknown choice value %d", s.Choice)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *PrivateIEID) Decode(r *per.Decoder) (err error) {

	c := per.ChoiceConstraints{
		Extensible: false,
		RootAlternatives: []per.AlternativeInfo{
			per.AlternativeInfo{Name: "local", Tag: 0},
			per.AlternativeInfo{Name: "global", Tag: 0},
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
		s.Local = new(INTEGER)
		if err = s.Local.Decode(r); err != nil {
			return fmt.Errorf("decode Local failed: %w", err)
		}
	case 1:
		val, err := r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: nil, Max: nil})
		if err != nil {
			return fmt.Errorf("decode Global failed: %w", err)
		}
		tmpStr := string(val)
		s.Global = &tmpStr
	default:
		return fmt.Errorf("decode choice of PrivateIEID with unknown choice index %d", choiceIndex)
	}
	return nil
}
