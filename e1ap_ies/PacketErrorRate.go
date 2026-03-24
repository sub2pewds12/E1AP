package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// PacketErrorRate is a generated SEQUENCE type.
type PacketErrorRate struct {
	PERScalar    PERScalar
	PERExponent  PERExponent
	IEExtensions *PacketErrorRateExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *PacketErrorRate) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "pER-Scalar", Optional: false},
			per.ComponentInfo{Name: "pER-Exponent", Optional: false},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.IEExtensions != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if err := seqEncoder.EncodePreamble(optionalBitmap); err != nil {
		return err
	}

	if err = w.EncodeInteger(int64(s.PERScalar.Value), per.ConstrainedExtensible(0, 9)); err != nil {
		return fmt.Errorf("encode PERScalar failed: %w", err)
	}
	if err = w.EncodeInteger(int64(s.PERExponent.Value), per.ConstrainedExtensible(0, 9)); err != nil {
		return fmt.Errorf("encode PERExponent failed: %w", err)
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
func (s *PacketErrorRate) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "pER-Scalar", Optional: false},
			per.ComponentInfo{Name: "pER-Exponent", Optional: false},
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
		val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 9))
		if err != nil {
			return fmt.Errorf("decode PERScalar failed: %w", err)
		}
		s.PERScalar.Value = val
	}

	{
		val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 9))
		if err != nil {
			return fmt.Errorf("decode PERExponent failed: %w", err)
		}
		s.PERExponent.Value = val
	}

	if seqDecoder.IsComponentPresent(2) {
		s.IEExtensions = new(PacketErrorRateExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
