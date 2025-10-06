package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// PDUSessionType is a generated ENUMERATED type.
type PDUSessionType struct {
	Value aper.Enumerated
}

const (
	PDUSessionTypeIpv4         aper.Enumerated = 0
	PDUSessionTypeIpv6         aper.Enumerated = 1
	PDUSessionTypeIpv4v6       aper.Enumerated = 2
	PDUSessionTypeEthernet     aper.Enumerated = 3
	PDUSessionTypeUnstructured aper.Enumerated = 4
)
