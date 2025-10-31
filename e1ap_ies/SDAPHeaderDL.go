package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// SDAPHeaderDL is a generated ENUMERATED type.
type SDAPHeaderDL struct {
	Value aper.Enumerated
}

const (
	SDAPHeaderDLPresent aper.Enumerated = 0
	SDAPHeaderDLAbsent  aper.Enumerated = 1
)

func (e *SDAPHeaderDL) Encode(w *aper.AperWriter) error {
	// Encode logic for enum SDAPHeaderDL to be generated here.
	return fmt.Errorf("Encode not implemented for enum SDAPHeaderDL")
}

func (e *SDAPHeaderDL) Decode(r *aper.AperReader) error {
	// Decode logic for enum SDAPHeaderDL to be generated here.
	return fmt.Errorf("Decode not implemented for enum SDAPHeaderDL")
}
