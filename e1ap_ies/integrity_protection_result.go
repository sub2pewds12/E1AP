package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// IntegrityProtectionResult From: 9_4_5_Information_Element_Definitions.txt:1258
// ASN.1 Data Type: ENUMERATED
const (
	IntegrityProtectionResultPerformed    aper.Enumerated = 0
	IntegrityProtectionResultNotPerformed aper.Enumerated = 1
)

type IntegrityProtectionResult struct {
	Value aper.Enumerated
}
