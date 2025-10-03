package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// DefaultDRB From: 9_4_5_Information_Element_Definitions.txt:385
// ASN.1 Data Type: ENUMERATED
const (
	DefaultDRBTrue  aper.Enumerated = 0
	DefaultDRBFalse aper.Enumerated = 1
)

type DefaultDRB struct {
	Value aper.Enumerated
}
