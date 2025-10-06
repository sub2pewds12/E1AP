package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// CNSupport is a generated ENUMERATED type.
type CNSupport struct {
	Value aper.Enumerated
}

const (
	CNSupportCEpc aper.Enumerated = 0
	CNSupportC5gc aper.Enumerated = 1
	CNSupportBoth aper.Enumerated = 2
)
