package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// QOSParametersSupportList is a generated SEQUENCE type.
type QOSParametersSupportList struct {
	EUTRANQOSSupportList []EUTRANQOSSupportItem      `aper:"optional,ext"`
	NGRANQOSSupportList  []NGRANQOSSupportItem       `aper:"optional,ext"`
	IEExtensions         *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *QOSParametersSupportList) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.EUTRANQOSSupportList != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.NGRANQOSSupportList != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(3), &aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if s.EUTRANQOSSupportList != nil {
		if err = s.EUTRANQOSSupportList.Encode(w); err != nil {
			return fmt.Errorf("Encode EUTRANQOSSupportList failed: %w", err)
		}
	}
	if s.NGRANQOSSupportList != nil {
		if err = s.NGRANQOSSupportList.Encode(w); err != nil {
			return fmt.Errorf("Encode NGRANQOSSupportList failed: %w", err)
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
func (s *QOSParametersSupportList) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.EUTRANQOSSupportList = new(EUTRANQOSSupportList)
		if err = s.EUTRANQOSSupportList.Decode(r); err != nil {
			return fmt.Errorf("Decode EUTRANQOSSupportList failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.NGRANQOSSupportList = new(NGRANQOSSupportList)
		if err = s.NGRANQOSSupportList.Decode(r); err != nil {
			return fmt.Errorf("Decode NGRANQOSSupportList failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for QOSParametersSupportList")
	}
	return nil
}
