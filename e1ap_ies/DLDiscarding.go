package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// DLDiscarding is a generated SEQUENCE type.
type DLDiscarding struct {
	DLDiscardingCountVal PDCPCount
	IEExtensions         *DLDiscardingExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *DLDiscarding) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: false,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "dLDiscardingCountVal", Optional: false},
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

	if err = s.DLDiscardingCountVal.Encode(w); err != nil {
		return fmt.Errorf("encode DLDiscardingCountVal failed: %w", err)
	}

	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DLDiscarding) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: false,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "dLDiscardingCountVal", Optional: false},
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

	if err = s.DLDiscardingCountVal.Decode(r); err != nil {
		return fmt.Errorf("Decode DLDiscardingCountVal failed: %w", err)
	}

	if seqDecoder.IsComponentPresent(1) {
		s.IEExtensions = new(DLDiscardingExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	return nil
}
