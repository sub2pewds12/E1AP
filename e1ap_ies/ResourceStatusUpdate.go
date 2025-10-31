package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// ResourceStatusUpdate is a generated SEQUENCE type.
type ResourceStatusUpdate struct {
	TransactionID                 TransactionID                                  `aper:"lb:0,ub:255,mandatory,ext"`
	GNBCUCPMeasurementID          ResourceStatusUpdateIEsIDGNBCUCPMeasurementID  `aper:"lb:1,ub:4095,mandatory,ext"`
	GNBCUUPMeasurementID          *ResourceStatusUpdateIEsIDGNBCUUPMeasurementID `aper:"lb:1,ub:4095,optional,ext"`
	TNLAvailableCapacityIndicator *TNLAvailableCapacityIndicator                 `aper:"optional,ext"`
	HWCapacityIndicator           HWCapacityIndicator                            `aper:"mandatory,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *ResourceStatusUpdate) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	var optionalityBitmap [1]byte
	if s.GNBCUUPMeasurementID != nil {
		optionalityBitmap[0] |= 1 << 7
	}
	if s.TNLAvailableCapacityIndicator != nil {
		optionalityBitmap[0] |= 1 << 6
	}
	if err = w.WriteBitString(optionalityBitmap[:], uint(2), &aper.Constraint{Lb: 2, Ub: 2}, false); err != nil {
		return fmt.Errorf("Encode optionality bitmap failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.TransactionID), &aper.Constraint{Lb: 0, Ub: 255}, true); err != nil {
		return fmt.Errorf("Encode TransactionID failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.GNBCUCPMeasurementID), &aper.Constraint{Lb: 1, Ub: 4095}, true); err != nil {
		return fmt.Errorf("Encode GNBCUCPMeasurementID failed: %w", err)
	}
	if s.GNBCUUPMeasurementID != nil {
		if err = w.WriteInteger(int64((*s.GNBCUUPMeasurementID)), &aper.Constraint{Lb: 1, Ub: 4095}, true); err != nil {
			return fmt.Errorf("Encode GNBCUUPMeasurementID failed: %w", err)
		}
	}
	if s.TNLAvailableCapacityIndicator != nil {
		if err = s.TNLAvailableCapacityIndicator.Encode(w); err != nil {
			return fmt.Errorf("Encode TNLAvailableCapacityIndicator failed: %w", err)
		}
	}
	if err = s.HWCapacityIndicator.Encode(w); err != nil {
		return fmt.Errorf("Encode HWCapacityIndicator failed: %w", err)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ResourceStatusUpdate) Decode(r *aper.AperReader) (err error) {
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
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 255}, true); err != nil {
			return fmt.Errorf("Decode TransactionID failed: %w", err)
		}
		s.TransactionID = TransactionID(val)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 4095}, true); err != nil {
			return fmt.Errorf("Decode GNBCUCPMeasurementID failed: %w", err)
		}
		s.GNBCUCPMeasurementID = ResourceStatusUpdateIEsIDGNBCUCPMeasurementID(val)
	}

	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<7) > 0 {

		{
			var val int64
			if val, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 4095}, true); err != nil {
				return fmt.Errorf("Decode GNBCUUPMeasurementID failed: %w", err)
			}
			tmp := ResourceStatusUpdateIEsIDGNBCUUPMeasurementID(val)
			s.GNBCUUPMeasurementID = &tmp
		}

	}
	if len(optionalityBitmap) > 0 && optionalityBitmap[0]&(1<<6) > 0 {
		s.TNLAvailableCapacityIndicator = new(TNLAvailableCapacityIndicator)
		if err = s.TNLAvailableCapacityIndicator.Decode(r); err != nil {
			return fmt.Errorf("Decode TNLAvailableCapacityIndicator failed: %w", err)
		}
	}
	if err = s.HWCapacityIndicator.Decode(r); err != nil {
		return fmt.Errorf("Decode HWCapacityIndicator failed: %w", err)
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for ResourceStatusUpdate")
	}
	return nil
}
