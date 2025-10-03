package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// Presence From: 9_4_6_Common_Definitions.txt:36
// ASN.1 Data Type: ENUMERATED
const (
	PresenceOptional    aper.Enumerated = 0
	PresenceConditional aper.Enumerated = 1
	PresenceMandatory   aper.Enumerated = 2
)

type Presence struct {
	Value aper.Enumerated
}
