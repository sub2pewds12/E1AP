package e1ap_ies

import (
	"fmt"
	"io"
)

// GNBCUUPStatusIndication is a generated SEQUENCE type.
type GNBCUUPStatusIndication struct {
	TransactionID              TransactionID              `aper:"lb:0,ub:255,mandatory,reject,ext"`
	GNBCUUPOverloadInformation GNBCUUPOverloadInformation `aper:"mandatory,reject,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *GNBCUUPStatusIndication) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for GNBCUUPStatusIndication")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GNBCUUPStatusIndication) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for GNBCUUPStatusIndication")
}
