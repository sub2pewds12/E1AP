package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// ConfidentialityProtectionIndication is a generated ENUMERATED type.
type ConfidentialityProtectionIndication struct {
	Value aper.Enumerated
}

const (
	ConfidentialityProtectionIndicationRequired  aper.Enumerated = 0
	ConfidentialityProtectionIndicationPreferred aper.Enumerated = 1
	ConfidentialityProtectionIndicationNotNeeded aper.Enumerated = 2
)
