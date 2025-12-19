package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// TSCTrafficCharacteristics is a generated SEQUENCE type.
type TSCTrafficCharacteristics struct {
	TSCTrafficCharacteristicsUL *TSCTrafficInformation      `aper:"optional"`
	TSCTrafficCharacteristicsDL *TSCTrafficInformation      `aper:"optional"`
	IEExtensions                *ProtocolExtensionContainer `aper:"optional"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *TSCTrafficCharacteristics) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(false); err != nil {
		return fmt.Errorf("encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.TSCTrafficCharacteristicsUL != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.TSCTrafficCharacteristicsDL != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(3), &aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return fmt.Errorf("encode optionality bitmap failed: %w", err)
	}
	if s.TSCTrafficCharacteristicsUL != nil {
		if err = s.TSCTrafficCharacteristicsUL.Encode(w); err != nil {
			return fmt.Errorf("encode TSCTrafficCharacteristicsUL failed: %w", err)
		}
	}
	if s.TSCTrafficCharacteristicsDL != nil {
		if err = s.TSCTrafficCharacteristicsDL.Encode(w); err != nil {
			return fmt.Errorf("encode TSCTrafficCharacteristicsDL failed: %w", err)
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
func (s *TSCTrafficCharacteristics) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	optionalityBitmap, _, err := r.ReadBitString(&aper.Constraint{Lb: 3, Ub: 3}, false)
	if err != nil {
		return fmt.Errorf("read optionality bitmap failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.TSCTrafficCharacteristicsUL = new(TSCTrafficInformation)
		if err = s.TSCTrafficCharacteristicsUL.Decode(r); err != nil {
			return fmt.Errorf("decode TSCTrafficCharacteristicsUL failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.TSCTrafficCharacteristicsDL = new(TSCTrafficInformation)
		if err = s.TSCTrafficCharacteristicsDL.Decode(r); err != nil {
			return fmt.Errorf("decode TSCTrafficCharacteristicsDL failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for TSCTrafficCharacteristics */
	}
	return nil
}
