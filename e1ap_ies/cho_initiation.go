package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// CHOInitiation From: 9_4_5_Information_Element_Definitions.txt:232
// ASN.1 Data Type: ENUMERATED
const (
	CHOInitiationTrue aper.Enumerated = 0
)

type CHOInitiation struct {
	Value aper.Enumerated
}
