package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// DataForwardingRequest is a generated ENUMERATED type.
type DataForwardingRequest struct {
	Value aper.Enumerated
}

const (
	DataForwardingRequestUL   aper.Enumerated = 0
	DataForwardingRequestDL   aper.Enumerated = 1
	DataForwardingRequestBoth aper.Enumerated = 2
)
