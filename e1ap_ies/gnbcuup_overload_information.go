package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// GNBCUUPOverloadInformation is a generated ENUMERATED type.
type GNBCUUPOverloadInformation struct {
	Value aper.Enumerated
}

const (
	GNBCUUPOverloadInformationOverloaded    aper.Enumerated = 0
	GNBCUUPOverloadInformationNotOverloaded aper.Enumerated = 1
)
