package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// DLTXStop From: 9_4_5_Information_Element_Definitions.txt:413
// ASN.1 Data Type: ENUMERATED
const (
	DLTXStopStop   aper.Enumerated = 0
	DLTXStopResume aper.Enumerated = 1
)

type DLTXStop struct {
	Value aper.Enumerated
}
