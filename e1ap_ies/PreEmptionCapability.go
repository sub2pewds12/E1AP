package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// PreEmptionCapability is a generated ENUMERATED type.
type PreEmptionCapability struct {
	Value aper.Enumerated
}

const (
	PreEmptionCapabilityShallNotTriggerPreEmption aper.Enumerated = 0
	PreEmptionCapabilityMayTriggerPreEmption      aper.Enumerated = 1
)

func (e *PreEmptionCapability) Encode(w *aper.AperWriter) error {
	// Encode logic for enum PreEmptionCapability to be generated here.
	return fmt.Errorf("Encode not implemented for enum PreEmptionCapability")
}

func (e *PreEmptionCapability) Decode(r *aper.AperReader) error {
	// Decode logic for enum PreEmptionCapability to be generated here.
	return fmt.Errorf("Decode not implemented for enum PreEmptionCapability")
}
