package e1ap_ies

import (
	"fmt"
	"io"
)

// CellGroupInformationItemExtensions is a generated type-safe wrapper for extensions.
type CellGroupInformationItemExtensions struct {
	NumberOfTunnels *NumberOfTunnels
}

// Encode implements the aper.AperMarshaller interface.
func (s *CellGroupInformationItemExtensions) Encode(w io.Writer) error {
	// TODO: Implement custom extension encoding
	return fmt.Errorf("Encode not implemented for CellGroupInformationItemExtensions")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *CellGroupInformationItemExtensions) Decode(r io.Reader) error {
	// TODO: Implement custom extension decoding
	return fmt.Errorf("Decode not implemented for CellGroupInformationItemExtensions")
}
