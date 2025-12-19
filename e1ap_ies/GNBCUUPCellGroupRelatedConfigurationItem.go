package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// GNBCUUPCellGroupRelatedConfigurationItem is a generated SEQUENCE type.
type GNBCUUPCellGroupRelatedConfigurationItem struct {
	CellGroupID      CellGroupID                 `aper:"lb:0,ub:3,mandatory"`
	UPTNLInformation UPTNLInformation            `aper:"mandatory"`
	ULConfiguration  *ULConfiguration            `aper:"optional"`
	IEExtensions     *ProtocolExtensionContainer `aper:"optional"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *GNBCUUPCellGroupRelatedConfigurationItem) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(false); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.ULConfiguration != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(2), &aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.CellGroupID.Value), &aper.Constraint{Lb: 0, Ub: 3}, true); err != nil {
		return fmt.Errorf("Encode CellGroupID failed: %w", err)
	}
	if err = s.UPTNLInformation.Encode(w); err != nil {
		return fmt.Errorf("Encode UPTNLInformation failed: %w", err)
	}
	if s.ULConfiguration != nil {
		if err = w.WriteEnumerate(uint64((*s.ULConfiguration).Value), aper.Constraint{Lb: 0, Ub: 2}, true); err != nil {
			return fmt.Errorf("Encode ULConfiguration failed: %w", err)
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
func (s *GNBCUUPCellGroupRelatedConfigurationItem) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}
	if err = s.CellGroupID.Decode(r); err != nil {
		return fmt.Errorf("Decode CellGroupID failed: %w", err)
	}
	if err = s.UPTNLInformation.Decode(r); err != nil {
		return fmt.Errorf("Decode UPTNLInformation failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.ULConfiguration = new(ULConfiguration)
		if err = s.ULConfiguration.Decode(r); err != nil {
			return fmt.Errorf("Decode ULConfiguration failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible { /* TODO: Implement extension skipping for GNBCUUPCellGroupRelatedConfigurationItem */
	}
	return nil
}
