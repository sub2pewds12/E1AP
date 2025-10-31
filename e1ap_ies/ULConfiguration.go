package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// ULConfiguration is a generated ENUMERATED type.
type ULConfiguration struct {
	Value aper.Enumerated
}

const (
	ULConfigurationNoData aper.Enumerated = 0
	ULConfigurationShared aper.Enumerated = 1
	ULConfigurationOnly   aper.Enumerated = 2
)

func (e *ULConfiguration) Encode(w *aper.AperWriter) error {
	// Encode logic for enum ULConfiguration to be generated here.
	return fmt.Errorf("Encode not implemented for enum ULConfiguration")
}

func (e *ULConfiguration) Decode(r *aper.AperReader) error {
	// Decode logic for enum ULConfiguration to be generated here.
	return fmt.Errorf("Decode not implemented for enum ULConfiguration")
}
