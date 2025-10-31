package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// GNBCUUPOverloadInformation is a generated ENUMERATED type.
type GNBCUUPOverloadInformation struct {
	Value aper.Enumerated
}

const (
	GNBCUUPOverloadInformationOverloaded    aper.Enumerated = 0
	GNBCUUPOverloadInformationNotOverloaded aper.Enumerated = 1
)

func (e *GNBCUUPOverloadInformation) Encode(w *aper.AperWriter) error {
	// Encode logic for enum GNBCUUPOverloadInformation to be generated here.
	return fmt.Errorf("Encode not implemented for enum GNBCUUPOverloadInformation")
}

func (e *GNBCUUPOverloadInformation) Decode(r *aper.AperReader) error {
	// Decode logic for enum GNBCUUPOverloadInformation to be generated here.
	return fmt.Errorf("Decode not implemented for enum GNBCUUPOverloadInformation")
}
