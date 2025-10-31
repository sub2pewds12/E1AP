package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// SDAPHeaderUL is a generated ENUMERATED type.
type SDAPHeaderUL struct {
	Value aper.Enumerated
}

const (
	SDAPHeaderULPresent aper.Enumerated = 0
	SDAPHeaderULAbsent  aper.Enumerated = 1
)

func (e *SDAPHeaderUL) Encode(w *aper.AperWriter) error {
	// Encode logic for enum SDAPHeaderUL to be generated here.
	return fmt.Errorf("Encode not implemented for enum SDAPHeaderUL")
}

func (e *SDAPHeaderUL) Decode(r *aper.AperReader) error {
	// Decode logic for enum SDAPHeaderUL to be generated here.
	return fmt.Errorf("Decode not implemented for enum SDAPHeaderUL")
}
