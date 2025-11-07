package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// PDUSessionResourceRequiredToModifyItem is a generated SEQUENCE type.
type PDUSessionResourceRequiredToModifyItem struct {
	PDUSessionID                 PDUSessionID                                      `aper:"lb:0,ub:255,mandatory,ext"`
	NGDLUPTNLInformation         *UPTNLInformation                                 `aper:"optional,ext"`
	DRBRequiredToModifyListNGRAN *DRBRequiredToModifyListNGRAN                     `aper:"optional,ext"`
	DRBRequiredToRemoveListNGRAN *DRBRequiredToRemoveListNGRAN                     `aper:"optional,ext"`
	IEExtensions                 *PDUSessionResourceRequiredToModifyItemExtensions `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDUSessionResourceRequiredToModifyItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.NGDLUPTNLInformation != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.DRBRequiredToModifyListNGRAN != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.DRBRequiredToRemoveListNGRAN != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 4
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(4), &aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.PDUSessionID.Value), &aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
		return fmt.Errorf("Encode PDUSessionID failed: %w", err)
	}
	if s.NGDLUPTNLInformation != nil {
		if err = s.NGDLUPTNLInformation.Encode(w); err != nil {
			return fmt.Errorf("Encode NGDLUPTNLInformation failed: %w", err)
		}
	}
	if s.DRBRequiredToModifyListNGRAN != nil {
		{
			itemPointers := make([]aper.AperMarshaller, len(s.DRBRequiredToModifyListNGRAN.Value))
			for i := 0; i < len(s.DRBRequiredToModifyListNGRAN.Value); i++ {
				itemPointers[i] = &(s.DRBRequiredToModifyListNGRAN.Value[i])
			}
			if err = aper.WriteSequenceOf(itemPointers, w, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
				return fmt.Errorf("Encode DRBRequiredToModifyListNGRAN failed: %w", err)
			}
		}
	}
	if s.DRBRequiredToRemoveListNGRAN != nil {
		{
			itemPointers := make([]aper.AperMarshaller, len(s.DRBRequiredToRemoveListNGRAN.Value))
			for i := 0; i < len(s.DRBRequiredToRemoveListNGRAN.Value); i++ {
				itemPointers[i] = &(s.DRBRequiredToRemoveListNGRAN.Value[i])
			}
			if err = aper.WriteSequenceOf(itemPointers, w, &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
				return fmt.Errorf("Encode DRBRequiredToRemoveListNGRAN failed: %w", err)
			}
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *PDUSessionResourceRequiredToModifyItem) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if err = s.PDUSessionID.Decode(r); err != nil {
		return fmt.Errorf("Decode PDUSessionID failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.NGDLUPTNLInformation = new(UPTNLInformation)
		if err = s.NGDLUPTNLInformation.Decode(r); err != nil {
			return fmt.Errorf("Decode NGDLUPTNLInformation failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.DRBRequiredToModifyListNGRAN = new(DRBRequiredToModifyListNGRAN)
		if err = s.DRBRequiredToModifyListNGRAN.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBRequiredToModifyListNGRAN failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.DRBRequiredToRemoveListNGRAN = new(DRBRequiredToRemoveListNGRAN)
		if err = s.DRBRequiredToRemoveListNGRAN.Decode(r); err != nil {
			return fmt.Errorf("Decode DRBRequiredToRemoveListNGRAN failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {
		s.IEExtensions = new(PDUSessionResourceRequiredToModifyItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for PDUSessionResourceRequiredToModifyItem")
	}
	return nil
}
