package e1ap_ies

import (
	"github.com/lvdund/ngap/aper"
)

// CauseProtocol From: 9_4_5_Information_Element_Definitions.txt:162
// ASN.1 Data Type: ENUMERATED
const (
	CauseProtocolTransferSyntaxError                          aper.Enumerated = 0
	CauseProtocolAbstractSyntaxErrorReject                    aper.Enumerated = 1
	CauseProtocolAbstractSyntaxErrorIgnoreAndNotify           aper.Enumerated = 2
	CauseProtocolMessageNotCompatibleWithReceiverState        aper.Enumerated = 3
	CauseProtocolSemanticError                                aper.Enumerated = 4
	CauseProtocolAbstractSyntaxErrorFalselyConstructedMessage aper.Enumerated = 5
	CauseProtocolUnspecified                                  aper.Enumerated = 6
)

type CauseProtocol struct {
	Value aper.Enumerated
}
