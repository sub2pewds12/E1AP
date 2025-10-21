package e1ap_ies

import (
	"fmt"
	"io"
)

// PDUSessionResourceToSetupItemExtensions is a generated type-safe wrapper for extensions.
type PDUSessionResourceToSetupItemExtensions struct {
	CommonNetworkInstance          *CommonNetworkInstance
	RedundantNGULUPTNLInformation  *UPTNLInformation
	RedundantCommonNetworkInstance *CommonNetworkInstance
	RedundantPDUSessionInformation *RedundantPDUSessionInformation
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDUSessionResourceToSetupItemExtensions) Encode(w io.Writer) error {
	// TODO: Implement custom extension encoding
	return fmt.Errorf("Encode not implemented for PDUSessionResourceToSetupItemExtensions")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *PDUSessionResourceToSetupItemExtensions) Decode(r io.Reader) error {
	// TODO: Implement custom extension decoding
	return fmt.Errorf("Decode not implemented for PDUSessionResourceToSetupItemExtensions")
}
