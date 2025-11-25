package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// CellGroupInformationItem is a generated SEQUENCE type.
type CellGroupInformationItem struct {
	CellGroupID     CellGroupID                         `aper:"lb:0,ub:3,mandatory,ext"`
	ULConfiguration *ULConfiguration                    `aper:"optional,ext"`
	DLTXStop        *DLTXStop                           `aper:"optional,ext"`
	RATType         *RATType                            `aper:"optional,ext"`
	IEExtensions    *CellGroupInformationItemExtensions `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *CellGroupInformationItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.ULConfiguration != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.DLTXStop != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.RATType != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 4
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(4), &aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.CellGroupID.Value), &aper.Constraint{Lb: 0, Ub: 3}, true); err != nil {
		return fmt.Errorf("Encode CellGroupID failed: %w", err)
	}
	if s.ULConfiguration != nil {
		if err = s.ULConfiguration.Encode(w); err != nil {
			return fmt.Errorf("Encode ULConfiguration failed: %w", err)
		}
	}
	if s.DLTXStop != nil {
		if err = s.DLTXStop.Encode(w); err != nil {
			return fmt.Errorf("Encode DLTXStop failed: %w", err)
		}
	}
	if s.RATType != nil {
		if err = s.RATType.Encode(w); err != nil {
			return fmt.Errorf("Encode RATType failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *CellGroupInformationItem) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 4, Ub: 4}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if err = s.CellGroupID.Decode(r); err != nil {
		return fmt.Errorf("Decode CellGroupID failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.ULConfiguration = new(ULConfiguration)
		if err = s.ULConfiguration.Decode(r); err != nil {
			return fmt.Errorf("Decode ULConfiguration failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.DLTXStop = new(DLTXStop)
		if err = s.DLTXStop.Decode(r); err != nil {
			return fmt.Errorf("Decode DLTXStop failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.RATType = new(RATType)
		if err = s.RATType.Decode(r); err != nil {
			return fmt.Errorf("Decode RATType failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<4) > 0 {
		s.IEExtensions = new(CellGroupInformationItemExtensions)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for CellGroupInformationItem */
	}
	return nil
}
