package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// UplinkOnlyROHC is a generated SEQUENCE type.
type UplinkOnlyROHC struct {
	MaxCID       UplinkOnlyROHCMaxCID
	ROHCProfiles UplinkOnlyROHCROHCProfiles
	ContinueROHC *UplinkOnlyROHCContinueROHC
	IEExtensions *UplinkOnlyROHCExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *UplinkOnlyROHC) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "maxCID", Optional: false},
			per.ComponentInfo{Name: "rOHC-Profiles", Optional: false},
			per.ComponentInfo{Name: "continueROHC", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.ContinueROHC != nil {
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

	if err = w.EncodeInteger(int64(s.MaxCID.Value), per.ConstrainedExtensible(0, 16383)); err != nil {
		return fmt.Errorf("encode MaxCID failed: %w", err)
	}
	if err = w.EncodeInteger(int64(s.ROHCProfiles.Value), per.ConstrainedExtensible(0, 511)); err != nil {
		return fmt.Errorf("encode ROHCProfiles failed: %w", err)
	}

	if s.ContinueROHC != nil {

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1), ExtValues: nil}
			if err = w.EncodeEnumerated(int64((*s.ContinueROHC).Value), enumC); err != nil {
				return fmt.Errorf("encode ContinueROHC failed: %w", err)
			}
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
func (s *UplinkOnlyROHC) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "maxCID", Optional: false},
			per.ComponentInfo{Name: "rOHC-Profiles", Optional: false},
			per.ComponentInfo{Name: "continueROHC", Optional: true},
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
		val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 16383))
		if err != nil {
			return fmt.Errorf("decode MaxCID failed: %w", err)
		}
		s.MaxCID.Value = val
	}

	{
		val, err := r.DecodeInteger(per.ConstrainedExtensible(0, 511))
		if err != nil {
			return fmt.Errorf("decode ROHCProfiles failed: %w", err)
		}
		s.ROHCProfiles.Value = val
	}

	if seqDecoder.IsComponentPresent(2) {
		s.ContinueROHC = new(UplinkOnlyROHCContinueROHC)

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 1), ExtValues: nil}
			val, err := r.DecodeEnumerated(enumC)
			if err != nil {
				return fmt.Errorf("decode ContinueROHC failed: %w", err)
			}
			s.ContinueROHC.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(3) {
		s.IEExtensions = new(UplinkOnlyROHCExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
