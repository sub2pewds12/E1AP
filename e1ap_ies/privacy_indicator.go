package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// PrivacyIndicator From: 9_4_5_Information_Element_Definitions.txt:1926
// ASN.1 Data Type: ENUMERATED
const (
	PrivacyIndicatorImmediateMDT aper.Enumerated = 0
	PrivacyIndicatorLoggedMDT    aper.Enumerated = 1
)

type PrivacyIndicator struct {
	Value aper.Enumerated
}
