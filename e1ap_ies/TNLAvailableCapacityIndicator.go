package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// TNLAvailableCapacityIndicator is a generated SEQUENCE type.
type TNLAvailableCapacityIndicator struct {
	DLTNLOfferedCapacity   TNLAvailableCapacityIndicatorDLTNLOfferedCapacity
	DLTNLAvailableCapacity TNLAvailableCapacityIndicatorDLTNLAvailableCapacity
	ULTNLOfferedCapacity   TNLAvailableCapacityIndicatorULTNLOfferedCapacity
	ULTNLAvailableCapacity TNLAvailableCapacityIndicatorULTNLAvailableCapacity
	IEExtensions           *TNLAvailableCapacityIndicatorExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *TNLAvailableCapacityIndicator) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "dL-TNL-OfferedCapacity", Optional: false},
			per.ComponentInfo{Name: "dL-TNL-AvailableCapacity", Optional: false},
			per.ComponentInfo{Name: "uL-TNL-OfferedCapacity", Optional: false},
			per.ComponentInfo{Name: "uL-TNL-AvailableCapacity", Optional: false},
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

	if err = w.EncodeInteger(int64(s.DLTNLOfferedCapacity.Value), per.ConstrainedExtensible(0, 16777216)); err != nil {
		return fmt.Errorf("encode DLTNLOfferedCapacity failed: %w", err)
	}
	if err = w.EncodeInteger(int64(s.DLTNLAvailableCapacity.Value), per.ConstrainedExtensible(0, 100)); err != nil {
		return fmt.Errorf("encode DLTNLAvailableCapacity failed: %w", err)
	}
	if err = w.EncodeInteger(int64(s.ULTNLOfferedCapacity.Value), per.ConstrainedExtensible(0, 16777216)); err != nil {
		return fmt.Errorf("encode ULTNLOfferedCapacity failed: %w", err)
	}
	if err = w.EncodeInteger(int64(s.ULTNLAvailableCapacity.Value), per.ConstrainedExtensible(0, 100)); err != nil {
		return fmt.Errorf("encode ULTNLAvailableCapacity failed: %w", err)
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
func (s *TNLAvailableCapacityIndicator) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "dL-TNL-OfferedCapacity", Optional: false},
			per.ComponentInfo{Name: "dL-TNL-AvailableCapacity", Optional: false},
			per.ComponentInfo{Name: "uL-TNL-OfferedCapacity", Optional: false},
			per.ComponentInfo{Name: "uL-TNL-AvailableCapacity", Optional: false},
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
		val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 16777216))
		if err != nil {
			return fmt.Errorf("decode DLTNLOfferedCapacity failed: %w", err)
		}
		s.DLTNLOfferedCapacity.Value = val
	}

	{
		val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 100))
		if err != nil {
			return fmt.Errorf("decode DLTNLAvailableCapacity failed: %w", err)
		}
		s.DLTNLAvailableCapacity.Value = val
	}

	{
		val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 16777216))
		if err != nil {
			return fmt.Errorf("decode ULTNLOfferedCapacity failed: %w", err)
		}
		s.ULTNLOfferedCapacity.Value = val
	}

	{
		val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 100))
		if err != nil {
			return fmt.Errorf("decode ULTNLAvailableCapacity failed: %w", err)
		}
		s.ULTNLAvailableCapacity.Value = val
	}

	if seqDecoder.IsComponentPresent(4) {
		s.IEExtensions = new(TNLAvailableCapacityIndicatorExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
