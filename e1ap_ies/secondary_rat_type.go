package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// SecondaryRATType From: manual_patch:-1
// ASN.1 Data Type: ENUMERATED
const (
	SecondaryRATTypeNr    aper.Enumerated = 0
	SecondaryRATTypeEUTRA aper.Enumerated = 1
)

type SecondaryRATType struct {
	Value aper.Enumerated
}
