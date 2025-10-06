package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// Presence is a generated ENUMERATED type.
type Presence struct {
	Value aper.Enumerated
}

const (
	PresenceOptional    aper.Enumerated = 0
	PresenceConditional aper.Enumerated = 1
	PresenceMandatory   aper.Enumerated = 2
)
