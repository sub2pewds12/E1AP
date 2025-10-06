package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// IntegrityProtectionResult is a generated ENUMERATED type.
type IntegrityProtectionResult struct {
	Value aper.Enumerated
}

const (
	IntegrityProtectionResultPerformed    aper.Enumerated = 0
	IntegrityProtectionResultNotPerformed aper.Enumerated = 1
)
