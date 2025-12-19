package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// SecurityIndication is a generated SEQUENCE type.
type SecurityIndication struct {
	IntegrityProtectionIndication       IntegrityProtectionIndication       `aper:"mandatory,ext"`
	ConfidentialityProtectionIndication ConfidentialityProtectionIndication `aper:"mandatory,ext"`
	MaximumIPdatarate                   *MaximumIPdatarate                  `aper:"optional,ext"`
	IEExtensions                        *ProtocolExtensionContainer         `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *SecurityIndication) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.MaximumIPdatarate != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(2), &aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteEnumerate(uint64(s.IntegrityProtectionIndication.Value), aper.Constraint{Lb: 0, Ub: 2}, true); err != nil {
		return fmt.Errorf("encode IntegrityProtectionIndication failed: %w", err)
	}
	if err = w.WriteEnumerate(uint64(s.ConfidentialityProtectionIndication.Value), aper.Constraint{Lb: 0, Ub: 2}, true); err != nil {
		return fmt.Errorf("encode ConfidentialityProtectionIndication failed: %w", err)
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
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *SecurityIndication) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	optionalityBitmap, _, err := r.ReadBitString(&aper.Constraint{Lb: 2, Ub: 2}, false)
	if err != nil {
		return fmt.Errorf("read optionality bitmap failed: %w", err)
	}
	if err = s.IntegrityProtectionIndication.Decode(r); err != nil {
		return fmt.Errorf("decode IntegrityProtectionIndication failed: %w", err)
	}
	if err = s.ConfidentialityProtectionIndication.Decode(r); err != nil {
		return fmt.Errorf("decode ConfidentialityProtectionIndication failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.MaximumIPdatarate = new(MaximumIPdatarate)
		if err = s.MaximumIPdatarate.Decode(r); err != nil {
			return fmt.Errorf("decode MaximumIPdatarate failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for SecurityIndication */
	}
	return nil
}
