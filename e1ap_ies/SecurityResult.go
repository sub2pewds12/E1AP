package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// SecurityResult is a generated SEQUENCE type.
type SecurityResult struct {
	IntegrityProtectionResult       IntegrityProtectionResult       `aper:"mandatory,ext"`
	ConfidentialityProtectionResult ConfidentialityProtectionResult `aper:"mandatory,ext"`
	IEExtensions                    *ProtocolExtensionContainer     `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *SecurityResult) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(1), &aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = s.IntegrityProtectionResult.Encode(w); err != nil {
		return fmt.Errorf("Encode IntegrityProtectionResult failed: %w", err)
	}
	if err = s.ConfidentialityProtectionResult.Encode(w); err != nil {
		return fmt.Errorf("Encode ConfidentialityProtectionResult failed: %w", err)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *SecurityResult) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if err = s.IntegrityProtectionResult.Decode(r); err != nil {
		return fmt.Errorf("Decode IntegrityProtectionResult failed: %w", err)
	}
	if err = s.ConfidentialityProtectionResult.Decode(r); err != nil {
		return fmt.Errorf("Decode ConfidentialityProtectionResult failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for SecurityResult")
	}
	return nil
}
