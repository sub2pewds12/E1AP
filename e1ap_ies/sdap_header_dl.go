package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// SDAPHeaderDL From: 9_4_5_Information_Element_Definitions.txt:2208
// ASN.1 Data Type: ENUMERATED
const (
	SDAPHeaderDLPresent aper.Enumerated = 0
	SDAPHeaderDLAbsent  aper.Enumerated = 1
)

type SDAPHeaderDL struct {
	Value aper.Enumerated
}
