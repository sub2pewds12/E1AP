package e1ap_ies

import "github.com/lvdund/ngap/aper"

// ConfidentialityProtectionResult From: 9_4_5_Information_Element_Definitions.txt:259
const (
	ConfidentialityProtectionResultPerformed    aper.Enumerated = 0
	ConfidentialityProtectionResultNotPerformed aper.Enumerated = 1
)

type ConfidentialityProtectionResult struct {
	Value aper.Enumerated
}
