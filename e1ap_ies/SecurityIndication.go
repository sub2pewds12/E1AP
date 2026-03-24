package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// SecurityIndication is a generated SEQUENCE type.
type SecurityIndication struct {
	IntegrityProtectionIndication       IntegrityProtectionIndication
	ConfidentialityProtectionIndication ConfidentialityProtectionIndication
	MaximumIPdatarate                   *MaximumIPdatarate
	IEExtensions                        *SecurityIndicationExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *SecurityIndication) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "integrityProtectionIndication", Optional: false},
			per.ComponentInfo{Name: "confidentialityProtectionIndication", Optional: false},
			per.ComponentInfo{Name: "maximumIPdatarate", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.MaximumIPdatarate != nil {
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

	{
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3), ExtValues: nil}
		if err = w.EncodeEnumerated(int64(s.IntegrityProtectionIndication.Value), enumC); err != nil {
			return fmt.Errorf("encode IntegrityProtectionIndication failed: %w", err)
		}
	}

	{
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3), ExtValues: nil}
		if err = w.EncodeEnumerated(int64(s.ConfidentialityProtectionIndication.Value), enumC); err != nil {
			return fmt.Errorf("encode ConfidentialityProtectionIndication failed: %w", err)
		}
	}

	if s.MaximumIPdatarate != nil {
		if err = s.MaximumIPdatarate.Encode(w); err != nil {
			return fmt.Errorf("encode MaximumIPdatarate failed: %w", err)
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
func (s *SecurityIndication) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "integrityProtectionIndication", Optional: false},
			per.ComponentInfo{Name: "confidentialityProtectionIndication", Optional: false},
			per.ComponentInfo{Name: "maximumIPdatarate", Optional: true},
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
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3), ExtValues: nil}
		val, err := r.DecodeEnumerated(enumC)
		if err != nil {
			return fmt.Errorf("decode IntegrityProtectionIndication failed: %w", err)
		}
		s.IntegrityProtectionIndication.Value = val
	}

	{
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 3), ExtValues: nil}
		val, err := r.DecodeEnumerated(enumC)
		if err != nil {
			return fmt.Errorf("decode ConfidentialityProtectionIndication failed: %w", err)
		}
		s.ConfidentialityProtectionIndication.Value = val
	}

	if seqDecoder.IsComponentPresent(2) {
		s.MaximumIPdatarate = new(MaximumIPdatarate)
		if err = s.MaximumIPdatarate.Decode(r); err != nil {
			return fmt.Errorf("Decode MaximumIPdatarate failed: %w", err)
		}
	}

	if seqDecoder.IsComponentPresent(3) {
		s.IEExtensions = new(SecurityIndicationExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
