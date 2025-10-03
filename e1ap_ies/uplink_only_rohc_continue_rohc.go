package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// UplinkOnlyROHCContinueROHC From: 9_4_5_Information_Element_Definitions.txt:2430
// ASN.1 Data Type: ENUMERATED
const (
	UplinkOnlyROHCContinueROHCTrue aper.Enumerated = 0
)

type UplinkOnlyROHCContinueROHC struct {
	Value aper.Enumerated
}
