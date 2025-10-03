package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// GNBCUUPOverloadInformation From: 9_4_5_Information_Element_Definitions.txt:1219
// ASN.1 Data Type: ENUMERATED
const (
	GNBCUUPOverloadInformationOverloaded    aper.Enumerated = 0
	GNBCUUPOverloadInformationNotOverloaded aper.Enumerated = 1
)

type GNBCUUPOverloadInformation struct {
	Value aper.Enumerated
}
