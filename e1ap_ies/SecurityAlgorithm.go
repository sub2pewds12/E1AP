package e1ap_ies

import (
	"fmt"

	"asn1go/per"
)

// SecurityAlgorithm is a generated SEQUENCE type.
type SecurityAlgorithm struct {
	CipheringAlgorithm           CipheringAlgorithm
	IntegrityProtectionAlgorithm *IntegrityProtectionAlgorithm
	IEExtensions                 *SecurityAlgorithmExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *SecurityAlgorithm) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "cipheringAlgorithm", Optional: false},
			per.ComponentInfo{Name: "integrityProtectionAlgorithm", Optional: true},
			per.ComponentInfo{Name: "iE-Extensions", Optional: true},
		},
	}
	seqEncoder := w.NewSequenceEncoder(c)
	if err := seqEncoder.EncodeExtensionBit(false); err != nil {
		return err
	}

	optionalBitmap := make([]bool, 0)

	if s.IntegrityProtectionAlgorithm != nil {
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
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 4), ExtValues: nil}
		if err = w.EncodeEnumerated(int64(s.CipheringAlgorithm.Value), enumC); err != nil {
			return fmt.Errorf("encode CipheringAlgorithm failed: %w", err)
		}
	}

	if s.IntegrityProtectionAlgorithm != nil {

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 4), ExtValues: nil}
			if err = w.EncodeEnumerated(int64((*s.IntegrityProtectionAlgorithm).Value), enumC); err != nil {
				return fmt.Errorf("encode IntegrityProtectionAlgorithm failed: %w", err)
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
func (s *SecurityAlgorithm) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "cipheringAlgorithm", Optional: false},
			per.ComponentInfo{Name: "integrityProtectionAlgorithm", Optional: true},
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
		enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 4), ExtValues: nil}
		val, err := r.DecodeEnumerated(enumC)
		if err != nil {
			return fmt.Errorf("decode CipheringAlgorithm failed: %w", err)
		}
		s.CipheringAlgorithm.Value = val
	}

	if seqDecoder.IsComponentPresent(1) {
		s.IntegrityProtectionAlgorithm = new(IntegrityProtectionAlgorithm)

		{
			enumC := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 4), ExtValues: nil}
			val, err := r.DecodeEnumerated(enumC)
			if err != nil {
				return fmt.Errorf("decode IntegrityProtectionAlgorithm failed: %w", err)
			}
			s.IntegrityProtectionAlgorithm.Value = val
		}
	}

	if seqDecoder.IsComponentPresent(2) {
		s.IEExtensions = new(SecurityAlgorithmExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
