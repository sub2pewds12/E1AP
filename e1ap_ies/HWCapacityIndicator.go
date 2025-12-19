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
		return fmt.Errorf("encode extensibility bool failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.OfferedThroughput.Value), &aper.Constraint{Lb: 1, Ub: 16777216}, true); err != nil {
		return fmt.Errorf("encode OfferedThroughput failed: %w", err)
	}
	if err = w.WriteInteger(int64(s.AvailableThroughput.Value), &aper.Constraint{Lb: 0, Ub: 100}, true); err != nil {
		return fmt.Errorf("encode AvailableThroughput failed: %w", err)
	}
	if err = s.IEExtensions.Encode(w); err != nil {
		return fmt.Errorf("encode IEExtensions failed: %w", err)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *HWCapacityIndicator) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	if err = s.OfferedThroughput.Decode(r); err != nil {
		return fmt.Errorf("decode OfferedThroughput failed: %w", err)
	}
	if err = s.AvailableThroughput.Decode(r); err != nil {
		return fmt.Errorf("decode AvailableThroughput failed: %w", err)
	}
	if err = s.IEExtensions.Decode(r); err != nil {
		return fmt.Errorf("decode IEExtensions failed: %w", err)
	}
	if isExtensible { /* TODO: Implement extension skipping for HWCapacityIndicator */
	}
	return nil
}
