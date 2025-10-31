package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// PDCPSNStatusRequest is a generated ENUMERATED type.
type PDCPSNStatusRequest struct {
	Value aper.Enumerated
}

const (
	PDCPSNStatusRequestRequested aper.Enumerated = 0
)

func (e *PDCPSNStatusRequest) Encode(w *aper.AperWriter) error {
	// Encode logic for enum PDCPSNStatusRequest to be generated here.
	return fmt.Errorf("Encode not implemented for enum PDCPSNStatusRequest")
}

func (e *PDCPSNStatusRequest) Decode(r *aper.AperReader) error {
	// Decode logic for enum PDCPSNStatusRequest to be generated here.
	return fmt.Errorf("Decode not implemented for enum PDCPSNStatusRequest")
}
