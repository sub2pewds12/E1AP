package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// ULUPTNLAddressToUpdateItem is a generated SEQUENCE type.
type ULUPTNLAddressToUpdateItem struct {
	OldTNLAdress TransportLayerAddress
	NewTNLAdress TransportLayerAddress
	IEExtensions *ULUPTNLAddressToUpdateItemExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *ULUPTNLAddressToUpdateItem) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "oldTNLAdress", Optional: false},
			per.ComponentInfo{Name: "newTNLAdress", Optional: false},
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

	if err = w.EncodeBitString(s.OldTNLAdress.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(160)}); err != nil {
		return fmt.Errorf("encode OldTNLAdress failed: %w", err)
	}
	if err = w.EncodeBitString(s.NewTNLAdress.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(160)}); err != nil {
		return fmt.Errorf("encode NewTNLAdress failed: %w", err)
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
func (s *ULUPTNLAddressToUpdateItem) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "oldTNLAdress", Optional: false},
			per.ComponentInfo{Name: "newTNLAdress", Optional: false},
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
		val, err := r.DecodeBitString(per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(160)})
		if err != nil {
			return fmt.Errorf("decode OldTNLAdress failed: %w", err)
		}
		s.OldTNLAdress.Value = val
	}

	{
		val, err := r.DecodeBitString(per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(160)})
		if err != nil {
			return fmt.Errorf("decode NewTNLAdress failed: %w", err)
		}
		s.NewTNLAdress.Value = val
	}

	if seqDecoder.IsComponentPresent(2) {
		s.IEExtensions = new(ULUPTNLAddressToUpdateItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
