package e1ap_ies

import (
	"fmt"

	"github.com/lvdund/ngap/aper"
)

// DirectForwardingPathAvailability is a generated ENUMERATED type.
type DirectForwardingPathAvailability struct {
	Value aper.Enumerated
}

const (
	DirectForwardingPathAvailabilityInterSystemDirectPathAvailable aper.Enumerated = 0
	DirectForwardingPathAvailabilityIntraSystemDirectPathAvailable aper.Enumerated = 1
)

func (e *DirectForwardingPathAvailability) Encode(w *aper.AperWriter) error {
	// Encode logic for enum DirectForwardingPathAvailability to be generated here.
	return fmt.Errorf("Encode not implemented for enum DirectForwardingPathAvailability")
}

func (e *DirectForwardingPathAvailability) Decode(r *aper.AperReader) error {
	// Decode logic for enum DirectForwardingPathAvailability to be generated here.
	return fmt.Errorf("Decode not implemented for enum DirectForwardingPathAvailability")
}
