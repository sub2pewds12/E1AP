package e1ap_ies

import (
	"fmt"
	"io"
)

// GNBCUCPTNLAToRemoveItemExtensions is a generated type-safe wrapper for extensions.
type GNBCUCPTNLAToRemoveItemExtensions struct {
	TNLAssociationTransportLayerAddressgNBCUUP *CPTNLInformation
}

// Encode implements the aper.AperMarshaller interface.
func (s *GNBCUCPTNLAToRemoveItemExtensions) Encode(w io.Writer) error {
	// TODO: Implement custom extension encoding
	return fmt.Errorf("Encode not implemented for GNBCUCPTNLAToRemoveItemExtensions")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *GNBCUCPTNLAToRemoveItemExtensions) Decode(r io.Reader) error {
	// TODO: Implement custom extension decoding
	return fmt.Errorf("Decode not implemented for GNBCUCPTNLAToRemoveItemExtensions")
}
