package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DRBSetupItemEUTRANS1DLUPUnchanged is a generated ENUMERATED type.
type DRBSetupItemEUTRANS1DLUPUnchanged struct {
	Value aper.Enumerated
}

const (
	DRBSetupItemEUTRANS1DLUPUnchangedTrue aper.Enumerated = 0
)

func (e *DRBSetupItemEUTRANS1DLUPUnchanged) Encode(w *aper.AperWriter) error {
	// Encode logic for enum DRBSetupItemEUTRANS1DLUPUnchanged to be generated here.
	return fmt.Errorf("Encode not implemented for enum DRBSetupItemEUTRANS1DLUPUnchanged")
}

func (e *DRBSetupItemEUTRANS1DLUPUnchanged) Decode(r *aper.AperReader) error {
	// Decode logic for enum DRBSetupItemEUTRANS1DLUPUnchanged to be generated here.
	return fmt.Errorf("Decode not implemented for enum DRBSetupItemEUTRANS1DLUPUnchanged")
}
