package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// M4Configuration is a generated SEQUENCE type.
type M4Configuration struct {
	M4period     M4period
	M4LinksToLog LinksToLog
	IEExtensions *M4ConfigurationExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *M4Configuration) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "m4period", Optional: false},
			per.ComponentInfo{Name: "m4-links-to-log", Optional: false},
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

	{
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 5), ExtValues: nil}
		if err = w.EncodeEnumerated(int64(s.M4period.Value), enumC); err != nil {
			return fmt.Errorf("encode M4period failed: %w", err)
		}
	}

	{
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3), ExtValues: nil}
		if err = w.EncodeEnumerated(int64(s.M4LinksToLog.Value), enumC); err != nil {
			return fmt.Errorf("encode M4LinksToLog failed: %w", err)
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
func (s *M4Configuration) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "m4period", Optional: false},
			per.ComponentInfo{Name: "m4-links-to-log", Optional: false},
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
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 5), ExtValues: nil}
		val, err := r.DecodeEnumerated(enumC)
		if err != nil {
			return fmt.Errorf("decode M4period failed: %w", err)
		}
		s.M4period.Value = val
	}

	{
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3), ExtValues: nil}
		val, err := r.DecodeEnumerated(enumC)
		if err != nil {
			return fmt.Errorf("decode M4LinksToLog failed: %w", err)
		}
		s.M4LinksToLog.Value = val
	}

	if seqDecoder.IsComponentPresent(2) {
		s.IEExtensions = new(M4ConfigurationExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
