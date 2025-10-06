package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// EarlyForwardingCOUNTReq is a generated ENUMERATED type.
type EarlyForwardingCOUNTReq struct {
	Value aper.Enumerated
}

const (
	EarlyForwardingCOUNTReqFirstDlCount aper.Enumerated = 0
	EarlyForwardingCOUNTReqDlDiscarding aper.Enumerated = 1
)
