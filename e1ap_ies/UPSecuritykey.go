package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// UPSecuritykey is a generated SEQUENCE type.
type UPSecuritykey struct {
	EncryptionKey          EncryptionKey               `aper:"mandatory,ext"`
	IntegrityProtectionKey *IntegrityProtectionKey     `aper:"optional,ext"`
	IEExtensions           *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *UPSecuritykey) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.IntegrityProtectionKey != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(2), &aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteOctetString([]byte(s.EncryptionKey.Value), &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
		return fmt.Errorf("encode EncryptionKey failed: %w", err)
	}
	if s.IntegrityProtectionKey != nil {
		if err = w.WriteOctetString([]byte((*s.IntegrityProtectionKey).Value), &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return fmt.Errorf("encode IntegrityProtectionKey failed: %w", err)
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
func (s *UPSecuritykey) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	optionalityBitmap, _, err := r.ReadBitString(&aper.Constraint{Lb: 2, Ub: 2}, false)
	if err != nil {
		return fmt.Errorf("read optionality bitmap failed: %w", err)
	}
	if err = s.EncryptionKey.Decode(r); err != nil {
		return fmt.Errorf("decode EncryptionKey failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.IntegrityProtectionKey = new(IntegrityProtectionKey)
		if err = s.IntegrityProtectionKey.Decode(r); err != nil {
			return fmt.Errorf("decode IntegrityProtectionKey failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for UPSecuritykey */
	}
	return nil
}
