package e1ap_ies

import "github.com/lvdund/ngap/aper"

// EarlyForwardingCOUNTReq From: 9_4_5_Information_Element_Definitions.txt:975
const (
	EarlyForwardingCOUNTReqFirstDlCount aper.Enumerated = 0
	EarlyForwardingCOUNTReqDlDiscarding aper.Enumerated = 1
)

type EarlyForwardingCOUNTReq struct {
	Value aper.Enumerated
}
