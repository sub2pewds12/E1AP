package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// DAPSRequestInfo is a generated SEQUENCE type.
type DAPSRequestInfo struct {
	DapsIndicator DAPSRequestInfoDapsIndicator
	IEExtensions  *DAPSRequestInfoExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *DAPSRequestInfo) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "dapsIndicator", Optional: false},
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
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1), ExtValues: nil}
		if err = w.EncodeEnumerated(int64(s.DapsIndicator.Value), enumC); err != nil {
			return fmt.Errorf("encode DapsIndicator failed: %w", err)
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
func (s *DAPSRequestInfo) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "dapsIndicator", Optional: false},
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
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1), ExtValues: nil}
		val, err := r.DecodeEnumerated(enumC)
		if err != nil {
			return fmt.Errorf("decode DapsIndicator failed: %w", err)
		}
		s.DapsIndicator.Value = val
	}

	if seqDecoder.IsComponentPresent(1) {
		s.IEExtensions = new(DAPSRequestInfoExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
