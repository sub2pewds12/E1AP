package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// QoSFlowLevelQoSParametersReflectiveQOSIndicator is a generated ENUMERATED type.
type QoSFlowLevelQoSParametersReflectiveQOSIndicator struct {
	Value aper.Enumerated
}

const (
	QoSFlowLevelQoSParametersReflectiveQOSIndicatorEnabled aper.Enumerated = 0
)

func (e *QoSFlowLevelQoSParametersReflectiveQOSIndicator) Encode(w *aper.AperWriter) error {
	// Encode logic for enum QoSFlowLevelQoSParametersReflectiveQOSIndicator to be generated here.
	return fmt.Errorf("Encode not implemented for enum QoSFlowLevelQoSParametersReflectiveQOSIndicator")
}

func (e *QoSFlowLevelQoSParametersReflectiveQOSIndicator) Decode(r *aper.AperReader) error {
	// Decode logic for enum QoSFlowLevelQoSParametersReflectiveQOSIndicator to be generated here.
	return fmt.Errorf("Decode not implemented for enum QoSFlowLevelQoSParametersReflectiveQOSIndicator")
}
