package e1ap_ies

import (
	"fmt"
	"io"
)

// E1ReleaseRequest is a generated SEQUENCE type.
type E1ReleaseRequest struct {
	TransactionID TransactionID `aper:"lb:0,ub:255,mandatory,reject,ext"`
	Cause         Cause         `aper:"mandatory,ignore,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *E1ReleaseRequest) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for E1ReleaseRequest")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *E1ReleaseRequest) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for E1ReleaseRequest")
}
