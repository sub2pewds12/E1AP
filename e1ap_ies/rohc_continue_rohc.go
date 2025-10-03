package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// ROHCContinueROHC From: 9_4_5_Information_Element_Definitions.txt:2115
// ASN.1 Data Type: ENUMERATED
const (
	ROHCContinueROHCTrue aper.Enumerated = 0
)

type ROHCContinueROHC struct {
	Value aper.Enumerated
}
