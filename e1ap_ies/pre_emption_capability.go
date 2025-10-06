package e1ap_ies

import (
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
