package e1ap_ies

import "github.com/lvdund/ngap/aper"

// IntegrityProtectionIndication From: 9_4_5_Information_Element_Definitions.txt:1241
const (
	IntegrityProtectionIndicationRequired  aper.Enumerated = 0
	IntegrityProtectionIndicationPreferred aper.Enumerated = 1
	IntegrityProtectionIndicationNotNeeded aper.Enumerated = 2
)

type IntegrityProtectionIndication struct {
	Value aper.Enumerated
}
