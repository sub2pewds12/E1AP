package e1ap_ies

import (
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
