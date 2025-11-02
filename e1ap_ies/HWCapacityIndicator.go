package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// HWCapacityIndicator is a generated SEQUENCE type.
type HWCapacityIndicator struct {
	OfferedThroughput   HWCapacityIndicatorOfferedThroughput   `aper:"lb:1,ub:16777216,mandatory,ext"`
	AvailableThroughput HWCapacityIndicatorAvailableThroughput `aper:"lb:0,ub:100,mandatory,ext"`
	IEExtensions        ProtocolExtensionContainer             `aper:"mandatory,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *HWCapacityIndicator) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.OfferedThroughput.Value), &aper.Constraint{Lb: 1, Ub: 16777216}, true); err != nil {
		return fmt.Errorf("Encode OfferedThroughput failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.AvailableThroughput.Value), &aper.Constraint{Lb: 0, Ub: 100}, true); err != nil {
		return fmt.Errorf("Encode AvailableThroughput failed: %w", err)
	}
	if err = s.IEExtensions.Encode(w); err != nil {
		return fmt.Errorf("Encode IEExtensions failed: %w", err)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *HWCapacityIndicator) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 1, Ub: 16777216}, true); err != nil {
			return fmt.Errorf("Decode OfferedThroughput failed: %w", err)
		}
		s.OfferedThroughput.Value = aper.Integer(val)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 100}, true); err != nil {
			return fmt.Errorf("Decode AvailableThroughput failed: %w", err)
		}
		s.AvailableThroughput.Value = aper.Integer(val)
	}
	if err = s.IEExtensions.Decode(r); err != nil {
		return fmt.Errorf("Decode IEExtensions failed: %w", err)
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for HWCapacityIndicator")
	}
	return nil
}
