package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// ROHC is a generated SEQUENCE type.
type ROHC struct {
	MaxCID       ROHCMaxCID                  `aper:"lb:0,ub:16383,mandatory,ext"`
	ROHCProfiles ROHCROHCProfiles            `aper:"lb:0,ub:511,mandatory,ext"`
	ContinueROHC *ROHCContinueROHC           `aper:"optional,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *ROHC) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.ContinueROHC != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(2), &aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.MaxCID.Value), &aper.Constraint{Lb: 0, Ub: 16383}, true); err != nil {
		return fmt.Errorf("Encode MaxCID failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.ROHCProfiles.Value), &aper.Constraint{Lb: 0, Ub: 511}, true); err != nil {
		return fmt.Errorf("Encode ROHCProfiles failed: %w", err)
	}
	if s.ContinueROHC != nil {
		if err = w.WriteEnumerate(uint64(s.ContinueROHC.Value), aper.Constraint{Lb: 0, Ub: 0}, true); err != nil {
			return fmt.Errorf("Encode ContinueROHC failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ROHC) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 16383}, true); err != nil {
			return fmt.Errorf("Decode MaxCID failed: %w", err)
		}
		s.MaxCID.Value = aper.Integer(val)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 511}, true); err != nil {
			return fmt.Errorf("Decode ROHCProfiles failed: %w", err)
		}
		s.ROHCProfiles.Value = aper.Integer(val)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.ContinueROHC = new(ROHCContinueROHC)

		{
			var val uint64
			if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true); err != nil {
				return fmt.Errorf("Decode ContinueROHC failed: %w", err)
			}
			s.ContinueROHC.Value = aper.Enumerated(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for ROHC")
	}
	return nil
}
