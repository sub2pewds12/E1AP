package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// NewULTNLInformationRequired From: 9_4_5_Information_Element_Definitions.txt:1411
// ASN.1 Data Type: ENUMERATED
const (
	NewULTNLInformationRequiredRequired aper.Enumerated = 0
)

type NewULTNLInformationRequired struct {
	Value aper.Enumerated
}
