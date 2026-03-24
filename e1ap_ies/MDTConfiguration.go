package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// MDTConfiguration is a generated SEQUENCE type.
type MDTConfiguration struct {
	MdtActivation MDTActivation
	MDTMode       MDTMode
	IEExtensions  *MDTConfigurationExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *MDTConfiguration) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "mdt-Activation", Optional: false},
			per.ComponentInfo{Name: "mDTMode", Optional: false},
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

	{
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
		if err = w.EncodeEnumerated(int64(s.MdtActivation.Value), enumC); err != nil {
			return fmt.Errorf("encode MdtActivation failed: %w", err)
		}
	}
	if err = s.MDTMode.Encode(w); err != nil {
		return fmt.Errorf("encode MDTMode failed: %w", err)
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
func (s *MDTConfiguration) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "mdt-Activation", Optional: false},
			per.ComponentInfo{Name: "mDTMode", Optional: false},
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
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
		val, err := r.DecodeEnumerated(enumC)
		if err != nil {
			return fmt.Errorf("decode MdtActivation failed: %w", err)
		}
		s.MdtActivation.Value = val
	}
	if err = s.MDTMode.Decode(r); err != nil {
		return fmt.Errorf("Decode MDTMode failed: %w", err)
	}

	if seqDecoder.IsComponentPresent(2) {
		s.IEExtensions = new(MDTConfigurationExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
