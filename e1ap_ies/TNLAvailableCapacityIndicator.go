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
		return fmt.Errorf("encode extensibility bool failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.DLTNLOfferedCapacity.Value), &aper.Constraint{Lb: 0, Ub: 16777216}, true); err != nil {
		return fmt.Errorf("encode DLTNLOfferedCapacity failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.DLTNLAvailableCapacity.Value), &aper.Constraint{Lb: 0, Ub: 100}, true); err != nil {
		return fmt.Errorf("encode DLTNLAvailableCapacity failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.ULTNLOfferedCapacity.Value), &aper.Constraint{Lb: 0, Ub: 16777216}, true); err != nil {
		return fmt.Errorf("encode ULTNLOfferedCapacity failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.ULTNLAvailableCapacity.Value), &aper.Constraint{Lb: 0, Ub: 100}, true); err != nil {
		return fmt.Errorf("encode ULTNLAvailableCapacity failed: %w", err)
	}
	if err = s.IEExtensions.Encode(w); err != nil {
		return fmt.Errorf("encode IEExtensions failed: %w", err)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *TNLAvailableCapacityIndicator) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	if err = s.DLTNLOfferedCapacity.Decode(r); err != nil {
		return fmt.Errorf("decode DLTNLOfferedCapacity failed: %w", err)
	}
	if err = s.DLTNLAvailableCapacity.Decode(r); err != nil {
		return fmt.Errorf("decode DLTNLAvailableCapacity failed: %w", err)
	}
	if err = s.ULTNLOfferedCapacity.Decode(r); err != nil {
		return fmt.Errorf("decode ULTNLOfferedCapacity failed: %w", err)
	}
	if err = s.ULTNLAvailableCapacity.Decode(r); err != nil {
		return fmt.Errorf("decode ULTNLAvailableCapacity failed: %w", err)
	}
	if err = s.IEExtensions.Decode(r); err != nil {
		return fmt.Errorf("decode IEExtensions failed: %w", err)
	}
	if isExtensible { /* TODO: Implement extension skipping for TNLAvailableCapacityIndicator */
	}
	return nil
}
