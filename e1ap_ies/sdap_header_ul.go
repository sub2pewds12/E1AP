package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// SDAPHeaderUL From: 9_4_5_Information_Element_Definitions.txt:2214
// ASN.1 Data Type: ENUMERATED
const (
	SDAPHeaderULPresent aper.Enumerated = 0
	SDAPHeaderULAbsent  aper.Enumerated = 1
)

type SDAPHeaderUL struct {
	Value aper.Enumerated
}
