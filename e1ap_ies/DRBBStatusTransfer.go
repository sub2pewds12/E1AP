package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// DRBBStatusTransfer is a generated SEQUENCE type.
type DRBBStatusTransfer struct {
	ReceiveStatusofPDCPSDU *DRBBStatusTransferReceiveStatusofPDCPSDU
	CountValue             PDCPCount
	IEExtension            *DRBBStatusTransferExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBBStatusTransfer) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "receiveStatusofPDCPSDU", Optional: true},
			per.ComponentInfo{Name: "countValue", Optional: false},
			per.ComponentInfo{Name: "iE-Extension", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.ReceiveStatusofPDCPSDU != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.IEExtension != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if err := seqEncoder.EncodePreamble(optionalBitmap); err != nil {
		return err
	}

	if s.ReceiveStatusofPDCPSDU != nil {
		if err = w.EncodeBitString((*s.ReceiveStatusofPDCPSDU).Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(131072)}); err != nil {
			return fmt.Errorf("encode ReceiveStatusofPDCPSDU failed: %w", err)
		}
	}
	if err = s.CountValue.Encode(w); err != nil {
		return fmt.Errorf("encode CountValue failed: %w", err)
	}

	if s.IEExtension != nil {
		if err = s.IEExtension.Encode(w); err != nil {
			return fmt.Errorf("encode IEExtension failed: %w", err)
		}
	}

	if err := seqEncoder.EncodeExtensionAdditions([]bool{}, [][]byte{}); err != nil {
		return err
	}

	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DRBBStatusTransfer) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "receiveStatusofPDCPSDU", Optional: true},
			per.ComponentInfo{Name: "countValue", Optional: false},
			per.ComponentInfo{Name: "iE-Extension", Optional: true},
		},
	}
	seqDecoder := r.NewSequenceDecoder(c)
	if err := seqDecoder.DecodeExtensionBit(); err != nil {
		return err
	}

	if err := seqDecoder.DecodePreamble(); err != nil {
		return err
	}

	if seqDecoder.IsComponentPresent(0) {
		s.ReceiveStatusofPDCPSDU = new(DRBBStatusTransferReceiveStatusofPDCPSDU)

		{
			val, err := r.DecodeBitString(per.SizeConstraints{Extensible: false, Min: int64Ptr(1), Max: int64Ptr(131072)})
			if err != nil {
				return fmt.Errorf("decode ReceiveStatusofPDCPSDU failed: %w", err)
			}
			s.ReceiveStatusofPDCPSDU.Value = val
		}
	}
	if err = s.CountValue.Decode(r); err != nil {
		return fmt.Errorf("Decode CountValue failed: %w", err)
	}

	if seqDecoder.IsComponentPresent(2) {
		s.IEExtension = new(DRBBStatusTransferExtensions)
		if err = s.IEExtension.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtension failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
