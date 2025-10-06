package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// CauseTransport is a generated ENUMERATED type.
type CauseTransport struct {
	Value aper.Enumerated
}

const (
	CauseTransportUnspecified                  aper.Enumerated = 0
	CauseTransportTransportResourceUnavailable aper.Enumerated = 1
	CauseTransportUnknownTNLAddressForIAB      aper.Enumerated = 2
)
