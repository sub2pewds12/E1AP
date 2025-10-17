package e1ap_ies

import (
	"fmt"
	"io"
)

// E1ReleaseResponse is a generated SEQUENCE type.
type E1ReleaseResponse struct {
	TransactionID TransactionID `aper:"lb:0,ub:255,mandatory,reject,ext"`
}

// Encode implements the aper.AperMarshaller interface.
func (s *E1ReleaseResponse) Encode(w io.Writer) error {
	_ = w // Placeholder to prevent unused import warning
	return fmt.Errorf("Encode not implemented for E1ReleaseResponse")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *E1ReleaseResponse) Decode(r io.Reader) error {
	_ = r // Placeholder to prevent unused import warning
	return fmt.Errorf("Decode not implemented for E1ReleaseResponse")
}
