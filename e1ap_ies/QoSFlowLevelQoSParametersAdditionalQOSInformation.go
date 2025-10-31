package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// QoSFlowLevelQoSParametersAdditionalQOSInformation is a generated ENUMERATED type.
type QoSFlowLevelQoSParametersAdditionalQOSInformation struct {
	Value aper.Enumerated
}

const (
	QoSFlowLevelQoSParametersAdditionalQOSInformationMoreLikely aper.Enumerated = 0
)

func (e *QoSFlowLevelQoSParametersAdditionalQOSInformation) Encode(w *aper.AperWriter) error {
	// Encode logic for enum QoSFlowLevelQoSParametersAdditionalQOSInformation to be generated here.
	return fmt.Errorf("Encode not implemented for enum QoSFlowLevelQoSParametersAdditionalQOSInformation")
}

func (e *QoSFlowLevelQoSParametersAdditionalQOSInformation) Decode(r *aper.AperReader) error {
	// Decode logic for enum QoSFlowLevelQoSParametersAdditionalQOSInformation to be generated here.
	return fmt.Errorf("Decode not implemented for enum QoSFlowLevelQoSParametersAdditionalQOSInformation")
}
