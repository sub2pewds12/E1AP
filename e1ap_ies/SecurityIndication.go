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
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.MaximumIPdatarate != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(2), &aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = s.IntegrityProtectionIndication.Encode(w); err != nil {
		return fmt.Errorf("Encode IntegrityProtectionIndication failed: %w", err)
	}
	if err = s.ConfidentialityProtectionIndication.Encode(w); err != nil {
		return fmt.Errorf("Encode ConfidentialityProtectionIndication failed: %w", err)
	}
	if s.MaximumIPdatarate != nil {
		if err = s.MaximumIPdatarate.Encode(w); err != nil {
			return fmt.Errorf("Encode MaximumIPdatarate failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *SecurityIndication) Decode(r *aper.AperReader) (err error) {
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
		if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true); err != nil {
			return fmt.Errorf("Decode IntegrityProtectionIndication failed: %w", err)
		}
		s.IntegrityProtectionIndication.Value = aper.Enumerated(val)
	}

	{
		var val uint64
		if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true); err != nil {
			return fmt.Errorf("Decode ConfidentialityProtectionIndication failed: %w", err)
		}
		s.ConfidentialityProtectionIndication.Value = aper.Enumerated(val)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.MaximumIPdatarate = new(MaximumIPdatarate)
		if err = s.MaximumIPdatarate.Decode(r); err != nil {
			return fmt.Errorf("Decode MaximumIPdatarate failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for SecurityIndication")
	}
	return nil
}
