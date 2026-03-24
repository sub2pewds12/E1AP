package e1ap_ies

import (
	"asn1go/per"
)

// TriggeringMessage is a generated ENUMERATED type.
type TriggeringMessage struct {
	Value int64
}

const (
	TriggeringMessageInitiatingMessage   int64 = 0
	TriggeringMessageSuccessfulOutcome   int64 = 1
	TriggeringMessageUnsuccessfulOutcome int64 = 2
)

// Encode implements the MessageEncoder interface for TriggeringMessage.
func (e *TriggeringMessage) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: false, RootValues: make([]int64, 3), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for TriggeringMessage.
func (e *TriggeringMessage) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: false, RootValues: make([]int64, 3), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
