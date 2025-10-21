package e1ap_ies

import (
	"fmt"
	"io"
)

// DRBToSetupItemNGRANExtensions is a generated type-safe wrapper for extensions.
type DRBToSetupItemNGRANExtensions struct {
	DRBQOS                      *QoSFlowLevelQoSParameters
	DAPSRequestInfo             *DAPSRequestInfo
	IgnoreMappingRuleIndication *IgnoreMappingRuleIndication
	QoSFlowsDRBRemapping        *QOSFlowsDRBRemapping
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBToSetupItemNGRANExtensions) Encode(w io.Writer) error {
	// TODO: Implement custom extension encoding
	return fmt.Errorf("Encode not implemented for DRBToSetupItemNGRANExtensions")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DRBToSetupItemNGRANExtensions) Decode(r io.Reader) error {
	// TODO: Implement custom extension decoding
	return fmt.Errorf("Decode not implemented for DRBToSetupItemNGRANExtensions")
}
