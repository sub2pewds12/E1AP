package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// EHCCommonParametersEhcCIDLength is a generated ENUMERATED type.
type EHCCommonParametersEhcCIDLength struct {
	Value aper.Enumerated
}

const (
	EHCCommonParametersEhcCIDLengthBits7  aper.Enumerated = 0
	EHCCommonParametersEhcCIDLengthBits15 aper.Enumerated = 1
)
