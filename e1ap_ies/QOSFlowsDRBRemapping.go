package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// QOSFlowsDRBRemapping is a generated ENUMERATED type.
type QOSFlowsDRBRemapping struct {
	Value aper.Enumerated
}

const (
	QOSFlowsDRBRemappingUpdate              aper.Enumerated = 0
	QOSFlowsDRBRemappingSourceConfiguration aper.Enumerated = 1
)

func (e *QOSFlowsDRBRemapping) Encode(w *aper.AperWriter) error {
	// Encode logic for enum QOSFlowsDRBRemapping to be generated here.
	return fmt.Errorf("Encode not implemented for enum QOSFlowsDRBRemapping")
}

func (e *QOSFlowsDRBRemapping) Decode(r *aper.AperReader) error {
	// Decode logic for enum QOSFlowsDRBRemapping to be generated here.
	return fmt.Errorf("Decode not implemented for enum QOSFlowsDRBRemapping")
}
