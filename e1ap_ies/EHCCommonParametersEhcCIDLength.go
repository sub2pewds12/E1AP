package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// EHCCommonParametersEhcCIDLength is a generated ENUMERATED type.
type EHCCommonParametersEhcCIDLength struct {
	Value aper.Enumerated
}

const (
	EHCCommonParametersEhcCIDLengthBits7  aper.Enumerated = 0
	EHCCommonParametersEhcCIDLengthBits15 aper.Enumerated = 1
)

func (e *EHCCommonParametersEhcCIDLength) Encode(w *aper.AperWriter) error {
	// Encode logic for enum EHCCommonParametersEhcCIDLength to be generated here.
	return fmt.Errorf("Encode not implemented for enum EHCCommonParametersEhcCIDLength")
}

func (e *EHCCommonParametersEhcCIDLength) Decode(r *aper.AperReader) error {
	// Decode logic for enum EHCCommonParametersEhcCIDLength to be generated here.
	return fmt.Errorf("Decode not implemented for enum EHCCommonParametersEhcCIDLength")
}
