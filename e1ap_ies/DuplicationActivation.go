package e1ap_ies

import (
	"asn1go/per"
)

// DuplicationActivation is a generated ENUMERATED type.
type DuplicationActivation struct {
	Value int64
}

const (
	DuplicationActivationActive   int64 = 0
	DuplicationActivationInactive int64 = 1
)

// Encode implements the MessageEncoder interface for DuplicationActivation.
func (e *DuplicationActivation) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for DuplicationActivation.
func (e *DuplicationActivation) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
