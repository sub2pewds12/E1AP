package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// M4Configuration is a generated SEQUENCE type.
type M4Configuration struct {
	M4period     M4period                    `aper:"mandatory,ext"`
	M4LinksToLog LinksToLog                  `aper:"mandatory,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *M4Configuration) Encode(w *aper.AperWriter) (err error) {
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
	if err = w.WriteEnumerate(uint64(s.M4period.Value), aper.Constraint{Lb: 0, Ub: 4}, true); err != nil {
		return fmt.Errorf("Encode M4period failed: %w", err)
	}
	if err = w.WriteEnumerate(uint64(s.M4LinksToLog.Value), aper.Constraint{Lb: 0, Ub: 2}, true); err != nil {
		return fmt.Errorf("Encode M4LinksToLog failed: %w", err)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *M4Configuration) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}

	{
		var val uint64
		if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 4}, true); err != nil {
			return fmt.Errorf("Decode M4period failed: %w", err)
		}
		s.M4period.Value = aper.Enumerated(val)
	}

	{
		var val uint64
		if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 2}, true); err != nil {
			return fmt.Errorf("Decode M4LinksToLog failed: %w", err)
		}
		s.M4LinksToLog.Value = aper.Enumerated(val)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for M4Configuration")
	}
	return nil
}
