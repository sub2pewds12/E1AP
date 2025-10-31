package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// ExtendedGNBCUUPName is a generated SEQUENCE type.
type ExtendedGNBCUUPName struct {
	GNBCUUPNameVisibleString *aper.OctetString           `aper:"optional,ext"`
	GNBCUUPNameUTF8String    *aper.OctetString           `aper:"optional,ext"`
	IEExtensions             *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *ExtendedGNBCUUPName) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
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
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if s.GNBCUUPNameVisibleString != nil {
		if err = s.GNBCUUPNameVisibleString.Encode(w); err != nil {
			return fmt.Errorf("Encode GNBCUUPNameVisibleString failed: %w", err)
		}
	}
	if s.GNBCUUPNameUTF8String != nil {
		if err = s.GNBCUUPNameUTF8String.Encode(w); err != nil {
			return fmt.Errorf("Encode GNBCUUPNameUTF8String failed: %w", err)
		}
	}
	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("Encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ExtendedGNBCUUPName) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {

		{
			var val []byte
			if val, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
				return fmt.Errorf("Decode GNBCUUPNameVisibleString failed: %w", err)
			}
			tmp := aper.OctetString(val)
			s.GNBCUUPNameVisibleString = &tmp
		}

	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {

		{
			var val []byte
			if val, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
				return fmt.Errorf("Decode GNBCUUPNameUTF8String failed: %w", err)
			}
			tmp := aper.OctetString(val)
			s.GNBCUUPNameUTF8String = &tmp
		}

	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for ExtendedGNBCUUPName")
	}
	return nil
}
