package e1ap_ies

import (
	"fmt"
	"io"
)

// QoSFlowLevelQoSParametersExtensions is a generated type-safe wrapper for extensions.
type QoSFlowLevelQoSParametersExtensions struct {
	QoSMonitoringRequest            *QosMonitoringRequest
	MCGOfferedGBRQoSFlowInfo        *GBRQoSFlowInformation
	QosMonitoringReportingFrequency *QosMonitoringReportingFrequency
	QoSMonitoringDisabled           *QosMonitoringDisabled
}

// Encode implements the aper.AperMarshaller interface.
func (s *QoSFlowLevelQoSParametersExtensions) Encode(w io.Writer) error {
	// TODO: Implement custom extension encoding
	return fmt.Errorf("Encode not implemented for QoSFlowLevelQoSParametersExtensions")
}

// Decode implements the aper.AperUnmarshaller interface.
func (s *QoSFlowLevelQoSParametersExtensions) Decode(r io.Reader) error {
	// TODO: Implement custom extension decoding
	return fmt.Errorf("Decode not implemented for QoSFlowLevelQoSParametersExtensions")
}
