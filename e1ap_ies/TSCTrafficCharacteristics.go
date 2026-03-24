package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// TSCTrafficCharacteristics is a generated SEQUENCE type.
type TSCTrafficCharacteristics struct {
	TSCTrafficCharacteristicsUL *TSCTrafficInformation
	TSCTrafficCharacteristicsDL *TSCTrafficInformation
	IEExtensions                *TSCTrafficCharacteristicsExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *TSCTrafficCharacteristics) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: false,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "tSCTrafficCharacteristicsUL", Optional: true},
			per.ComponentInfo{Name: "tSCTrafficCharacteristicsDL", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.TSCTrafficCharacteristicsUL != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.TSCTrafficCharacteristicsDL != nil {
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

	if s.TSCTrafficCharacteristicsUL != nil {
		if err = s.TSCTrafficCharacteristicsUL.Encode(w); err != nil {
			return fmt.Errorf("encode TSCTrafficCharacteristicsUL failed: %w", err)
		}
	}

	if s.TSCTrafficCharacteristicsDL != nil {
		if err = s.TSCTrafficCharacteristicsDL.Encode(w); err != nil {
			return fmt.Errorf("encode TSCTrafficCharacteristicsDL failed: %w", err)
		}
	}

	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *TSCTrafficCharacteristics) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: false,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "tSCTrafficCharacteristicsUL", Optional: true},
			per.ComponentInfo{Name: "tSCTrafficCharacteristicsDL", Optional: true},
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

	if seqDecoder.IsComponentPresent(0) {
		s.TSCTrafficCharacteristicsUL = new(TSCTrafficInformation)
		if err = s.TSCTrafficCharacteristicsUL.Decode(r); err != nil {
			return fmt.Errorf("Decode TSCTrafficCharacteristicsUL failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(1) {
		s.TSCTrafficCharacteristicsDL = new(TSCTrafficInformation)
		if err = s.TSCTrafficCharacteristicsDL.Decode(r); err != nil {
			return fmt.Errorf("Decode TSCTrafficCharacteristicsDL failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(2) {
		s.IEExtensions = new(TSCTrafficCharacteristicsExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	return nil
}
