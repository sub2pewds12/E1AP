package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// QoSFlowLevelQoSParametersExtensions is a generated type-safe wrapper for extensions.
type QoSFlowLevelQoSParametersExtensions struct {
	QoSMonitoringRequest            *QosMonitoringRequest
	MCGOfferedGBRQoSFlowInfo        *GBRQoSFlowInformation
	QosMonitoringReportingFrequency *QosMonitoringReportingFrequency
	QoSMonitoringDisabled           *QosMonitoringDisabled
}

// Encode implements the aper.AperMarshaller interface.
func (s *QoSFlowLevelQoSParametersExtensions) Encode(w *aper.AperWriter) error {
	// TODO: Implement the complex APER extension container encoding logic.
	// This involves creating a bitmap of present extensions and encoding each one as an open type.
	return fmt.Errorf("Encode not yet implemented for QoSFlowLevelQoSParametersExtensions")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *QoSFlowLevelQoSParametersExtensions) Decode(r *aper.AperReader) error {
	// TODO: Implement the complex APER extension container decoding logic.
	// This involves reading a bitmap and then decoding each present extension as an open type.
	return fmt.Errorf("Decode not yet implemented for QoSFlowLevelQoSParametersExtensions")
}
