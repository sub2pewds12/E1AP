package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// QOSFlowFailedItem is a generated SEQUENCE type.
type QOSFlowFailedItem struct {
	QOSFlowIdentifier QOSFlowIdentifier
	Cause             Cause
	IEExtensions      *QOSFlowFailedItemExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *QOSFlowFailedItem) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "qoS-Flow-Identifier", Optional: false},
			per.ComponentInfo{Name: "cause", Optional: false},
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

	if err = w.EncodeInteger(int64(s.QOSFlowIdentifier.Value), per.Constrained(0, 63)); err != nil {
		return fmt.Errorf("encode QOSFlowIdentifier failed: %w", err)
	}
	if err = s.Cause.Encode(w); err != nil {
		return fmt.Errorf("encode Cause failed: %w", err)
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
func (s *QOSFlowFailedItem) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "qoS-Flow-Identifier", Optional: false},
			per.ComponentInfo{Name: "cause", Optional: false},
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
		val, err := r.DecodeInteger(per.Constrained(0, 63))
		if err != nil {
			return fmt.Errorf("decode QOSFlowIdentifier failed: %w", err)
		}
		s.QOSFlowIdentifier.Value = val
	}
	if err = s.Cause.Decode(r); err != nil {
		return fmt.Errorf("Decode Cause failed: %w", err)
	}

	if seqDecoder.IsComponentPresent(2) {
		s.IEExtensions = new(QOSFlowFailedItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
