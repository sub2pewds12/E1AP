package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// DRBMeasurementResultsInformationItem is a generated SEQUENCE type.
type DRBMeasurementResultsInformationItem struct {
	DRBID        DRBID
	ULD1Result   *DRBMeasurementResultsInformationItemULD1Result
	IEExtensions *DRBMeasurementResultsInformationItemExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBMeasurementResultsInformationItem) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "dRB-ID", Optional: false},
			per.ComponentInfo{Name: "uL-D1-Result", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.ULD1Result != nil {
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

	if err = w.EncodeInteger(int64(s.DRBID.Value), per.ConstrainedExtensible(1, 32)); err != nil {
		return fmt.Errorf("encode DRBID failed: %w", err)
	}

	if s.ULD1Result != nil {
		if err = w.EncodeInteger(int64((*s.ULD1Result).Value), per.ConstrainedExtensible(0, 10000)); err != nil {
			return fmt.Errorf("encode ULD1Result failed: %w", err)
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
func (s *DRBMeasurementResultsInformationItem) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "dRB-ID", Optional: false},
			per.ComponentInfo{Name: "uL-D1-Result", Optional: true},
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
		val, err := r.DecodeInteger(per.ConstrainedExtensible(1, 32))
		if err != nil {
			return fmt.Errorf("decode DRBID failed: %w", err)
		}
		s.DRBID.Value = val
	}

	if seqDecoder.IsComponentPresent(1) {
		s.ULD1Result = new(DRBMeasurementResultsInformationItemULD1Result)

		{
			val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 10000))
			if err != nil {
				return fmt.Errorf("decode ULD1Result failed: %w", err)
			}
			s.ULD1Result.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(2) {
		s.IEExtensions = new(DRBMeasurementResultsInformationItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
