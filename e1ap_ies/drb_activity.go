package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// DRBActivity From: 9_4_5_Information_Element_Definitions.txt:419
// ASN.1 Data Type: ENUMERATED
const (
	DRBActivityActive    aper.Enumerated = 0
	DRBActivityNotActive aper.Enumerated = 1
)

type DRBActivity struct {
	Value aper.Enumerated
}
