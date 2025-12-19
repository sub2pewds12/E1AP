package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// SecurityInformation is a generated SEQUENCE type.
type SecurityInformation struct {
	SecurityAlgorithm SecurityAlgorithm           `aper:"mandatory,ext"`
	UPSecuritykey     UPSecuritykey               `aper:"mandatory,ext"`
	IEExtensions      *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *SecurityInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(1), &aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("encode optionality bitmap failed: %w", err)
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
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *SecurityInformation) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	optionalityBitmap, _, err := r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 1}, false)
	if err != nil {
		return fmt.Errorf("read optionality bitmap failed: %w", err)
	}
	if err = s.SecurityAlgorithm.Decode(r); err != nil {
		return fmt.Errorf("decode SecurityAlgorithm failed: %w", err)
	}
	if err = s.UPSecuritykey.Decode(r); err != nil {
		return fmt.Errorf("decode UPSecuritykey failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for SecurityInformation */
	}
	return nil
}
