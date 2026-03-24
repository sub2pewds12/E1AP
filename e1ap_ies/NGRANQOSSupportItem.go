package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// NGRANQOSSupportItem is a generated SEQUENCE type.
type NGRANQOSSupportItem struct {
	NonDynamic5QIDescriptor NonDynamic5QIDescriptor
	IEExtensions            *NGRANQOSSupportItemExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *NGRANQOSSupportItem) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: false,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "non-Dynamic5QIDescriptor", Optional: false},
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

	if err = s.NonDynamic5QIDescriptor.Encode(w); err != nil {
		return fmt.Errorf("encode NonDynamic5QIDescriptor failed: %w", err)
	}

	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *NGRANQOSSupportItem) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: false,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "non-Dynamic5QIDescriptor", Optional: false},
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

	if err = s.NonDynamic5QIDescriptor.Decode(r); err != nil {
		return fmt.Errorf("Decode NonDynamic5QIDescriptor failed: %w", err)
	}

	if seqDecoder.IsComponentPresent(1) {
		s.IEExtensions = new(NGRANQOSSupportItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	return nil
}
