package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// TNLAvailableCapacityIndicator is a generated SEQUENCE type.
type TNLAvailableCapacityIndicator struct {
	DLTNLOfferedCapacity   TNLAvailableCapacityIndicatorDLTNLOfferedCapacity   `aper:"lb:0,ub:16777216,mandatory,ext"`
	DLTNLAvailableCapacity TNLAvailableCapacityIndicatorDLTNLAvailableCapacity `aper:"lb:0,ub:100,mandatory,ext"`
	ULTNLOfferedCapacity   TNLAvailableCapacityIndicatorULTNLOfferedCapacity   `aper:"lb:0,ub:16777216,mandatory,ext"`
	ULTNLAvailableCapacity TNLAvailableCapacityIndicatorULTNLAvailableCapacity `aper:"lb:0,ub:100,mandatory,ext"`
	IEExtensions           ProtocolExtensionContainer                          `aper:"mandatory,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *TNLAvailableCapacityIndicator) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(true); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.DLTNLOfferedCapacity.Value), &aper.Constraint{Lb: 0, Ub: 16777216}, true); err != nil {
		return fmt.Errorf("Encode DLTNLOfferedCapacity failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.DLTNLAvailableCapacity.Value), &aper.Constraint{Lb: 0, Ub: 100}, true); err != nil {
		return fmt.Errorf("Encode DLTNLAvailableCapacity failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.ULTNLOfferedCapacity.Value), &aper.Constraint{Lb: 0, Ub: 16777216}, true); err != nil {
		return fmt.Errorf("Encode ULTNLOfferedCapacity failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.ULTNLAvailableCapacity.Value), &aper.Constraint{Lb: 0, Ub: 100}, true); err != nil {
		return fmt.Errorf("Encode ULTNLAvailableCapacity failed: %w", err)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *TNLAvailableCapacityIndicator) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 16777216}, true); err != nil {
			return fmt.Errorf("Decode DLTNLOfferedCapacity failed: %w", err)
		}
		s.DLTNLOfferedCapacity.Value = aper.Integer(val)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 100}, true); err != nil {
			return fmt.Errorf("Decode DLTNLAvailableCapacity failed: %w", err)
		}
		s.DLTNLAvailableCapacity.Value = aper.Integer(val)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 16777216}, true); err != nil {
			return fmt.Errorf("Decode ULTNLOfferedCapacity failed: %w", err)
		}
		s.ULTNLOfferedCapacity.Value = aper.Integer(val)
	}

	{
		var val int64
		if val, err = r.ReadInteger(&aper.Constraint{Lb: 0, Ub: 100}, true); err != nil {
			return fmt.Errorf("Decode ULTNLAvailableCapacity failed: %w", err)
		}
		s.ULTNLAvailableCapacity.Value = aper.Integer(val)
	}
	if err = s.IEExtensions.Decode(r); err != nil {
		return fmt.Errorf("Decode IEExtensions failed: %w", err)
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for TNLAvailableCapacityIndicator")
	}
	return nil
}
