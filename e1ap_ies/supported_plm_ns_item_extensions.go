package e1ap_ies

import (
	"fmt"
	"io"
)

// SupportedPLMNsItemExtensions is a generated type-safe wrapper for extensions.
type SupportedPLMNsItemExtensions struct {
	NPNSupportInfo           *NPNSupportInfo
	ExtendedSliceSupportList *ExtendedSliceSupportList
	ExtendedNRCGISupportList *ExtendedNRCGISupportList
}

// Encode implements the aper.AperMarshaller interface.
func (s *SupportedPLMNsItemExtensions) Encode(w io.Writer) error {
	// TODO: Implement custom extension encoding
	return fmt.Errorf("Encode not implemented for SupportedPLMNsItemExtensions")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *SupportedPLMNsItemExtensions) Decode(r io.Reader) error {
	// TODO: Implement custom extension decoding
	return fmt.Errorf("Decode not implemented for SupportedPLMNsItemExtensions")
}
