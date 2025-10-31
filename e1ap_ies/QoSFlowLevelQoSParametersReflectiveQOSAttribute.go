package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// QoSFlowLevelQoSParametersReflectiveQOSAttribute is a generated ENUMERATED type.
type QoSFlowLevelQoSParametersReflectiveQOSAttribute struct {
	Value aper.Enumerated
}

const (
	QoSFlowLevelQoSParametersReflectiveQOSAttributeSubjectTo aper.Enumerated = 0
)

func (e *QoSFlowLevelQoSParametersReflectiveQOSAttribute) Encode(w *aper.AperWriter) error {
	// Encode logic for enum QoSFlowLevelQoSParametersReflectiveQOSAttribute to be generated here.
	return fmt.Errorf("Encode not implemented for enum QoSFlowLevelQoSParametersReflectiveQOSAttribute")
}

func (e *QoSFlowLevelQoSParametersReflectiveQOSAttribute) Decode(r *aper.AperReader) error {
	// Decode logic for enum QoSFlowLevelQoSParametersReflectiveQOSAttribute to be generated here.
	return fmt.Errorf("Decode not implemented for enum QoSFlowLevelQoSParametersReflectiveQOSAttribute")
}
