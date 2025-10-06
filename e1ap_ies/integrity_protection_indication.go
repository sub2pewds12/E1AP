package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// IntegrityProtectionIndication is a generated ENUMERATED type.
type IntegrityProtectionIndication struct {
	Value aper.Enumerated
}

const (
	IntegrityProtectionIndicationRequired  aper.Enumerated = 0
	IntegrityProtectionIndicationPreferred aper.Enumerated = 1
	IntegrityProtectionIndicationNotNeeded aper.Enumerated = 2
)
