package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// IABUPTNLAddressUpdate is a generated SEQUENCE type.
type IABUPTNLAddressUpdate struct {
	TransactionID              TransactionID                `aper:"lb:0,ub:255,mandatory,ext"`
	DLUPTNLAddressToUpdateList []DLUPTNLAddressToUpdateItem `aper:"ub:MaxnoofTNLAddresses,optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *IABUPTNLAddressUpdate) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.DLUPTNLAddressToUpdateList != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(1), &aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.TransactionID.Value), &aper.Constraint{Lb: 0, Ub: 255}, true); err != nil {
		return fmt.Errorf("Encode TransactionID failed: %w", err)
	}
	if s.DLUPTNLAddressToUpdateList != nil {
		if err = s.DLUPTNLAddressToUpdateList.Encode(w); err != nil {
			return fmt.Errorf("Encode DLUPTNLAddressToUpdateList failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *IABUPTNLAddressUpdate) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, true); err != nil {
			return fmt.Errorf("Decode TransactionID failed: %w", err)
		}
		s.TransactionID.Value = aper.Integer(val)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.DLUPTNLAddressToUpdateList = new(DLUPTNLAddressToUpdateList)
		if err = s.DLUPTNLAddressToUpdateList.Decode(r); err != nil {
			return fmt.Errorf("Decode DLUPTNLAddressToUpdateList failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for IABUPTNLAddressUpdate")
	}
	return nil
}
