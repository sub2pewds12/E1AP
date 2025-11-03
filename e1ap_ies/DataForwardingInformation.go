package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DataForwardingInformation is a generated SEQUENCE type.
type DataForwardingInformation struct {
	ULDataForwarding *UPTNLInformation                    `aper:"optional,ext"`
	DLDataForwarding *UPTNLInformation                    `aper:"optional,ext"`
	IEExtensions     *DataForwardingInformationExtensions `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *DataForwardingInformation) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.ULDataForwarding != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.DLDataForwarding != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(3), &aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if s.ULDataForwarding != nil {
		if err = s.ULDataForwarding.Encode(w); err != nil {
			return fmt.Errorf("Encode ULDataForwarding failed: %w", err)
		}
	}
	if s.DLDataForwarding != nil {
		if err = s.DLDataForwarding.Encode(w); err != nil {
			return fmt.Errorf("Encode DLDataForwarding failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DataForwardingInformation) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.ULDataForwarding = new(UPTNLInformation)
		if err = s.ULDataForwarding.Decode(r); err != nil {
			return fmt.Errorf("Decode ULDataForwarding failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.DLDataForwarding = new(UPTNLInformation)
		if err = s.DLDataForwarding.Decode(r); err != nil {
			return fmt.Errorf("Decode DLDataForwarding failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.IEExtensions = new(DataForwardingInformationExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for DataForwardingInformation")
	}
	return nil
}
