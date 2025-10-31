package e1ap_ies

import (
	"fmt"

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

func (e *CauseProtocol) Encode(w *aper.AperWriter) error {
	// Encode logic for enum CauseProtocol to be generated here.
	return fmt.Errorf("Encode not implemented for enum CauseProtocol")
}

func (e *CauseProtocol) Decode(r *aper.AperReader) error {
	// Decode logic for enum CauseProtocol to be generated here.
	return fmt.Errorf("Decode not implemented for enum CauseProtocol")
}
