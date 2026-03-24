package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// SDAPConfiguration is a generated SEQUENCE type.
type SDAPConfiguration struct {
	DefaultDRB   DefaultDRB
	SDAPHeaderUL SDAPHeaderUL
	SDAPHeaderDL SDAPHeaderDL
	IEExtensions *SDAPConfigurationExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *SDAPConfiguration) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "defaultDRB", Optional: false},
			per.ComponentInfo{Name: "sDAP-Header-UL", Optional: false},
			per.ComponentInfo{Name: "sDAP-Header-DL", Optional: false},
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
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
		if err = w.EncodeEnumerated(int64(s.DefaultDRB.Value), enumC); err != nil {
			return fmt.Errorf("encode DefaultDRB failed: %w", err)
		}
	}

	{
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
		if err = w.EncodeEnumerated(int64(s.SDAPHeaderUL.Value), enumC); err != nil {
			return fmt.Errorf("encode SDAPHeaderUL failed: %w", err)
		}
	}

	{
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
		if err = w.EncodeEnumerated(int64(s.SDAPHeaderDL.Value), enumC); err != nil {
			return fmt.Errorf("encode SDAPHeaderDL failed: %w", err)
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
func (s *SDAPConfiguration) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "defaultDRB", Optional: false},
			per.ComponentInfo{Name: "sDAP-Header-UL", Optional: false},
			per.ComponentInfo{Name: "sDAP-Header-DL", Optional: false},
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
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
		val, err := r.DecodeEnumerated(enumC)
		if err != nil {
			return fmt.Errorf("decode DefaultDRB failed: %w", err)
		}
		s.DefaultDRB.Value = val
	}

	{
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
		val, err := r.DecodeEnumerated(enumC)
		if err != nil {
			return fmt.Errorf("decode SDAPHeaderUL failed: %w", err)
		}
		s.SDAPHeaderUL.Value = val
	}

	{
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
		val, err := r.DecodeEnumerated(enumC)
		if err != nil {
			return fmt.Errorf("decode SDAPHeaderDL failed: %w", err)
		}
		s.SDAPHeaderDL.Value = val
	}

	if seqDecoder.IsComponentPresent(3) {
		s.IEExtensions = new(SDAPConfigurationExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
