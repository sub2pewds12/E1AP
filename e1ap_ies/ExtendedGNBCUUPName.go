package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// ExtendedGNBCUUPName is a generated SEQUENCE type.
type ExtendedGNBCUUPName struct {
	GNBCUUPNameVisibleString *GNBCUUPNameVisibleString
	GNBCUUPNameUTF8String    *GNBCUUPNameUTF8String
	IEExtensions             *ExtendedGNBCUUPNameExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *ExtendedGNBCUUPName) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "gNB-CU-UP-NameVisibleString", Optional: true},
			per.ComponentInfo{Name: "gNB-CU-UP-NameUTF8String", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.GNBCUUPNameVisibleString != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.GNBCUUPNameUTF8String != nil {
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

	if s.GNBCUUPNameVisibleString != nil {
		if err = s.GNBCUUPNameVisibleString.Encode(w); err != nil {
			return fmt.Errorf("encode GNBCUUPNameVisibleString failed: %w", err)
		}
	}

	if s.GNBCUUPNameUTF8String != nil {
		if err = s.GNBCUUPNameUTF8String.Encode(w); err != nil {
			return fmt.Errorf("encode GNBCUUPNameUTF8String failed: %w", err)
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
func (s *ExtendedGNBCUUPName) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "gNB-CU-UP-NameVisibleString", Optional: true},
			per.ComponentInfo{Name: "gNB-CU-UP-NameUTF8String", Optional: true},
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
		s.GNBCUUPNameVisibleString = new(GNBCUUPNameVisibleString)
		if err = s.GNBCUUPNameVisibleString.Decode(r); err != nil {
			return fmt.Errorf("Decode GNBCUUPNameVisibleString failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(1) {
		s.GNBCUUPNameUTF8String = new(GNBCUUPNameUTF8String)
		if err = s.GNBCUUPNameUTF8String.Decode(r); err != nil {
			return fmt.Errorf("Decode GNBCUUPNameUTF8String failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(2) {
		s.IEExtensions = new(ExtendedGNBCUUPNameExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
