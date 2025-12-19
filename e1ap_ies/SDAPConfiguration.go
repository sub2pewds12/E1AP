package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// SDAPConfiguration is a generated SEQUENCE type.
type SDAPConfiguration struct {
	DefaultDRB   DefaultDRB                  `aper:"mandatory,ext"`
	SDAPHeaderUL SDAPHeaderUL                `aper:"mandatory,ext"`
	SDAPHeaderDL SDAPHeaderDL                `aper:"mandatory,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *SDAPConfiguration) Encode(w *aper.AperWriter) (err error) {
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
	if err = w.WriteEnumerate(uint64(s.DefaultDRB.Value), aper.Constraint{Lb: 0, Ub: 1}, true); err != nil {
		return fmt.Errorf("Encode DefaultDRB failed: %w", err)
	}
	if err = w.WriteEnumerate(uint64(s.SDAPHeaderUL.Value), aper.Constraint{Lb: 0, Ub: 1}, true); err != nil {
		return fmt.Errorf("Encode SDAPHeaderUL failed: %w", err)
	}
	if err = w.WriteEnumerate(uint64(s.SDAPHeaderDL.Value), aper.Constraint{Lb: 0, Ub: 1}, true); err != nil {
		return fmt.Errorf("Encode SDAPHeaderDL failed: %w", err)
	}
	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("Encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *SDAPConfiguration) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if err = s.DefaultDRB.Decode(r); err != nil {
		return fmt.Errorf("Decode DefaultDRB failed: %w", err)
	}
	if err = s.SDAPHeaderUL.Decode(r); err != nil {
		return fmt.Errorf("Decode SDAPHeaderUL failed: %w", err)
	}
	if err = s.SDAPHeaderDL.Decode(r); err != nil {
		return fmt.Errorf("Decode SDAPHeaderDL failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for SDAPConfiguration */
	}
	return nil
}
