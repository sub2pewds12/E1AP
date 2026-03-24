package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// EUTRANQOS is a generated SEQUENCE type.
type EUTRANQOS struct {
	QCI                                  QCI
	EUTRANallocationAndRetentionPriority EUTRANAllocationAndRetentionPriority
	GbrQosInformation                    *GBRQosInformation
	IEExtensions                         *EUTRANQOSExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *EUTRANQOS) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "qCI", Optional: false},
			per.ComponentInfo{Name: "eUTRANallocationAndRetentionPriority", Optional: false},
			per.ComponentInfo{Name: "gbrQosInformation", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.GbrQosInformation != nil {
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

	if err = w.EncodeInteger(int64(s.QCI.Value), per.Constrained(0, 255)); err != nil {
		return fmt.Errorf("encode QCI failed: %w", err)
	}
	if err = s.EUTRANallocationAndRetentionPriority.Encode(w); err != nil {
		return fmt.Errorf("encode EUTRANallocationAndRetentionPriority failed: %w", err)
	}

	if s.GbrQosInformation != nil {
		if err = s.GbrQosInformation.Encode(w); err != nil {
			return fmt.Errorf("encode GbrQosInformation failed: %w", err)
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
func (s *EUTRANQOS) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "qCI", Optional: false},
			per.ComponentInfo{Name: "eUTRANallocationAndRetentionPriority", Optional: false},
			per.ComponentInfo{Name: "gbrQosInformation", Optional: true},
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
			return fmt.Errorf("decode QCI failed: %w", err)
		}
		s.QCI.Value = val
	}
	if err = s.EUTRANallocationAndRetentionPriority.Decode(r); err != nil {
		return fmt.Errorf("Decode EUTRANallocationAndRetentionPriority failed: %w", err)
	}

	if seqDecoder.IsComponentPresent(2) {
		s.GbrQosInformation = new(GBRQosInformation)
		if err = s.GbrQosInformation.Decode(r); err != nil {
			return fmt.Errorf("Decode GbrQosInformation failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(3) {
		s.IEExtensions = new(EUTRANQOSExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
