package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// SDAPHeaderUL is a generated ENUMERATED type.
type SDAPHeaderUL struct {
	Value aper.Enumerated
}

const (
	SDAPHeaderULPresent aper.Enumerated = 0
	SDAPHeaderULAbsent  aper.Enumerated = 1
)
