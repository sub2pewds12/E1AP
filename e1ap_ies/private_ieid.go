package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// PrivateIEID From: 9_4_6_Common_Definitions.txt:38
// ASN.1 Data Type: CHOICE
const (
	PrivateIEIDPresentNothing uint64 = iota
	PrivateIEIDPresentLocal
	PrivateIEIDPresentGlobal
)

type PrivateIEID struct {
	Choice uint64
	Local  *aper.Integer
	Global *string
}
