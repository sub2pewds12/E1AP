package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// ExtendedGNBCUCPName is a generated SEQUENCE type.
type ExtendedGNBCUCPName struct {
	GNBCUCPNameVisibleString *GNBCUCPNameVisibleString   `aper:"optional,ext"`
	GNBCUCPNameUTF8String    *GNBCUCPNameUTF8String      `aper:"optional,ext"`
	IEExtensions             *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *ExtendedGNBCUCPName) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.GNBCUCPNameVisibleString != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.GNBCUCPNameUTF8String != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(3), &aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if s.GNBCUCPNameVisibleString != nil {
		if err = w.WriteOctetString([]byte(s.GNBCUCPNameVisibleString.Value), &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return fmt.Errorf("Encode GNBCUCPNameVisibleString failed: %w", err)
		}
	}
	if s.GNBCUCPNameUTF8String != nil {
		if err = w.WriteOctetString([]byte(s.GNBCUCPNameUTF8String.Value), &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return fmt.Errorf("Encode GNBCUCPNameUTF8String failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ExtendedGNBCUCPName) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {

		var val []byte
		if val, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return fmt.Errorf("Decode GNBCUCPNameVisibleString failed: %w", err)
		}
		s.GNBCUCPNameVisibleString = new(GNBCUCPNameVisibleString)
		s.GNBCUCPNameVisibleString.Value = aper.OctetString(val)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {

		var val []byte
		if val, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return fmt.Errorf("Decode GNBCUCPNameUTF8String failed: %w", err)
		}
		s.GNBCUCPNameUTF8String = new(GNBCUCPNameUTF8String)
		s.GNBCUCPNameUTF8String.Value = aper.OctetString(val)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for ExtendedGNBCUCPName")
	}
	return nil
}
