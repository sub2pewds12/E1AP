package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// CauseProtocol is a generated ENUMERATED type.
type CauseProtocol struct {
	Value aper.Enumerated
}

const (
	CauseProtocolTransferSyntaxError                          aper.Enumerated = 0
	CauseProtocolAbstractSyntaxErrorReject                    aper.Enumerated = 1
	CauseProtocolAbstractSyntaxErrorIgnoreAndNotify           aper.Enumerated = 2
	CauseProtocolMessageNotCompatibleWithReceiverState        aper.Enumerated = 3
	CauseProtocolSemanticError                                aper.Enumerated = 4
	CauseProtocolAbstractSyntaxErrorFalselyConstructedMessage aper.Enumerated = 5
	CauseProtocolUnspecified                                  aper.Enumerated = 6
)
