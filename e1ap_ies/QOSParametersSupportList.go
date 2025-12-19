package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// QOSParametersSupportList is a generated SEQUENCE type.
type QOSParametersSupportList struct {
	EUTRANQOSSupportList *EUTRANQOSSupportList       `aper:"lb:1,ub:MaxnoofEUTRANQOSParameters,optional,ext"`
	NGRANQOSSupportList  *NGRANQOSSupportList        `aper:"lb:1,ub:MaxnoofNGRANQOSParameters,optional,ext"`
	IEExtensions         *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *QOSParametersSupportList) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("encode extensibility bool failed: %w", err)
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
		return fmt.Errorf("encode optionality bitmap failed: %w", err)
	}
	if s.EUTRANQOSSupportList != nil {
		if err = s.EUTRANQOSSupportList.Encode(w); err != nil {
			return fmt.Errorf("encode EUTRANQOSSupportList failed: %w", err)
		}
	}
	if s.NGRANQOSSupportList != nil {
		if err = s.NGRANQOSSupportList.Encode(w); err != nil {
			return fmt.Errorf("encode NGRANQOSSupportList failed: %w", err)
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
func (s *QOSParametersSupportList) Decode(r *aper.AperReader) (err error) {
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
		s.EUTRANQOSSupportList = new(EUTRANQOSSupportList)
		if err = s.EUTRANQOSSupportList.Decode(r); err != nil {
			return fmt.Errorf("decode EUTRANQOSSupportList failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.NGRANQOSSupportList = new(NGRANQOSSupportList)
		if err = s.NGRANQOSSupportList.Decode(r); err != nil {
			return fmt.Errorf("decode NGRANQOSSupportList failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for QOSParametersSupportList */
	}
	return nil
}
