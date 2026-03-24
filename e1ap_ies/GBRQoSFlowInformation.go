package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// GBRQoSFlowInformation is a generated SEQUENCE type.
type GBRQoSFlowInformation struct {
	MaxFlowBitRateDownlink        *BitRate
	MaxFlowBitRateUplink          *BitRate
	GuaranteedFlowBitRateDownlink *BitRate
	GuaranteedFlowBitRateUplink   *BitRate
	MaxPacketLossRateDownlink     *MaxPacketLossRate
	MaxPacketLossRateUplink       *MaxPacketLossRate
	IEExtensions                  *GBRQoSFlowInformationExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *GBRQoSFlowInformation) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "maxFlowBitRateDownlink", Optional: true},
			per.ComponentInfo{Name: "maxFlowBitRateUplink", Optional: true},
			per.ComponentInfo{Name: "guaranteedFlowBitRateDownlink", Optional: true},
			per.ComponentInfo{Name: "guaranteedFlowBitRateUplink", Optional: true},
			per.ComponentInfo{Name: "maxPacketLossRateDownlink", Optional: true},
			per.ComponentInfo{Name: "maxPacketLossRateUplink", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.MaxFlowBitRateDownlink != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.MaxFlowBitRateUplink != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.GuaranteedFlowBitRateDownlink != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.GuaranteedFlowBitRateUplink != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.MaxPacketLossRateDownlink != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.MaxPacketLossRateUplink != nil {
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

	if s.MaxFlowBitRateDownlink != nil {
		if err = w.EncodeInteger(int64((*s.MaxFlowBitRateDownlink).Value), per.ConstrainedExtensible(0, 4000000000000)); err != nil {
			return fmt.Errorf("encode MaxFlowBitRateDownlink failed: %w", err)
		}
	}

	if s.MaxFlowBitRateUplink != nil {
		if err = w.EncodeInteger(int64((*s.MaxFlowBitRateUplink).Value), per.ConstrainedExtensible(0, 4000000000000)); err != nil {
			return fmt.Errorf("encode MaxFlowBitRateUplink failed: %w", err)
		}
	}

	if s.GuaranteedFlowBitRateDownlink != nil {
		if err = w.EncodeInteger(int64((*s.GuaranteedFlowBitRateDownlink).Value), per.ConstrainedExtensible(0, 4000000000000)); err != nil {
			return fmt.Errorf("encode GuaranteedFlowBitRateDownlink failed: %w", err)
		}
	}

	if s.GuaranteedFlowBitRateUplink != nil {
		if err = w.EncodeInteger(int64((*s.GuaranteedFlowBitRateUplink).Value), per.ConstrainedExtensible(0, 4000000000000)); err != nil {
			return fmt.Errorf("encode GuaranteedFlowBitRateUplink failed: %w", err)
		}
	}

	if s.MaxPacketLossRateDownlink != nil {
		if err = w.EncodeInteger(int64((*s.MaxPacketLossRateDownlink).Value), per.ConstrainedExtensible(0, 1000)); err != nil {
			return fmt.Errorf("encode MaxPacketLossRateDownlink failed: %w", err)
		}
	}

	if s.MaxPacketLossRateUplink != nil {
		if err = w.EncodeInteger(int64((*s.MaxPacketLossRateUplink).Value), per.ConstrainedExtensible(0, 1000)); err != nil {
			return fmt.Errorf("encode MaxPacketLossRateUplink failed: %w", err)
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
func (s *GBRQoSFlowInformation) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "maxFlowBitRateDownlink", Optional: true},
			per.ComponentInfo{Name: "maxFlowBitRateUplink", Optional: true},
			per.ComponentInfo{Name: "guaranteedFlowBitRateDownlink", Optional: true},
			per.ComponentInfo{Name: "guaranteedFlowBitRateUplink", Optional: true},
			per.ComponentInfo{Name: "maxPacketLossRateDownlink", Optional: true},
			per.ComponentInfo{Name: "maxPacketLossRateUplink", Optional: true},
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

	if seqDecoder.IsComponentPresent(0) {
		s.MaxFlowBitRateDownlink = new(BitRate)

		{
			val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 4000000000000))
			if err != nil {
				return fmt.Errorf("decode MaxFlowBitRateDownlink failed: %w", err)
			}
			s.MaxFlowBitRateDownlink.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(1) {
		s.MaxFlowBitRateUplink = new(BitRate)

		{
			val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 4000000000000))
			if err != nil {
				return fmt.Errorf("decode MaxFlowBitRateUplink failed: %w", err)
			}
			s.MaxFlowBitRateUplink.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(2) {
		s.GuaranteedFlowBitRateDownlink = new(BitRate)

		{
			val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 4000000000000))
			if err != nil {
				return fmt.Errorf("decode GuaranteedFlowBitRateDownlink failed: %w", err)
			}
			s.GuaranteedFlowBitRateDownlink.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(3) {
		s.GuaranteedFlowBitRateUplink = new(BitRate)

		{
			val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 4000000000000))
			if err != nil {
				return fmt.Errorf("decode GuaranteedFlowBitRateUplink failed: %w", err)
			}
			s.GuaranteedFlowBitRateUplink.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(4) {
		s.MaxPacketLossRateDownlink = new(MaxPacketLossRate)

		{
			val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 1000))
			if err != nil {
				return fmt.Errorf("decode MaxPacketLossRateDownlink failed: %w", err)
			}
			s.MaxPacketLossRateDownlink.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(5) {
		s.MaxPacketLossRateUplink = new(MaxPacketLossRate)

		{
			val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 1000))
			if err != nil {
				return fmt.Errorf("decode MaxPacketLossRateUplink failed: %w", err)
			}
			s.MaxPacketLossRateUplink.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(6) {
		s.IEExtensions = new(GBRQoSFlowInformationExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
