package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// GBRQosInformation is a generated SEQUENCE type.
type GBRQosInformation struct {
	ERABMaximumBitrateDL    *BitRate
	ERABMaximumBitrateUL    *BitRate
	ERABGuaranteedBitrateDL *BitRate
	ERABGuaranteedBitrateUL *BitRate
	IEExtensions            *GBRQosInformationExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *GBRQosInformation) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "e-RAB-MaximumBitrateDL", Optional: true},
			per.ComponentInfo{Name: "e-RAB-MaximumBitrateUL", Optional: true},
			per.ComponentInfo{Name: "e-RAB-GuaranteedBitrateDL", Optional: true},
			per.ComponentInfo{Name: "e-RAB-GuaranteedBitrateUL", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.ERABMaximumBitrateDL != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.ERABMaximumBitrateUL != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.ERABGuaranteedBitrateDL != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.ERABGuaranteedBitrateUL != nil {
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

	if s.ERABMaximumBitrateDL != nil {
		if err = w.EncodeInteger(int64((*s.ERABMaximumBitrateDL).Value), per.ConstrainedExtensible(0, 4000000000000)); err != nil {
			return fmt.Errorf("encode ERABMaximumBitrateDL failed: %w", err)
		}
	}

	if s.ERABMaximumBitrateUL != nil {
		if err = w.EncodeInteger(int64((*s.ERABMaximumBitrateUL).Value), per.ConstrainedExtensible(0, 4000000000000)); err != nil {
			return fmt.Errorf("encode ERABMaximumBitrateUL failed: %w", err)
		}
	}

	if s.ERABGuaranteedBitrateDL != nil {
		if err = w.EncodeInteger(int64((*s.ERABGuaranteedBitrateDL).Value), per.ConstrainedExtensible(0, 4000000000000)); err != nil {
			return fmt.Errorf("encode ERABGuaranteedBitrateDL failed: %w", err)
		}
	}

	if s.ERABGuaranteedBitrateUL != nil {
		if err = w.EncodeInteger(int64((*s.ERABGuaranteedBitrateUL).Value), per.ConstrainedExtensible(0, 4000000000000)); err != nil {
			return fmt.Errorf("encode ERABGuaranteedBitrateUL failed: %w", err)
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
func (s *GBRQosInformation) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "e-RAB-MaximumBitrateDL", Optional: true},
			per.ComponentInfo{Name: "e-RAB-MaximumBitrateUL", Optional: true},
			per.ComponentInfo{Name: "e-RAB-GuaranteedBitrateDL", Optional: true},
			per.ComponentInfo{Name: "e-RAB-GuaranteedBitrateUL", Optional: true},
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
		s.ERABMaximumBitrateDL = new(BitRate)

		{
			val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 4000000000000))
			if err != nil {
				return fmt.Errorf("decode ERABMaximumBitrateDL failed: %w", err)
			}
			s.ERABMaximumBitrateDL.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(1) {
		s.ERABMaximumBitrateUL = new(BitRate)

		{
			val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 4000000000000))
			if err != nil {
				return fmt.Errorf("decode ERABMaximumBitrateUL failed: %w", err)
			}
			s.ERABMaximumBitrateUL.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(2) {
		s.ERABGuaranteedBitrateDL = new(BitRate)

		{
			val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 4000000000000))
			if err != nil {
				return fmt.Errorf("decode ERABGuaranteedBitrateDL failed: %w", err)
			}
			s.ERABGuaranteedBitrateDL.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(3) {
		s.ERABGuaranteedBitrateUL = new(BitRate)

		{
			val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 4000000000000))
			if err != nil {
				return fmt.Errorf("decode ERABGuaranteedBitrateUL failed: %w", err)
			}
			s.ERABGuaranteedBitrateUL.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(4) {
		s.IEExtensions = new(GBRQosInformationExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
