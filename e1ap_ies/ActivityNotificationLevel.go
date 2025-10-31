package e1ap_ies

import (
	"fmt"

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

func (e *ActivityNotificationLevel) Encode(w *aper.AperWriter) error {
	// Encode logic for enum ActivityNotificationLevel to be generated here.
	return fmt.Errorf("Encode not implemented for enum ActivityNotificationLevel")
}

func (e *ActivityNotificationLevel) Decode(r *aper.AperReader) error {
	// Decode logic for enum ActivityNotificationLevel to be generated here.
	return fmt.Errorf("Decode not implemented for enum ActivityNotificationLevel")
}
