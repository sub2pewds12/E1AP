package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// ProtocolIEContainerList is a generated SEQUENCE type.
type ProtocolIEContainerList struct {
}

// Encode implements the aper.AperMarshaller interface.
func (s *ProtocolIEContainerList) Encode(w *aper.AperWriter) (err error) {
	if err = w.WriteBool(false); err != nil {
		return fmt.Errorf("encode extensibility bool failed: %w", err)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ProtocolIEContainerList) Decode(r *aper.AperReader) (err error) {
	isExtensible, err := r.ReadBool()
	if err != nil {
		return fmt.Errorf("read extensibility bool failed: %w", err)
	}
	_ = isExtensible
	if isExtensible { /* TODO: Implement extension skipping for ProtocolIEContainerList */
	}
	return nil
}
