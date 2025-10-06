package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// PrivacyIndicator is a generated ENUMERATED type.
type PrivacyIndicator struct {
	Value aper.Enumerated
}

const (
	PrivacyIndicatorImmediateMDT aper.Enumerated = 0
	PrivacyIndicatorLoggedMDT    aper.Enumerated = 1
)
