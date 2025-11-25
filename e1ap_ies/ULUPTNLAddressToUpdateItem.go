package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// ULUPTNLAddressToUpdateItem is a generated SEQUENCE type.
type ULUPTNLAddressToUpdateItem struct {
	OldTNLAdress TransportLayerAddress       `aper:"lb:1,ub:160,mandatory,ext"`
	NewTNLAdress TransportLayerAddress       `aper:"lb:1,ub:160,mandatory,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *ULUPTNLAddressToUpdateItem) Encode(w *aper.AperWriter) (err error) {
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
	if err = w.WriteBitString(s.OldTNLAdress.Value.Bytes, uint(s.OldTNLAdress.Value.NumBits), &aper.Constraint{Lb: 1, Ub: 160}, false); err != nil {
		return fmt.Errorf("Encode OldTNLAdress failed: %w", err)
	}
	if err = w.WriteBitString(s.NewTNLAdress.Value.Bytes, uint(s.NewTNLAdress.Value.NumBits), &aper.Constraint{Lb: 1, Ub: 160}, false); err != nil {
		return fmt.Errorf("Encode NewTNLAdress failed: %w", err)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ULUPTNLAddressToUpdateItem) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if err = s.OldTNLAdress.Decode(r); err != nil {
		return fmt.Errorf("Decode OldTNLAdress failed: %w", err)
	}
	if err = s.NewTNLAdress.Decode(r); err != nil {
		return fmt.Errorf("Decode NewTNLAdress failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for ULUPTNLAddressToUpdateItem */
	}
	return nil
}
