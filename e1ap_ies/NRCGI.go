package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// NRCGI is a generated SEQUENCE type.
type NRCGI struct {
	PLMNIdentity   PLMNIdentity
	NRCellIdentity NRCellIdentity
	IEExtensions   *NRCGIExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *NRCGI) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: false,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "pLMN-Identity", Optional: false},
			per.ComponentInfo{Name: "nR-Cell-Identity", Optional: false},
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

	if err = w.EncodeOctetString([]byte(s.PLMNIdentity.Value), per.SizeConstraints{Extensible: false, Min: int64Ptr(3), Max: int64Ptr(3)}); err != nil {
		return fmt.Errorf("encode PLMNIdentity failed: %w", err)
	}
	if err = w.EncodeBitString(s.NRCellIdentity.Value, per.SizeConstraints{Extensible: false, Min: int64Ptr(36), Max: int64Ptr(36)}); err != nil {
		return fmt.Errorf("encode NRCellIdentity failed: %w", err)
	}

	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *NRCGI) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: false,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "pLMN-Identity", Optional: false},
			per.ComponentInfo{Name: "nR-Cell-Identity", Optional: false},
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
		val, err := r.DecodeOctetString(per.SizeConstraints{Extensible: false, Min: int64Ptr(3), Max: int64Ptr(3)})
		if err != nil {
			return fmt.Errorf("decode PLMNIdentity failed: %w", err)
		}
		s.PLMNIdentity.Value = val
	}

	{
		val, err := r.DecodeBitString(per.SizeConstraints{Extensible: false, Min: int64Ptr(36), Max: int64Ptr(36)})
		if err != nil {
			return fmt.Errorf("decode NRCellIdentity failed: %w", err)
		}
		s.NRCellIdentity.Value = val
	}

	if seqDecoder.IsComponentPresent(2) {
		s.IEExtensions = new(NRCGIExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	return nil
}
