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
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.IntegrityProtectionKey != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(2), &aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = s.EncryptionKey.Encode(w); err != nil {
		return fmt.Errorf("Encode EncryptionKey failed: %w", err)
	}
	if s.IntegrityProtectionKey != nil {
		if err = s.IntegrityProtectionKey.Encode(w); err != nil {
			return fmt.Errorf("Encode IntegrityProtectionKey failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *UPSecuritykey) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if err = s.EncryptionKey.Decode(r); err != nil {
		return fmt.Errorf("Decode EncryptionKey failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.IntegrityProtectionKey = new(IntegrityProtectionKey)
		if err = s.IntegrityProtectionKey.Decode(r); err != nil {
			return fmt.Errorf("Decode IntegrityProtectionKey failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for UPSecuritykey")
	}
	return nil
}
