package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// UEActivity From: 9_4_5_Information_Element_Definitions.txt:2358
// ASN.1 Data Type: ENUMERATED
const (
	UEActivityActive    aper.Enumerated = 0
	UEActivityNotActive aper.Enumerated = 1
)

type UEActivity struct {
	Value aper.Enumerated
}
