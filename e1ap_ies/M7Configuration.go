package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// M7Configuration is a generated SEQUENCE type.
type M7Configuration struct {
	M7period     M7period
	M7LinksToLog LinksToLog
	IEExtensions *M7ConfigurationExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *M7Configuration) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "m7period", Optional: false},
			per.ComponentInfo{Name: "m7-links-to-log", Optional: false},
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

	if err = w.EncodeInteger(int64(s.M7period.Value), per.ConstrainedExtensible(1, 60)); err != nil {
		return fmt.Errorf("encode M7period failed: %w", err)
	}

	{
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3), ExtValues: nil}
		if err = w.EncodeEnumerated(int64(s.M7LinksToLog.Value), enumC); err != nil {
			return fmt.Errorf("encode M7LinksToLog failed: %w", err)
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
func (s *M7Configuration) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "m7period", Optional: false},
			per.ComponentInfo{Name: "m7-links-to-log", Optional: false},
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
		val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 60))
		if err != nil {
			return fmt.Errorf("decode M7period failed: %w", err)
		}
		s.M7period.Value = val
	}

	{
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3), ExtValues: nil}
		val, err := r.DecodeEnumerated(enumC)
		if err != nil {
			return fmt.Errorf("decode M7LinksToLog failed: %w", err)
		}
		s.M7LinksToLog.Value = val
	}

	if seqDecoder.IsComponentPresent(2) {
		s.IEExtensions = new(M7ConfigurationExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
