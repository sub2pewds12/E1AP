package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DRBRequiredToModifyItemNGRAN is a generated SEQUENCE type.
type DRBRequiredToModifyItemNGRAN struct {
	DRBID                                DRBID                                      `aper:"lb:1,ub:32,mandatory,ext"`
	GNBCUUPCellGroupRelatedConfiguration []GNBCUUPCellGroupRelatedConfigurationItem `aper:"optional,ext"`
	FlowToRemove                         []QOSFlowItem                              `aper:"optional,ext"`
	Cause                                Cause                                      `aper:"mandatory,ext"`
	IEExtensions                         *ProtocolExtensionContainer                `aper:"optional,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBRequiredToModifyItemNGRAN) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.GNBCUUPCellGroupRelatedConfiguration != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.FlowToRemove != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if s.IEExtensions != nil {
		optionalityBitmap[0] |= 1 << 5
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(3), &aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.DRBID), &aper.Constraint{Lb: 1, Ub: 32}, true); err != nil {
		return fmt.Errorf("Encode DRBID failed: %w", err)
	}
	if s.GNBCUUPCellGroupRelatedConfiguration != nil {
		if err = s.GNBCUUPCellGroupRelatedConfiguration.Encode(w); err != nil {
			return fmt.Errorf("Encode GNBCUUPCellGroupRelatedConfiguration failed: %w", err)
		}
	}
	if s.FlowToRemove != nil {
		if err = s.FlowToRemove.Encode(w); err != nil {
			return fmt.Errorf("Encode FlowToRemove failed: %w", err)
		}
	}
	if err = s.Cause.Encode(w); err != nil {
		return fmt.Errorf("Encode Cause failed: %w", err)
	}
	if s.IEExtensions != nil {
		if err = s.IEExtensions.Encode(w); err != nil {
			return fmt.Errorf("Encode IEExtensions failed: %w", err)
		}
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DRBRequiredToModifyItemNGRAN) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	var optionalityBitmap []byte
	if optionalityBitmap, _, err = r.ReadBitString(&aper.Constraint{Lb: 3, Ub: 3}, false); err != nil {
		return fmt.Errorf("Read optionality bitmap failed: %w", err)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 32}, true); err != nil {
			return fmt.Errorf("Decode DRBID failed: %w", err)
		}
		s.DRBID = DRBID(val)
	}

	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {
		s.GNBCUUPCellGroupRelatedConfiguration = new(GNBCUUPCellGroupRelatedConfiguration)
		if err = s.GNBCUUPCellGroupRelatedConfiguration.Decode(r); err != nil {
			return fmt.Errorf("Decode GNBCUUPCellGroupRelatedConfiguration failed: %w", err)
		}
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.FlowToRemove = new(QOSFlowList)
		if err = s.FlowToRemove.Decode(r); err != nil {
			return fmt.Errorf("Decode FlowToRemove failed: %w", err)
		}
	}
	if err = s.Cause.Decode(r); err != nil {
		return fmt.Errorf("Decode Cause failed: %w", err)
	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<5) > 0 {
		s.IEExtensions = new(ProtocolExtensionContainer)
		if err = s.IEExtensions.Decode(r); err != nil {
			return fmt.Errorf("Decode IEExtensions failed: %w", err)
		}
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for DRBRequiredToModifyItemNGRAN")
	}
	return nil
}
