package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// UEAssociatedLogicalE1ConnectionItemRes is a generated SEQUENCE type.
type UEAssociatedLogicalE1ConnectionItemRes struct {
	UEAssociatedLogicalE1ConnectionItem UEAssociatedLogicalE1ConnectionItem `aper:"mandatory"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *UEAssociatedLogicalE1ConnectionItemRes) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(false); err != nil {
		return fmt.Errorf("encode extensibility bool failed: %w", err)
	}
	if err = s.UEAssociatedLogicalE1ConnectionItem.Encode(w); err != nil {
		return fmt.Errorf("encode UEAssociatedLogicalE1ConnectionItem failed: %w", err)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *UEAssociatedLogicalE1ConnectionItemRes) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	if err = s.UEAssociatedLogicalE1ConnectionItem.Decode(r); err != nil {
		return fmt.Errorf("decode UEAssociatedLogicalE1ConnectionItem failed: %w", err)
	}
	if isExtensible { /* TODO: Implement extension skipping for UEAssociatedLogicalE1ConnectionItemRes */
	}
	return nil
}
