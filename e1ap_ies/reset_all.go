package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// ResetAll From: 9_4_4_PDU_Definitions.txt:271
// ASN.1 Data Type: ENUMERATED
const (
	ResetAllResetAll aper.Enumerated = 0
)

type ResetAll struct {
	Value aper.Enumerated
}
