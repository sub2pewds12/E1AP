package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// NPNSupportInfoSNPN is a generated SEQUENCE type.
type NPNSupportInfoSNPN struct {
	NID          NID
	IEExtensions *NPNSupportInfoSNPNExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *NPNSupportInfoSNPN) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: false,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "nID", Optional: false},
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

	if err = w.EncodeBitString(s.NID.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(44), Max: int64Ptr(44)}); err != nil {
		return fmt.Errorf("encode NID failed: %w", err)
	}

	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *NPNSupportInfoSNPN) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: false,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "nID", Optional: false},
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
		val, err := r.DecodeBitString(per.SizeConstraints{Extensible: false, Min: int64Ptr(44), Max: int64Ptr(44)})
		if err != nil {
			return fmt.Errorf("decode NID failed: %w", err)
		}
		s.NID.Value = val
	}

	if seqDecoder.IsComponentPresent(1) {
		s.IEExtensions = new(NPNSupportInfoSNPNExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	return nil
}
