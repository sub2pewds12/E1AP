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
		return fmt.Errorf("encode extensibility bool failed: %w", err)
	}
	if err = s.Value.Encode(w); err != nil {
		return fmt.Errorf("encode Value failed: %w", err)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ProtocolIESingleContainer) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	if err = s.Value.Decode(r); err != nil {
		return fmt.Errorf("decode Value failed: %w", err)
	}
	if isExtensible { /* TODO: Implement extension skipping for ProtocolIESingleContainer */
	}
	return nil
}
