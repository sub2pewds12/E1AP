package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DRBToSetupItemNGRANExtensions is a generated type-safe wrapper for extensions.
type DRBToSetupItemNGRANExtensions struct {
	DRBQOS                      *QoSFlowLevelQoSParameters
	DAPSRequestInfo             *DAPSRequestInfo
	IgnoreMappingRuleIndication *IgnoreMappingRuleIndication
	QoSFlowsDRBRemapping        *QOSFlowsDRBRemapping
}

// Encode implements the aper.AperMarshaller interface.
func (s *DRBToSetupItemNGRANExtensions) Encode(w *aper.AperWriter) error {
	// TODO: Implement the complex APER extension container encoding logic.
	// This involves creating a bitmap of present extensions and encoding each one as an open type.
	return fmt.Errorf("Encode not yet implemented for DRBToSetupItemNGRANExtensions")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *DRBToSetupItemNGRANExtensions) Decode(r *aper.AperReader) error {
	// TODO: Implement the complex APER extension container decoding logic.
	// This involves reading a bitmap and then decoding each present extension as an open type.
	return fmt.Errorf("Decode not yet implemented for DRBToSetupItemNGRANExtensions")
}
