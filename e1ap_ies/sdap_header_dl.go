package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// SDAPHeaderDL is a generated ENUMERATED type.
type SDAPHeaderDL struct {
	Value aper.Enumerated
}

const (
	SDAPHeaderDLPresent aper.Enumerated = 0
	SDAPHeaderDLAbsent  aper.Enumerated = 1
)
