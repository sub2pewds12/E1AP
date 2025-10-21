package e1ap_ies

import (
	"fmt"
	"io"
)

// TraceActivationExtensions is a generated type-safe wrapper for extensions.
type TraceActivationExtensions struct {
	MDTConfiguration         *MDTConfiguration
	TraceCollectionEntityURI *URIaddress
}

// Encode implements the aper.AperMarshaller interface.
func (s *TraceActivationExtensions) Encode(w io.Writer) error {
	// TODO: Implement custom extension encoding
	return fmt.Errorf("Encode not implemented for TraceActivationExtensions")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *TraceActivationExtensions) Decode(r io.Reader) error {
	// TODO: Implement custom extension decoding
	return fmt.Errorf("Decode not implemented for TraceActivationExtensions")
}
