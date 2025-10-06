package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// MDTActivation is a generated ENUMERATED type.
type MDTActivation struct {
	Value aper.Enumerated
}

const (
	MDTActivationImmediateMDTOnly     aper.Enumerated = 0
	MDTActivationImmediateMDTAndTrace aper.Enumerated = 1
)
