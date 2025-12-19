package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// ExtendedGNBCUUPName is a generated SEQUENCE type.
type ExtendedGNBCUUPName struct {
	GNBCUUPNameVisibleString *GNBCUUPNameVisibleString   `aper:"optional,ext"`
	GNBCUUPNameUTF8String    *GNBCUUPNameUTF8String      `aper:"optional,ext"`
	IEExtensions             *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *ExtendedGNBCUUPName) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.GNBCUUPNameVisibleString != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.GNBCUUPNameUTF8String != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(3), &aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return fmt.Errorf("encode optionality bitmap failed: %w", err)
	}
	if s.GNBCUUPNameVisibleString != nil {
		if err = s.GNBCUUPNameVisibleString.Encode(w); err != nil {
			return fmt.Errorf("encode GNBCUUPNameVisibleString failed: %w", err)
		}
	}
	if s.GNBCUUPNameUTF8String != nil {
		if err = s.GNBCUUPNameUTF8String.Encode(w); err != nil {
			return fmt.Errorf("encode GNBCUUPNameUTF8String failed: %w", err)
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
func (s *ExtendedGNBCUUPName) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	optionalityBitmap, _, err := r.ReadBitString(&aper.Constraint{Lb: 3, Ub: 3}, false)
	if err != nil {
		return fmt.Errorf("read optionality bitmap failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.GNBCUUPNameVisibleString = new(GNBCUUPNameVisibleString)
		if err = s.GNBCUUPNameVisibleString.Decode(r); err != nil {
			return fmt.Errorf("decode GNBCUUPNameVisibleString failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.GNBCUUPNameUTF8String = new(GNBCUUPNameUTF8String)
		if err = s.GNBCUUPNameUTF8String.Decode(r); err != nil {
			return fmt.Errorf("decode GNBCUUPNameUTF8String failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for ExtendedGNBCUUPName */
	}
	return nil
}
