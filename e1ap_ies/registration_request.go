package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// RegistrationRequest From: 9_4_5_Information_Element_Definitions.txt:2081
// ASN.1 Data Type: ENUMERATED
const (
	RegistrationRequestStart aper.Enumerated = 0
	RegistrationRequestStop  aper.Enumerated = 1
)

type RegistrationRequest struct {
	Value aper.Enumerated
}
