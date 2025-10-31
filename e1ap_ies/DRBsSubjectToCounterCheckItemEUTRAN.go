package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DRBsSubjectToCounterCheckItemEUTRAN is a generated SEQUENCE type.
type DRBsSubjectToCounterCheckItemEUTRAN struct {
	DRBID        DRBID                       `aper:"lb:1,ub:32,mandatory,ext"`
	PDCPULCount  PDCPCount                   `aper:"mandatory,ext"`
	PDCPDLCount  PDCPCount                   `aper:"mandatory,ext"`
	IEExtensions *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBsSubjectToCounterCheckItemEUTRAN) Encode(w *aper.AperWriter) (err error) {
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
	if err = w.WriteInteger(int64(s.DRBID), &aper.Constraint{Lb: 1, Ub: 32}, true); err != nil {
		return fmt.Errorf("Encode DRBID failed: %w", err)
	}
	if err = s.PDCPULCount.Encode(w); err != nil {
		return fmt.Errorf("Encode PDCPULCount failed: %w", err)
	}
	if err = s.PDCPDLCount.Encode(w); err != nil {
		return fmt.Errorf("Encode PDCPDLCount failed: %w", err)
	}
	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("Encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DRBsSubjectToCounterCheckItemEUTRAN) Decode(r *aper.AperReader) (err error) {
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
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 32}, true); err != nil {
			return fmt.Errorf("Decode DRBID failed: %w", err)
		}
		s.DRBID = DRBID(val)
	}

	if err = s.PDCPULCount.Decode(r); err != nil {
		return fmt.Errorf("Decode PDCPULCount failed: %w", err)
	}
	if err = s.PDCPDLCount.Decode(r); err != nil {
		return fmt.Errorf("Decode PDCPDLCount failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for DRBsSubjectToCounterCheckItemEUTRAN")
	}
	return nil
}
