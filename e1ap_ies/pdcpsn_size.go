package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// PDCPSNSize is a generated ENUMERATED type.
type PDCPSNSize struct {
	Value aper.Enumerated
}

const (
	PDCPSNSizeS12 aper.Enumerated = 0
	PDCPSNSizeS18 aper.Enumerated = 1
)
