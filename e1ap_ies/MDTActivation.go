package e1ap_ies

import (
	"asn1go/per"
)

// MDTActivation is a generated ENUMERATED type.
type MDTActivation struct {
	Value int64
}

const (
	MDTActivationImmediateMDTOnly     int64 = 0
	MDTActivationImmediateMDTAndTrace int64 = 1
)

// Encode implements the MessageEncoder interface for MDTActivation.
func (e *MDTActivation) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for MDTActivation.
func (e *MDTActivation) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
