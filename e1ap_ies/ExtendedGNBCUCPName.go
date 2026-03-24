package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// ExtendedGNBCUCPName is a generated SEQUENCE type.
type ExtendedGNBCUCPName struct {
	GNBCUCPNameVisibleString *GNBCUCPNameVisibleString
	GNBCUCPNameUTF8String    *GNBCUCPNameUTF8String
	IEExtensions             *ExtendedGNBCUCPNameExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *ExtendedGNBCUCPName) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "gNB-CU-CP-NameVisibleString", Optional: true},
			per.ComponentInfo{Name: "gNB-CU-CP-NameUTF8String", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.GNBCUCPNameVisibleString != nil {
		optionalBitmap = append(optionalBitmap, true)
	} else {
		optionalBitmap = append(optionalBitmap, false)
	}

	if s.GNBCUCPNameUTF8String != nil {
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

	if s.GNBCUCPNameVisibleString != nil {
		if err = s.GNBCUCPNameVisibleString.Encode(w); err != nil {
			return fmt.Errorf("encode GNBCUCPNameVisibleString failed: %w", err)
		}
	}

	if s.GNBCUCPNameUTF8String != nil {
		if err = s.GNBCUCPNameUTF8String.Encode(w); err != nil {
			return fmt.Errorf("encode GNBCUCPNameUTF8String failed: %w", err)
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
func (s *ExtendedGNBCUCPName) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "gNB-CU-CP-NameVisibleString", Optional: true},
			per.ComponentInfo{Name: "gNB-CU-CP-NameUTF8String", Optional: true},
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
		s.GNBCUCPNameVisibleString = new(GNBCUCPNameVisibleString)
		if err = s.GNBCUCPNameVisibleString.Decode(r); err != nil {
			return fmt.Errorf("Decode GNBCUCPNameVisibleString failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(1) {
		s.GNBCUCPNameUTF8String = new(GNBCUCPNameUTF8String)
		if err = s.GNBCUCPNameUTF8String.Decode(r); err != nil {
			return fmt.Errorf("Decode GNBCUCPNameUTF8String failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(2) {
		s.IEExtensions = new(ExtendedGNBCUCPNameExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
