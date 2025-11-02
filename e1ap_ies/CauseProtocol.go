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

// Encode implements the aper.AperMarshaller interface.
func (e *CauseProtocol) Encode(w *aper.AperWriter) error {
	return w.WriteEnumerate(uint64(e.Value), aper.Constraint{Lb: 0, Ub: 6}, true)
}

// Decode implements the aper.AperUnmarshaller interface.
func (e *CauseProtocol) Decode(r *aper.AperReader) error {

	val, err := r.ReadEnumerate(aper.Constraint{Lb: 0, Ub: 6}, true)
	if err != nil {
		return err
	}
	e.Value = aper.Enumerated(val)
	return nil
}
