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
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.BurstArrivalTime != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(2), &aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.Periodicity), &aper.Constraint{Lb: 1, Ub: 640000}, true); err != nil {
		return fmt.Errorf("Encode Periodicity failed: %w", err)
	}
	if s.BurstArrivalTime != nil {
		if err = w.WriteOctetString([]byte((*s.BurstArrivalTime)), &aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
			return fmt.Errorf("Encode BurstArrivalTime failed: %w", err)
		}
	}
	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("Encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *TSCTrafficInformation) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 640000}, true); err != nil {
			return fmt.Errorf("Decode Periodicity failed: %w", err)
		}
		s.Periodicity = Periodicity(val)
	}

	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {

		{
			var val []byte
			if val, err = r.ReadOctetString(&aper.Constraint{Lb: 0, Ub: 0}, false); err != nil {
				return fmt.Errorf("Decode BurstArrivalTime failed: %w", err)
			}
			tmp := BurstArrivalTime(val)
			s.BurstArrivalTime = &tmp
		}

	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for TSCTrafficInformation")
	}
	return nil
}
