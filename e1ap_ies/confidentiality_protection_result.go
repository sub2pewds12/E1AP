package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// ConfidentialityProtectionResult is a generated ENUMERATED type.
type ConfidentialityProtectionResult struct {
	Value aper.Enumerated
}

const (
	ConfidentialityProtectionResultPerformed    aper.Enumerated = 0
	ConfidentialityProtectionResultNotPerformed aper.Enumerated = 1
)
