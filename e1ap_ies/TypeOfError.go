package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// TypeOfError is a generated ENUMERATED type.
type TypeOfError struct {
	Value aper.Enumerated
}

const (
	TypeOfErrorNotUnderstood aper.Enumerated = 0
	TypeOfErrorMissing       aper.Enumerated = 1
)

func (e *TypeOfError) Encode(w *aper.AperWriter) error {
	// Encode logic for enum TypeOfError to be generated here.
	return fmt.Errorf("Encode not implemented for enum TypeOfError")
}

func (e *TypeOfError) Decode(r *aper.AperReader) error {
	// Decode logic for enum TypeOfError to be generated here.
	return fmt.Errorf("Decode not implemented for enum TypeOfError")
}
