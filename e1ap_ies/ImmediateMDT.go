package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// ImmediateMDT is a generated SEQUENCE type.
type ImmediateMDT struct {
	MeasurementsToActivate MeasurementsToActivate
	MeasurementFour        *M4Configuration
	MeasurementSix         *M6Configuration
	MeasurementSeven       *M7Configuration
	IEExtensions           *ImmediateMDTExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *ImmediateMDT) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "measurementsToActivate", Optional: false},
			per.ComponentInfo{Name: "measurementFour", Optional: true},
			per.ComponentInfo{Name: "measurementSix", Optional: true},
			per.ComponentInfo{Name: "measurementSeven", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.MeasurementFour != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.MeasurementSix != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.MeasurementSeven != nil {
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

	if err = w.EncodeBitString(s.MeasurementsToActivate.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(8), Max: int64Ptr(8)}); err != nil {
		return fmt.Errorf("encode MeasurementsToActivate failed: %w", err)
	}

	if s.MeasurementFour != nil {
		if err = s.MeasurementFour.Encode(w); err != nil {
			return fmt.Errorf("encode MeasurementFour failed: %w", err)
		}
	}

	if s.MeasurementSix != nil {
		if err = s.MeasurementSix.Encode(w); err != nil {
			return fmt.Errorf("encode MeasurementSix failed: %w", err)
		}
	}

	if s.MeasurementSeven != nil {
		if err = s.MeasurementSeven.Encode(w); err != nil {
			return fmt.Errorf("encode MeasurementSeven failed: %w", err)
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
func (s *ImmediateMDT) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "measurementsToActivate", Optional: false},
			per.ComponentInfo{Name: "measurementFour", Optional: true},
			per.ComponentInfo{Name: "measurementSix", Optional: true},
			per.ComponentInfo{Name: "measurementSeven", Optional: true},
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
		val, err := r.DecodeBitString(per.SizeConstraints{Extensible: false, Min: int64Ptr(8), Max: int64Ptr(8)})
		if err != nil {
			return fmt.Errorf("decode MeasurementsToActivate failed: %w", err)
		}
		s.MeasurementsToActivate.Value = val
	}

	if seqDecoder.IsComponentPresent(1) {
		s.MeasurementFour = new(M4Configuration)
		if err = s.MeasurementFour.Decode(r); err != nil {
			return fmt.Errorf("Decode MeasurementFour failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(2) {
		s.MeasurementSix = new(M6Configuration)
		if err = s.MeasurementSix.Decode(r); err != nil {
			return fmt.Errorf("Decode MeasurementSix failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(3) {
		s.MeasurementSeven = new(M7Configuration)
		if err = s.MeasurementSeven.Decode(r); err != nil {
			return fmt.Errorf("Decode MeasurementSeven failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(4) {
		s.IEExtensions = new(ImmediateMDTExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
