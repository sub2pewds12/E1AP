package e1ap_ies

import (
	"fmt"
	"io"
)

// QOSFlowItemExtensions is a generated type-safe wrapper for extensions.
type QOSFlowItemExtensions struct {
	QoSFlowMappingIndication *QOSFlowMappingIndication
}

// Encode implements the aper.AperMarshaller interface.
func (s *QOSFlowItemExtensions) Encode(w io.Writer) error {
	// TODO: Implement custom extension encoding
	return fmt.Errorf("Encode not implemented for QOSFlowItemExtensions")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *QOSFlowItemExtensions) Decode(r io.Reader) error {
	// TODO: Implement custom extension decoding
	return fmt.Errorf("Decode not implemented for QOSFlowItemExtensions")
}
