package e1ap_ies

import (
	"fmt"
	"io"
)

// PDUSessionResourceSetupModItemExtensions is a generated type-safe wrapper for extensions.
type PDUSessionResourceSetupModItemExtensions struct {
	RedundantNGDLUPTNLInformation *UPTNLInformation
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDUSessionResourceSetupModItemExtensions) Encode(w io.Writer) error {
	// TODO: Implement custom extension encoding
	return fmt.Errorf("Encode not implemented for PDUSessionResourceSetupModItemExtensions")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *PDUSessionResourceSetupModItemExtensions) Decode(r io.Reader) error {
	// TODO: Implement custom extension decoding
	return fmt.Errorf("Decode not implemented for PDUSessionResourceSetupModItemExtensions")
}
