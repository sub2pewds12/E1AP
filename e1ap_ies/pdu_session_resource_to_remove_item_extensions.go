package e1ap_ies

import (
	"fmt"
	"io"
)

// PDUSessionResourceToRemoveItemExtensions is a generated type-safe wrapper for extensions.
type PDUSessionResourceToRemoveItemExtensions struct {
	Cause *Cause
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDUSessionResourceToRemoveItemExtensions) Encode(w io.Writer) error {
	// TODO: Implement custom extension encoding
	return fmt.Errorf("Encode not implemented for PDUSessionResourceToRemoveItemExtensions")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *PDUSessionResourceToRemoveItemExtensions) Decode(r io.Reader) error {
	// TODO: Implement custom extension decoding
	return fmt.Errorf("Decode not implemented for PDUSessionResourceToRemoveItemExtensions")
}
