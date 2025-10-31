package e1ap_ies

import (
	"fmt"

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

func (e *EarlyForwardingCOUNTReq) Encode(w *aper.AperWriter) error {
	// Encode logic for enum EarlyForwardingCOUNTReq to be generated here.
	return fmt.Errorf("Encode not implemented for enum EarlyForwardingCOUNTReq")
}

func (e *EarlyForwardingCOUNTReq) Decode(r *aper.AperReader) error {
	// Decode logic for enum EarlyForwardingCOUNTReq to be generated here.
	return fmt.Errorf("Decode not implemented for enum EarlyForwardingCOUNTReq")
}
