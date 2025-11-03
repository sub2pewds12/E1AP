package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// MaximumIPdatarate is a generated SEQUENCE type.
type MaximumIPdatarate struct {
	MaxIPrate    MaxIPrate                   `aper:"mandatory,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *MaximumIPdatarate) Encode(w *aper.AperWriter) (err error) {
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
	if err = w.WriteEnumerate(uint64(s.MaxIPrate.Value), aper.Constraint{Lb: 0, Ub: 1}, true); err != nil {
		return fmt.Errorf("Encode MaxIPrate failed: %w", err)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *MaximumIPdatarate) Decode(r *aper.AperReader) (err error) {
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
		if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 1}, true); err != nil {
			return fmt.Errorf("Decode MaxIPrate failed: %w", err)
		}
		s.MaxIPrate.Value = aper.Enumerated(val)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for MaximumIPdatarate")
	}
	return nil
}
