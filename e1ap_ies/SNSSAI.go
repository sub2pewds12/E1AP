package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// SNSSAI is a generated SEQUENCE type.
type SNSSAI struct {
	SST          SNSSAISST
	SD           *SNSSAISD
	IEExtensions *SNSSAIExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *SNSSAI) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "sST", Optional: false},
			per.ComponentInfo{Name: "sD", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.SD != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.IEExtensions != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if err := seqEncoder.EncodePreamble(optionalBitmap); err != nil {
		return err
	}

	if err = w.EncodeOctetString([]byte(s.SST.Value), per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(1)}); err != nil {
		return fmt.Errorf("encode SST failed: %w", err)
	}

	if s.SD != nil {
		if err = w.EncodeOctetString([]byte((*s.SD).Value), per.SizeConstraints{Extensible: false, Min: int64Ptr(3), Max: int64Ptr(3)}); err != nil {
			return fmt.Errorf("encode SD failed: %w", err)
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
func (s *SNSSAI) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "sST", Optional: false},
			per.ComponentInfo{Name: "sD", Optional: true},
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
		val, err := r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(1)})
		if err != nil {
			return fmt.Errorf("decode SST failed: %w", err)
		}
		s.SST.Value = val
	}

	if seqDecoder.IsComponentPresent(1) {
		s.SD = new(SNSSAISD)

		{
			val, err := r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(3), Max: int64Ptr(3)})
			if err != nil {
				return fmt.Errorf("decode SD failed: %w", err)
			}
			s.SD.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(2) {
		s.IEExtensions = new(SNSSAIExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
