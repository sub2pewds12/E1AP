package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// BearerContextModificationConfirm is a generated SEQUENCE type.
type BearerContextModificationConfirm struct {
	GNBCUCPUEE1APID                        GNBCUCPUEE1APID                         `aper:"lb:0,ub:4294967295,mandatory,ext"`
	GNBCUUPUEE1APID                        GNBCUUPUEE1APID                         `aper:"lb:0,ub:4294967295,mandatory,ext"`
	SystemBearerContextModificationConfirm *SystemBearerContextModificationConfirm `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *BearerContextModificationConfirm) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.SystemBearerContextModificationConfirm != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(1), &aper.Constraint{Lb: 1, Ub: 1}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.GNBCUCPUEE1APID.Value), &aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
		return fmt.Errorf("Encode GNBCUCPUEE1APID failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.GNBCUUPUEE1APID.Value), &aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
		return fmt.Errorf("Encode GNBCUUPUEE1APID failed: %w", err)
	}
	if s.SystemBearerContextModificationConfirm != nil {
		if err = s.SystemBearerContextModificationConfirm.Encode(w); err != nil {
			return fmt.Errorf("Encode SystemBearerContextModificationConfirm failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *BearerContextModificationConfirm) Decode(r *aper.AperReader) (err error) {
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
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
			return fmt.Errorf("Decode GNBCUCPUEE1APID failed: %w", err)
		}
		s.GNBCUCPUEE1APID.Value = aper.Integer(val)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 4294967295}, false); err != nil {
			return fmt.Errorf("Decode GNBCUUPUEE1APID failed: %w", err)
		}
		s.GNBCUUPUEE1APID.Value = aper.Integer(val)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.SystemBearerContextModificationConfirm = new(SystemBearerContextModificationConfirm)
		if err = s.SystemBearerContextModificationConfirm.Decode(r); err != nil {
			return fmt.Errorf("Decode SystemBearerContextModificationConfirm failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for BearerContextModificationConfirm")
	}
	return nil
}
