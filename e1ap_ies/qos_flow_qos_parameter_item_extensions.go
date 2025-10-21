package e1ap_ies

import (
	"fmt"
	"io"
)

// QOSFlowQOSParameterItemExtensions is a generated type-safe wrapper for extensions.
type QOSFlowQOSParameterItemExtensions struct {
	RedundantQosFlowIndicator *RedundantQoSFlowIndicator
	TSCTrafficCharacteristics *TSCTrafficCharacteristics
}

// Encode implements the aper.AperMarshaller interface.
func (s *QOSFlowQOSParameterItemExtensions) Encode(w io.Writer) error {
	// TODO: Implement custom extension encoding
	return fmt.Errorf("Encode not implemented for QOSFlowQOSParameterItemExtensions")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *QOSFlowQOSParameterItemExtensions) Decode(r io.Reader) error {
	// TODO: Implement custom extension decoding
	return fmt.Errorf("Decode not implemented for QOSFlowQOSParameterItemExtensions")
}
