package e1ap_ies

import (
	"fmt"
	"io"
)

// DRBToSetupModItemNGRANExtensions is a generated type-safe wrapper for extensions.
type DRBToSetupModItemNGRANExtensions struct {
	DRBQOS                      *QoSFlowLevelQoSParameters
	IgnoreMappingRuleIndication *IgnoreMappingRuleIndication
	DAPSRequestInfo             *DAPSRequestInfo
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBToSetupModItemNGRANExtensions) Encode(w io.Writer) error {
	// TODO: Implement custom extension encoding
	return fmt.Errorf("Encode not implemented for DRBToSetupModItemNGRANExtensions")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DRBToSetupModItemNGRANExtensions) Decode(r io.Reader) error {
	// TODO: Implement custom extension decoding
	return fmt.Errorf("Decode not implemented for DRBToSetupModItemNGRANExtensions")
}
