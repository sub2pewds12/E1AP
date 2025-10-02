package e1ap_ies

import "github.com/lvdund/ngap/aper"

// ActivityNotificationLevel From: 9_4_5_Information_Element_Definitions.txt:96
const (
	ActivityNotificationLevelDRB        aper.Enumerated = 0
	ActivityNotificationLevelPDUSession aper.Enumerated = 1
	ActivityNotificationLevelUe         aper.Enumerated = 2
)

type ActivityNotificationLevel struct {
	Value aper.Enumerated
}
