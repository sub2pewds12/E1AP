package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// SecurityAlgorithm is a generated SEQUENCE type.
type SecurityAlgorithm struct {
	CipheringAlgorithm           CipheringAlgorithm            `aper:"mandatory,ext"`
	IntegrityProtectionAlgorithm *IntegrityProtectionAlgorithm `aper:"optional,ext"`
	IEExtensions                 *ProtocolExtensionContainer   `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *SecurityAlgorithm) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.IntegrityProtectionAlgorithm != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(2), &aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteEnumerate(uint64(s.CipheringAlgorithm.Value), aper.Constraint{Lb: 0, Ub: 3}, true); err != nil {
		return fmt.Errorf("Encode CipheringAlgorithm failed: %w", err)
	}
	if s.IntegrityProtectionAlgorithm != nil {
		if err = w.WriteEnumerate(uint64(s.IntegrityProtectionAlgorithm.Value), aper.Constraint{Lb: 0, Ub: 3}, true); err != nil {
			return fmt.Errorf("Encode IntegrityProtectionAlgorithm failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *SecurityAlgorithm) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}

	{
		var val uint64
		if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, true); err != nil {
			return fmt.Errorf("Decode CipheringAlgorithm failed: %w", err)
		}
		s.CipheringAlgorithm.Value = aper.Enumerated(val)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.IntegrityProtectionAlgorithm = new(IntegrityProtectionAlgorithm)

		{
			var val uint64
			if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 3}, true); err != nil {
				return fmt.Errorf("Decode IntegrityProtectionAlgorithm failed: %w", err)
			}
			s.IntegrityProtectionAlgorithm.Value = aper.Enumerated(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for SecurityAlgorithm")
	}
	return nil
}
