package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// EHCDownlinkParameters is a generated SEQUENCE type.
type EHCDownlinkParameters struct {
	DRBContinueEHCDL EHCDownlinkParametersDRBContinueEHCDL `aper:"mandatory,ext"`
	IEExtensions     *EHCDownlinkParametersExtensions      `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *EHCDownlinkParameters) Encode(w *aper.AperWriter) (err error) {
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
	if err = s.DRBContinueEHCDL.Encode(w); err != nil {
		return fmt.Errorf("Encode DRBContinueEHCDL failed: %w", err)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *EHCDownlinkParameters) Decode(r *aper.AperReader) (err error) {
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
		if val, err = r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 0}, true); err != nil {
			return fmt.Errorf("Decode DRBContinueEHCDL failed: %w", err)
		}
		s.DRBContinueEHCDL.Value = aper.Enumerated(val)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.IEExtensions = new(EHCDownlinkParametersExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for EHCDownlinkParameters")
	}
	return nil
}
