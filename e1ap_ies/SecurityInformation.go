package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/asn1go/per"
)

// SecurityInformation is a generated SEQUENCE type.
type SecurityInformation struct {
	SecurityAlgorithm SecurityAlgorithm
	UPSecuritykey     UPSecuritykey
	IEExtensions      *SecurityInformationExtensions
}

// Encode implements the aper.AperMarshaller interface.
func (s *SecurityInformation) Encode(w *per.Encoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "securityAlgorithm", Optional: false},
			per.ComponentInfo{Name: "uPSecuritykey", Optional: false},
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

	if err = s.SecurityAlgorithm.Encode(w); err != nil {
		return fmt.Errorf("encode SecurityAlgorithm failed: %w", err)
	}
	if err = s.UPSecuritykey.Encode(w); err != nil {
		return fmt.Errorf("encode UPSecuritykey failed: %w", err)
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
func (s *SecurityInformation) Decode(r *per.Decoder) (err error) {

	c := per.SequenceConstraints{
		Extensible: true,
		RootComponents: []per.ComponentInfo{
			per.ComponentInfo{Name: "securityAlgorithm", Optional: false},
			per.ComponentInfo{Name: "uPSecuritykey", Optional: false},
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

	if err = s.SecurityAlgorithm.Decode(r); err != nil {
		return fmt.Errorf("Decode SecurityAlgorithm failed: %w", err)
	}
	if err = s.UPSecuritykey.Decode(r); err != nil {
		return fmt.Errorf("Decode UPSecuritykey failed: %w", err)
	}

	if seqDecoder.IsComponentPresent(2) {
		s.IEExtensions = new(SecurityInformationExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}

	if _, err := seqDecoder.DecodeExtensionAdditions(); err != nil {
		return err
	}

	return nil
}
