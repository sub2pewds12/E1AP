package e1ap_ies

import (
	"fmt"
	"io"
)

// PDUSessionResourceToSetupModItemExtensions is a generated type-safe wrapper for extensions.
type PDUSessionResourceToSetupModItemExtensions struct {
	NetworkInstance                *NetworkInstance
	CommonNetworkInstance          *CommonNetworkInstance
	RedundantNGULUPTNLInformation  *UPTNLInformation
	RedundantCommonNetworkInstance *CommonNetworkInstance
}

// Encode implements the aper.AperMarshaller interface.
func (s *PDUSessionResourceToSetupModItemExtensions) Encode(w io.Writer) error {
	// TODO: Implement custom extension encoding
	return fmt.Errorf("Encode not implemented for PDUSessionResourceToSetupModItemExtensions")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *PDUSessionResourceToSetupModItemExtensions) Decode(r io.Reader) error {
	// TODO: Implement custom extension decoding
	return fmt.Errorf("Decode not implemented for PDUSessionResourceToSetupModItemExtensions")
}
