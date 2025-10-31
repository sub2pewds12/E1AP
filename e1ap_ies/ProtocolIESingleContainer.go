package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// ProtocolIESingleContainer is a generated SEQUENCE type.
type ProtocolIESingleContainer struct {
	Value ProtocolIEField `aper:"mandatory"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *ProtocolIESingleContainer) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(false); err != nil {
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	if err = s.Value.Encode(w); err != nil {
		return fmt.Errorf("Encode Value failed: %w", err)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ProtocolIESingleContainer) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	if err = s.Value.Decode(r); err != nil {
		return fmt.Errorf("Decode Value failed: %w", err)
	}
	if isExtensible {
		return fmt.Errorf("Extensions not yet implemented for ProtocolIESingleContainer")
	}
	return nil
}
