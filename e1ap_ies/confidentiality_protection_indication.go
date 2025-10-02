package e1ap_ies

import "github.com/lvdund/ngap/aper"

// ConfidentialityProtectionIndication From: 9_4_5_Information_Element_Definitions.txt:251
const (
	ConfidentialityProtectionIndicationRequired  aper.Enumerated = 0
	ConfidentialityProtectionIndicationPreferred aper.Enumerated = 1
	ConfidentialityProtectionIndicationNotNeeded aper.Enumerated = 2
)

type ConfidentialityProtectionIndication struct {
	Value aper.Enumerated
}
