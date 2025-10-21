package e1ap_ies

import (
	"fmt"
	"io"
)

// PDUSessionResourceModifiedItemExtensions is a generated type-safe wrapper for extensions.
type PDUSessionResourceModifiedItemExtensions struct {
	RedundantNGDLUPTNLInformation *UPTNLInformation
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDUSessionResourceModifiedItemExtensions) Encode(w io.Writer) error {
	// TODO: Implement custom extension encoding
	return fmt.Errorf("Encode not implemented for PDUSessionResourceModifiedItemExtensions")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *PDUSessionResourceModifiedItemExtensions) Decode(r io.Reader) error {
	// TODO: Implement custom extension decoding
	return fmt.Errorf("Decode not implemented for PDUSessionResourceModifiedItemExtensions")
}
