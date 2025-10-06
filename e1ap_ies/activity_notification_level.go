package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// ActivityNotificationLevel is a generated ENUMERATED type.
type ActivityNotificationLevel struct {
	Value aper.Enumerated
}

const (
	ActivityNotificationLevelDRB        aper.Enumerated = 0
	ActivityNotificationLevelPDUSession aper.Enumerated = 1
	ActivityNotificationLevelUe         aper.Enumerated = 2
)
