package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// PDUSessionResourceRequiredToModifyItem is a generated SEQUENCE type.
type PDUSessionResourceRequiredToModifyItem struct {
	PDUSessionID                 PDUSessionID                                      `aper:"lb:0,ub:255,mandatory,ext"`
	NGDLUPTNLInformation         *UPTNLInformation                                 `aper:"optional,ext"`
	DRBRequiredToModifyListNGRAN *DRBRequiredToModifyListNGRAN                     `aper:"lb:1,ub:MaxnoofDRBs,optional,ext"`
	DRBRequiredToRemoveListNGRAN *DRBRequiredToRemoveListNGRAN                     `aper:"lb:1,ub:MaxnoofDRBs,optional,ext"`
	IEExtensions                 *PDUSessionResourceRequiredToModifyItemExtensions `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDUSessionResourceRequiredToModifyItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("encode extensibility bool failed: %w", err)
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
		return fmt.Errorf("encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.PDUSessionID.Value), &aper.Constraint{Lb: 0, Ub: 255}, false); err != nil {
		return fmt.Errorf("encode PDUSessionID failed: %w", err)
	}
	if s.NGDLUPTNLInformation != nil {
		if err = s.NGDLUPTNLInformation.Encode(w); err != nil {
			return fmt.Errorf("encode NGDLUPTNLInformation failed: %w", err)
		}
	}
	if s.DRBRequiredToModifyListNGRAN != nil {
		if err = s.DRBRequiredToModifyListNGRAN.Encode(w); err != nil {
			return fmt.Errorf("encode DRBRequiredToModifyListNGRAN failed: %w", err)
		}
	}
	if s.DRBRequiredToRemoveListNGRAN != nil {
		if err = s.DRBRequiredToRemoveListNGRAN.Encode(w); err != nil {
			return fmt.Errorf("encode DRBRequiredToRemoveListNGRAN failed: %w", err)
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
func (s *PDUSessionResourceRequiredToModifyItem) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	optionalityBitmap, _, err := r.ReadBitString(&aper.Constraint{Lb: 4, Ub: 4}, false)
	if err != nil {
		return fmt.Errorf("read optionality bitmap failed: %w", err)
	}
	if err = s.PDUSessionID.Decode(r); err != nil {
		return fmt.Errorf("decode PDUSessionID failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.NGDLUPTNLInformation = new(UPTNLInformation)
		if err = s.NGDLUPTNLInformation.Decode(r); err != nil {
			return fmt.Errorf("decode NGDLUPTNLInformation failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.DRBRequiredToModifyListNGRAN = new(DRBRequiredToModifyListNGRAN)
		if err = s.DRBRequiredToModifyListNGRAN.Decode(r); err != nil {
			return fmt.Errorf("decode DRBRequiredToModifyListNGRAN failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.DRBRequiredToRemoveListNGRAN = new(DRBRequiredToRemoveListNGRAN)
		if err = s.DRBRequiredToRemoveListNGRAN.Decode(r); err != nil {
			return fmt.Errorf("decode DRBRequiredToRemoveListNGRAN failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {
		s.IEExtensions = new(PDUSessionResourceRequiredToModifyItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for PDUSessionResourceRequiredToModifyItem */
	}
	return nil
}
