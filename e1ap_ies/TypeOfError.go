package e1ap_ies

import (
	"github.com/lvdund/asn1go/per"
)

// TypeOfError is a generated ENUMERATED type.
type TypeOfError struct {
	Value int64
}

const (
	TypeOfErrorNotUnderstood int64 = 0
	TypeOfErrorMissing       int64 = 1
)

// Encode implements the MessageEncoder interface for TypeOfError.
func (e *TypeOfError) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for TypeOfError.
func (e *TypeOfError) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
