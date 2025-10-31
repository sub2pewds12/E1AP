package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// UplinkOnlyROHCContinueROHC is a generated ENUMERATED type.
type UplinkOnlyROHCContinueROHC struct {
	Value aper.Enumerated
}

const (
	UplinkOnlyROHCContinueROHCTrue aper.Enumerated = 0
)

func (e *UplinkOnlyROHCContinueROHC) Encode(w *aper.AperWriter) error {
	// Encode logic for enum UplinkOnlyROHCContinueROHC to be generated here.
	return fmt.Errorf("Encode not implemented for enum UplinkOnlyROHCContinueROHC")
}

func (e *UplinkOnlyROHCContinueROHC) Decode(r *aper.AperReader) error {
	// Decode logic for enum UplinkOnlyROHCContinueROHC to be generated here.
	return fmt.Errorf("Decode not implemented for enum UplinkOnlyROHCContinueROHC")
}
