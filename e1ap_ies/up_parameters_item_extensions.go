package e1ap_ies

import (
	"fmt"
	"io"
)

// UPParametersItemExtensions is a generated type-safe wrapper for extensions.
type UPParametersItemExtensions struct {
	QOSMappingInformation *QOSMappingInformation
}

// Encode implements the aper.AperMarshaller interface.
func (s *UPParametersItemExtensions) Encode(w io.Writer) error {
	// TODO: Implement custom extension encoding
	return fmt.Errorf("Encode not implemented for UPParametersItemExtensions")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *UPParametersItemExtensions) Decode(r io.Reader) error {
	// TODO: Implement custom extension decoding
	return fmt.Errorf("Decode not implemented for UPParametersItemExtensions")
}
