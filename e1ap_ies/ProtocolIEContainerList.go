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
		return fmt.Errorf("Encode extensibility bool failed: %w", err)
	}
	return nil
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *ProtocolIEContainerList) Decode(r *aper.AperReader) (err error) {
	var isExtensible bool
	if isExtensible, err = r.ReadBool(); err != nil {
		return fmt.Errorf("Read extensibility bool failed: %w", err)
	}
	if isExtensible { /* TODO: Implement extension skipping for ProtocolIEContainerList */
	}
	return nil
}
