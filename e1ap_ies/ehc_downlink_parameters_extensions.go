package e1ap_ies

import (
	"fmt"
	"io"
)

// EHCDownlinkParametersExtensions is a generated type-safe wrapper for extensions.
type EHCDownlinkParametersExtensions struct {
	MaxCIDEHCDL *MaxCIDEHCDL
}

// Encode implements the aper.AperMarshaller interface.
func (s *EHCDownlinkParametersExtensions) Encode(w io.Writer) error {
	// TODO: Implement custom extension encoding
	return fmt.Errorf("Encode not implemented for EHCDownlinkParametersExtensions")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *EHCDownlinkParametersExtensions) Decode(r io.Reader) error {
	// TODO: Implement custom extension decoding
	return fmt.Errorf("Decode not implemented for EHCDownlinkParametersExtensions")
}
