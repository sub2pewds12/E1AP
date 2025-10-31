package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// PDCPDataRecovery is a generated ENUMERATED type.
type PDCPDataRecovery struct {
	Value aper.Enumerated
}

const (
	PDCPDataRecoveryTrue aper.Enumerated = 0
)

func (e *PDCPDataRecovery) Encode(w *aper.AperWriter) error {
	// Encode logic for enum PDCPDataRecovery to be generated here.
	return fmt.Errorf("Encode not implemented for enum PDCPDataRecovery")
}

func (e *PDCPDataRecovery) Decode(r *aper.AperReader) error {
	// Decode logic for enum PDCPDataRecovery to be generated here.
	return fmt.Errorf("Decode not implemented for enum PDCPDataRecovery")
}
