package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// MaxIPrate From: 9_4_5_Information_Element_Definitions.txt:1307
// ASN.1 Data Type: ENUMERATED
const (
	MaxIPrateBitrate64kbs aper.Enumerated = 0
	MaxIPrateMaxUErate    aper.Enumerated = 1
)

type MaxIPrate struct {
	Value aper.Enumerated
}
