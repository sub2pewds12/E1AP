package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// TypeOfError From: 9_4_5_Information_Element_Definitions.txt:2314
// ASN.1 Data Type: ENUMERATED
const (
	TypeOfErrorNotUnderstood aper.Enumerated = 0
	TypeOfErrorMissing       aper.Enumerated = 1
)

type TypeOfError struct {
	Value aper.Enumerated
}
