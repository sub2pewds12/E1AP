package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// TSCTrafficInformation is a generated SEQUENCE type.
type TSCTrafficInformation struct {
	Periodicity      Periodicity                 `aper:"lb:1,ub:640000,mandatory"`
	BurstArrivalTime *BurstArrivalTime           `aper:"optional"`
	IEExtensions     *ProtocolExtensionContainer `aper:"optional"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *TSCTrafficInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(false); err != nil {
		return fmt.Errorf("encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.BurstArrivalTime != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(2), &aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.Periodicity.Value), &aper.Constraint{Lb: 1, Ub: 640000}, true); err != nil {
		return fmt.Errorf("encode Periodicity failed: %w", err)
	}
	if s.BurstArrivalTime != nil {
		if err = w.WriteOctetString([]byte((*s.BurstArrivalTime).Value), &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return fmt.Errorf("encode BurstArrivalTime failed: %w", err)
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
func (s *TSCTrafficInformation) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	optionalityBitmap, _, err := r.ReadBitString(&aper.Constraint{Lb: 2, Ub: 2}, false)
	if err != nil {
		return fmt.Errorf("read optionality bitmap failed: %w", err)
	}
	if err = s.Periodicity.Decode(r); err != nil {
		return fmt.Errorf("decode Periodicity failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.BurstArrivalTime = new(BurstArrivalTime)
		if err = s.BurstArrivalTime.Decode(r); err != nil {
			return fmt.Errorf("decode BurstArrivalTime failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for TSCTrafficInformation */
	}
	return nil
}
