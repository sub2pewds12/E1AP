package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// EHCUplinkParametersDRBContinueEHCUL is a generated ENUMERATED type.
type EHCUplinkParametersDRBContinueEHCUL struct {
	Value aper.Enumerated
}

const (
	EHCUplinkParametersDRBContinueEHCULTrue aper.Enumerated = 0
)

func (e *EHCUplinkParametersDRBContinueEHCUL) Encode(w *aper.AperWriter) error {
	// Encode logic for enum EHCUplinkParametersDRBContinueEHCUL to be generated here.
	return fmt.Errorf("Encode not implemented for enum EHCUplinkParametersDRBContinueEHCUL")
}

func (e *EHCUplinkParametersDRBContinueEHCUL) Decode(r *aper.AperReader) error {
	// Decode logic for enum EHCUplinkParametersDRBContinueEHCUL to be generated here.
	return fmt.Errorf("Decode not implemented for enum EHCUplinkParametersDRBContinueEHCUL")
}
