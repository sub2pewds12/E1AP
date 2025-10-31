package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// AdditionalHandoverInfo is a generated ENUMERATED type.
type AdditionalHandoverInfo struct {
	Value aper.Enumerated
}

const (
	AdditionalHandoverInfoDiscardPdpcSN aper.Enumerated = 0
)

func (e *AdditionalHandoverInfo) Encode(w *aper.AperWriter) error {
	// Encode logic for enum AdditionalHandoverInfo to be generated here.
	return fmt.Errorf("Encode not implemented for enum AdditionalHandoverInfo")
}

func (e *AdditionalHandoverInfo) Decode(r *aper.AperReader) error {
	// Decode logic for enum AdditionalHandoverInfo to be generated here.
	return fmt.Errorf("Decode not implemented for enum AdditionalHandoverInfo")
}
