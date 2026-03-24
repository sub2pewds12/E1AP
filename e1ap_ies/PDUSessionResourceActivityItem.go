package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// PDUSessionResourceActivityItem is a generated SEQUENCE type.
type PDUSessionResourceActivityItem struct {
	PDUSessionID               PDUSessionID
	PDUSessionResourceActivity PDUSessionResourceActivity
	IEExtensions               *PDUSessionResourceActivityItemExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDUSessionResourceActivityItem) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "pDU-Session-ID", Optional: false},
			per.ComponentInfo{Name: "pDU-Session-Resource-Activity", Optional: false},
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

	if err = w.EncodeInteger(int64(s.PDUSessionID.Value), per.Constrained(0, 255)); err != nil {
		return fmt.Errorf("encode PDUSessionID failed: %w", err)
	}

	{
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
		if err = w.EncodeEnumerated(int64(s.PDUSessionResourceActivity.Value), enumC); err != nil {
			return fmt.Errorf("encode PDUSessionResourceActivity failed: %w", err)
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
func (s *PDUSessionResourceActivityItem) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "pDU-Session-ID", Optional: false},
			per.ComponentInfo{Name: "pDU-Session-Resource-Activity", Optional: false},
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
		val, err := r.DecodeInteger(per.Constrained(0, 255))
		if err != nil {
			return fmt.Errorf("decode PDUSessionID failed: %w", err)
		}
		s.PDUSessionID.Value = val
	}

	{
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
		val, err := r.DecodeEnumerated(enumC)
		if err != nil {
			return fmt.Errorf("decode PDUSessionResourceActivity failed: %w", err)
		}
		s.PDUSessionResourceActivity.Value = val
	}

	if seqDecoder.IsComponentPresent(2) {
		s.IEExtensions = new(PDUSessionResourceActivityItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
