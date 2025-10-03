package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// PDCPReestablishment From: 9_4_5_Information_Element_Definitions.txt:1590
// ASN.1 Data Type: ENUMERATED
const (
	PDCPReestablishmentTrue aper.Enumerated = 0
)

type PDCPReestablishment struct {
	Value aper.Enumerated
}
