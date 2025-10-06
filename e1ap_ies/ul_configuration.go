package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// ULConfiguration is a generated ENUMERATED type.
type ULConfiguration struct {
	Value aper.Enumerated
}

const (
	ULConfigurationNoData aper.Enumerated = 0
	ULConfigurationShared aper.Enumerated = 1
	ULConfigurationOnly   aper.Enumerated = 2
)
