package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// ImmediateMDT is a generated SEQUENCE type.
type ImmediateMDT struct {
	MeasurementsToActivate MeasurementsToActivate      `aper:"lb:8,ub:8,mandatory,ext"`
	MeasurementFour        *M4Configuration            `aper:"optional,ext"`
	MeasurementSix         *M6Configuration            `aper:"optional,ext"`
	MeasurementSeven       *M7Configuration            `aper:"optional,ext"`
	IEExtensions           *ProtocolExtensionContainer `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *ImmediateMDT) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.MeasurementFour != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.MeasurementSix != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.MeasurementSeven != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 4
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(4), &aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return fmt.Errorf("encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteBitString(s.MeasurementsToActivate.Value.Bytes, uint(s.MeasurementsToActivate.Value.NumBits), &aper.Constraint{Lb: 8, Ub: 8}, false); err != nil {
		return fmt.Errorf("encode MeasurementsToActivate failed: %w", err)
	}
	if s.MeasurementFour != nil {
		if err = s.MeasurementFour.Encode(w); err != nil {
			return fmt.Errorf("encode MeasurementFour failed: %w", err)
		}
	}
	if s.MeasurementSix != nil {
		if err = s.MeasurementSix.Encode(w); err != nil {
			return fmt.Errorf("encode MeasurementSix failed: %w", err)
		}
	}
	if s.MeasurementSeven != nil {
		if err = s.MeasurementSeven.Encode(w); err != nil {
			return fmt.Errorf("encode MeasurementSeven failed: %w", err)
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
func (s *ImmediateMDT) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	optionalityBitmap, _, err := r.ReadBitString(&aper.Constraint{Lb: 4, Ub: 4}, false)
	if err != nil {
		return fmt.Errorf("read optionality bitmap failed: %w", err)
	}
	if err = s.MeasurementsToActivate.Decode(r); err != nil {
		return fmt.Errorf("decode MeasurementsToActivate failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.MeasurementFour = new(M4Configuration)
		if err = s.MeasurementFour.Decode(r); err != nil {
			return fmt.Errorf("decode MeasurementFour failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.MeasurementSix = new(M6Configuration)
		if err = s.MeasurementSix.Decode(r); err != nil {
			return fmt.Errorf("decode MeasurementSix failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.MeasurementSeven = new(M7Configuration)
		if err = s.MeasurementSeven.Decode(r); err != nil {
			return fmt.Errorf("decode MeasurementSeven failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for ImmediateMDT */
	}
	return nil
}
