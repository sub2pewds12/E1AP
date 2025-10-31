package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DataDiscardRequired is a generated ENUMERATED type.
type DataDiscardRequired struct {
	Value aper.Enumerated
}

const (
	DataDiscardRequiredRequired aper.Enumerated = 0
)

func (e *DataDiscardRequired) Encode(w *aper.AperWriter) error {
	// Encode logic for enum DataDiscardRequired to be generated here.
	return fmt.Errorf("Encode not implemented for enum DataDiscardRequired")
}

func (e *DataDiscardRequired) Decode(r *aper.AperReader) error {
	// Decode logic for enum DataDiscardRequired to be generated here.
	return fmt.Errorf("Decode not implemented for enum DataDiscardRequired")
}
