package e1ap_ies

import (
	"asn1go/per"
)

// UEAssociatedLogicalE1ConnectionListResItemAckItem is a generated SEQUENCE type.
type UEAssociatedLogicalE1ConnectionListResItemAckItem struct {
}

// Encode implements the aper.AperMarshaller interface.
func (s *UEAssociatedLogicalE1ConnectionListResItemAckItem) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible:     false,
		RootComponents: []per.ComponentInfo{},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if err := seqEncoder.EncodePreamble(optionalBitmap); err != nil {
		return err
	}

	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *UEAssociatedLogicalE1ConnectionListResItemAckItem) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible:     false,
		RootComponents: []per.ComponentInfo{},
	}
	seqDecoder := r.NewSequenceDecoder(c)
	if err := seqDecoder.DecodeExtensionBit(); err != nil {
		return err
	}

	if err := seqDecoder.DecodePreamble(); err != nil {
		return err
	}

	return nil
}
