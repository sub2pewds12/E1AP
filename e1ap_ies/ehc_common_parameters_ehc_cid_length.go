package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// EHCCommonParametersEhcCIDLength From: 9_4_5_Information_Element_Definitions.txt:977
// ASN.1 Data Type: ENUMERATED
const (
	EHCCommonParametersEhcCIDLengthBits7  aper.Enumerated = 0
	EHCCommonParametersEhcCIDLengthBits15 aper.Enumerated = 1
)

type EHCCommonParametersEhcCIDLength struct {
	Value aper.Enumerated
}
