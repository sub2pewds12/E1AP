package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// UEAssociatedLogicalE1ConnectionItemResAck is a generated SEQUENCE type.
type UEAssociatedLogicalE1ConnectionItemResAck struct {
	UEAssociatedLogicalE1ConnectionItem UEAssociatedLogicalE1ConnectionItem `aper:"mandatory"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *UEAssociatedLogicalE1ConnectionItemResAck) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(false); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	if err = s.UEAssociatedLogicalE1ConnectionItem.Encode(w); err != nil {
		return fmt.Errorf("Encode UEAssociatedLogicalE1ConnectionItem failed: %w", err)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *UEAssociatedLogicalE1ConnectionItemResAck) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	if err = s.UEAssociatedLogicalE1ConnectionItem.Decode(r); err != nil {
		return fmt.Errorf("Decode UEAssociatedLogicalE1ConnectionItem failed: %w", err)
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for UEAssociatedLogicalE1ConnectionItemResAck")
	}
	return nil
}
