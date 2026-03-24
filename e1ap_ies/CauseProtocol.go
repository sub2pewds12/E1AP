package e1ap_ies

import (
	"github.com/lvdund/asn1go/per"
)

// CauseProtocol is a generated ENUMERATED type.
type CauseProtocol struct {
	Value int64
}

const (
	CauseProtocolTransferSyntaxError                          int64 = 0
	CauseProtocolAbstractSyntaxErrorReject                    int64 = 1
	CauseProtocolAbstractSyntaxErrorIgnoreAndNotify           int64 = 2
	CauseProtocolMessageNotCompatibleWithReceiverState        int64 = 3
	CauseProtocolSemanticError                                int64 = 4
	CauseProtocolAbstractSyntaxErrorFalselyConstructedMessage int64 = 5
	CauseProtocolUnspecified                                  int64 = 6
)

// Encode implements the MessageEncoder interface for CauseProtocol.
func (e *CauseProtocol) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 7), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for CauseProtocol.
func (e *CauseProtocol) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 7), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
