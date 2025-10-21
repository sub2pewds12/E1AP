package e1ap_ies

import (
	"fmt"
	"io"
)

// PDUSessionResourceSetupItemExtensions is a generated type-safe wrapper for extensions.
type PDUSessionResourceSetupItemExtensions struct {
	RedundantNGDLUPTNLInformation      *UPTNLInformation
	RedundantPDUSessionInformationUsed *RedundantPDUSessionInformation
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDUSessionResourceSetupItemExtensions) Encode(w io.Writer) error {
	// TODO: Implement custom extension encoding
	return fmt.Errorf("Encode not implemented for PDUSessionResourceSetupItemExtensions")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *PDUSessionResourceSetupItemExtensions) Decode(r io.Reader) error {
	// TODO: Implement custom extension decoding
	return fmt.Errorf("Decode not implemented for PDUSessionResourceSetupItemExtensions")
}
