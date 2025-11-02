package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// GBRQosInformation is a generated SEQUENCE type.
type GBRQosInformation struct {
	ERABMaximumBitrateDL    *BitRate                    `aper:"lb:0,ub:4000000000000,optional,ext"`
	ERABMaximumBitrateUL    *BitRate                    `aper:"lb:0,ub:4000000000000,optional,ext"`
	ERABGuaranteedBitrateDL *BitRate                    `aper:"lb:0,ub:4000000000000,optional,ext"`
	ERABGuaranteedBitrateUL *BitRate                    `aper:"lb:0,ub:4000000000000,optional,ext"`
	IEExtensions            *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *GBRQosInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.ERABMaximumBitrateDL != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.ERABMaximumBitrateUL != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.ERABGuaranteedBitrateDL != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if s.ERABGuaranteedBitrateUL != nil {
		optionalityBitmap[0] |= 1 << 4
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 3
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(5), &aper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if s.ERABMaximumBitrateDL != nil {
		if err = w.WriteInteger(int64(s.ERABMaximumBitrateDL.Value), &aper.Constraint{Lb: 0, Ub: 4000000000000}, true); err != nil {
			return fmt.Errorf("Encode ERABMaximumBitrateDL failed: %w", err)
		}
	}
	if s.ERABMaximumBitrateUL != nil {
		if err = w.WriteInteger(int64(s.ERABMaximumBitrateUL.Value), &aper.Constraint{Lb: 0, Ub: 4000000000000}, true); err != nil {
			return fmt.Errorf("Encode ERABMaximumBitrateUL failed: %w", err)
		}
	}
	if s.ERABGuaranteedBitrateDL != nil {
		if err = w.WriteInteger(int64(s.ERABGuaranteedBitrateDL.Value), &aper.Constraint{Lb: 0, Ub: 4000000000000}, true); err != nil {
			return fmt.Errorf("Encode ERABGuaranteedBitrateDL failed: %w", err)
		}
	}
	if s.ERABGuaranteedBitrateUL != nil {
		if err = w.WriteInteger(int64(s.ERABGuaranteedBitrateUL.Value), &aper.Constraint{Lb: 0, Ub: 4000000000000}, true); err != nil {
			return fmt.Errorf("Encode ERABGuaranteedBitrateUL failed: %w", err)
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
func (s *GBRQosInformation) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 5, Ub: 5}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4000000000000}, true); err != nil {
				return fmt.Errorf("Decode ERABMaximumBitrateDL failed: %w", err)
			}
			s.ERABMaximumBitrateDL = new(BitRate)
			s.ERABMaximumBitrateDL.Value = aper.Integer(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4000000000000}, true); err != nil {
				return fmt.Errorf("Decode ERABMaximumBitrateUL failed: %w", err)
			}
			s.ERABMaximumBitrateUL = new(BitRate)
			s.ERABMaximumBitrateUL.Value = aper.Integer(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4000000000000}, true); err != nil {
				return fmt.Errorf("Decode ERABGuaranteedBitrateDL failed: %w", err)
			}
			s.ERABGuaranteedBitrateDL = new(BitRate)
			s.ERABGuaranteedBitrateDL.Value = aper.Integer(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4000000000000}, true); err != nil {
				return fmt.Errorf("Decode ERABGuaranteedBitrateUL failed: %w", err)
			}
			s.ERABGuaranteedBitrateUL = new(BitRate)
			s.ERABGuaranteedBitrateUL.Value = aper.Integer(val)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<3) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for GBRQosInformation")
	}
	return nil
}
