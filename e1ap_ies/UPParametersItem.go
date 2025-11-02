package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// UPParametersItem is a generated SEQUENCE type.
type UPParametersItem struct {
	UPTNLInformation UPTNLInformation            `aper:"mandatory,ext"`
	CellGroupID      CellGroupID                 `aper:"lb:0,ub:3,mandatory,ext"`
	IEExtensions     *UPParametersItemExtensions `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *UPParametersItem) Encode(w *aper.AperWriter) (err error) {
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
	if err = s.UPTNLInformation.Encode(w); err != nil {
		return fmt.Errorf("Encode UPTNLInformation failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.CellGroupID.Value), &aper.Constraint{Lb: 0, Ub: 3}, true); err != nil {
		return fmt.Errorf("Encode CellGroupID failed: %w", err)
	}
	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("Encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *UPParametersItem) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if err = s.UPTNLInformation.Decode(r); err != nil {
		return fmt.Errorf("Decode UPTNLInformation failed: %w", err)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 3}, true); err != nil {
			return fmt.Errorf("Decode CellGroupID failed: %w", err)
		}
		s.CellGroupID.Value = aper.Integer(val)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.IEExtensions = new(UPParametersItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for UPParametersItem")
	}
	return nil
}
