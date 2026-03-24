package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// AlternativeQoSParaSetItem is a generated SEQUENCE type.
type AlternativeQoSParaSetItem struct {
	AlternativeQoSParameterIndex AlternativeQoSParaSetItemAlternativeQoSParameterIndex
	GuaranteedFlowBitRateDL      *BitRate
	GuaranteedFlowBitRateUL      *BitRate
	PacketDelayBudget            *PacketDelayBudget
	PacketErrorRate              *PacketErrorRate
	IEExtensions                 *AlternativeQoSParaSetItemExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *AlternativeQoSParaSetItem) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "alternativeQoSParameterIndex", Optional: false},
			per.ComponentInfo{Name: "guaranteedFlowBitRateDL", Optional: true},
			per.ComponentInfo{Name: "guaranteedFlowBitRateUL", Optional: true},
			per.ComponentInfo{Name: "packetDelayBudget", Optional: true},
			per.ComponentInfo{Name: "packetErrorRate", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.GuaranteedFlowBitRateDL != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.GuaranteedFlowBitRateUL != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.PacketDelayBudget != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.PacketErrorRate != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.IEExtensions != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if err := seqEncoder.EncodePreamble(optionalBitmap); err != nil {
		return err
	}

	if err = w.EncodeInteger(int64(s.AlternativeQoSParameterIndex.Value), per.ConstrainedExtensible(1, 8)); err != nil {
		return fmt.Errorf("encode AlternativeQoSParameterIndex failed: %w", err)
	}

	if s.GuaranteedFlowBitRateDL != nil {
		if err = w.EncodeInteger(int64((*s.GuaranteedFlowBitRateDL).Value), per.ConstrainedExtensible(0, 4000000000000)); err != nil {
			return fmt.Errorf("encode GuaranteedFlowBitRateDL failed: %w", err)
		}
	}

	if s.GuaranteedFlowBitRateUL != nil {
		if err = w.EncodeInteger(int64((*s.GuaranteedFlowBitRateUL).Value), per.ConstrainedExtensible(0, 4000000000000)); err != nil {
			return fmt.Errorf("encode GuaranteedFlowBitRateUL failed: %w", err)
		}
	}

	if s.PacketDelayBudget != nil {
		if err = w.EncodeInteger(int64((*s.PacketDelayBudget).Value), per.ConstrainedExtensible(0, 1023)); err != nil {
			return fmt.Errorf("encode PacketDelayBudget failed: %w", err)
		}
	}

	if s.PacketErrorRate != nil {
		if err = s.PacketErrorRate.Encode(w); err != nil {
			return fmt.Errorf("encode PacketErrorRate failed: %w", err)
		}
	}

	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("encode IEExtensions failed: %w", err)
		}
	}

	if err := seqEncoder.EncodeExtensionAdditions([]bool{}, [][]byte{}); err != nil {
		return err
	}

	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *AlternativeQoSParaSetItem) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "alternativeQoSParameterIndex", Optional: false},
			per.ComponentInfo{Name: "guaranteedFlowBitRateDL", Optional: true},
			per.ComponentInfo{Name: "guaranteedFlowBitRateUL", Optional: true},
			per.ComponentInfo{Name: "packetDelayBudget", Optional: true},
			per.ComponentInfo{Name: "packetErrorRate", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqDecoder := r.NewSequenceDecoder(c)
	if err := seqDecoder.DecodeExtensionBit(); err != nil {
		return err
	}

	if err := seqDecoder.DecodePreamble(); err != nil {
		return err
	}

	{
		val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 8))
		if err != nil {
			return fmt.Errorf("decode AlternativeQoSParameterIndex failed: %w", err)
		}
		s.AlternativeQoSParameterIndex.Value = val
	}

	if seqDecoder.IsComponentPresent(1) {
		s.GuaranteedFlowBitRateDL = new(BitRate)

		{
			val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 4000000000000))
			if err != nil {
				return fmt.Errorf("decode GuaranteedFlowBitRateDL failed: %w", err)
			}
			s.GuaranteedFlowBitRateDL.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(2) {
		s.GuaranteedFlowBitRateUL = new(BitRate)

		{
			val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 4000000000000))
			if err != nil {
				return fmt.Errorf("decode GuaranteedFlowBitRateUL failed: %w", err)
			}
			s.GuaranteedFlowBitRateUL.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(3) {
		s.PacketDelayBudget = new(PacketDelayBudget)

		{
			val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 1023))
			if err != nil {
				return fmt.Errorf("decode PacketDelayBudget failed: %w", err)
			}
			s.PacketDelayBudget.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(4) {
		s.PacketErrorRate = new(PacketErrorRate)
		if err = s.PacketErrorRate.Decode(r); err != nil {
			return fmt.Errorf("Decode PacketErrorRate failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(5) {
		s.IEExtensions = new(AlternativeQoSParaSetItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
