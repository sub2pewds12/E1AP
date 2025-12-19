package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DRBMeasurementResultsInformationItem is a generated SEQUENCE type.
type DRBMeasurementResultsInformationItem struct {
	DRBID        DRBID                                           `aper:"lb:1,ub:32,mandatory,ext"`
	ULD1Result   *DRBMeasurementResultsInformationItemULD1Result `aper:"lb:0,ub:10000,optional,ext"`
	IEExtensions *ProtocolExtensionContainer                     `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBMeasurementResultsInformationItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.ULD1Result != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(2), &aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.DRBID.Value), &aper.Constraint{Lb: 1, Ub: 32}, true); err != nil {
		return fmt.Errorf("encode DRBID failed: %w", err)
	}
	if s.ULD1Result != nil {
		if err = w.WriteInteger(int64((*s.ULD1Result).Value), &aper.Constraint{Lb: 0, Ub: 10000}, true); err != nil {
			return fmt.Errorf("encode ULD1Result failed: %w", err)
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
func (s *DRBMeasurementResultsInformationItem) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	optionalityBitmap, _, err := r.ReadBitString(&aper.Constraint{Lb: 2, Ub: 2}, false)
	if err != nil {
		return fmt.Errorf("read optionality bitmap failed: %w", err)
	}
	if err = s.DRBID.Decode(r); err != nil {
		return fmt.Errorf("decode DRBID failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.ULD1Result = new(DRBMeasurementResultsInformationItemULD1Result)
		if err = s.ULD1Result.Decode(r); err != nil {
			return fmt.Errorf("decode ULD1Result failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for DRBMeasurementResultsInformationItem */
	}
	return nil
}
