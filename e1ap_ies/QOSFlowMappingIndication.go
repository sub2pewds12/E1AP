package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// QOSFlowMappingIndication is a generated ENUMERATED type.
type QOSFlowMappingIndication struct {
	Value aper.Enumerated
}

const (
	QOSFlowMappingIndicationUl aper.Enumerated = 0
	QOSFlowMappingIndicationDl aper.Enumerated = 1
)

func (e *QOSFlowMappingIndication) Encode(w *aper.AperWriter) error {
	// Encode logic for enum QOSFlowMappingIndication to be generated here.
	return fmt.Errorf("Encode not implemented for enum QOSFlowMappingIndication")
}

func (e *QOSFlowMappingIndication) Decode(r *aper.AperReader) error {
	// Decode logic for enum QOSFlowMappingIndication to be generated here.
	return fmt.Errorf("Decode not implemented for enum QOSFlowMappingIndication")
}
