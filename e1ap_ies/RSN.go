package e1ap_ies

import (
	"github.com/lvdund/asn1go/per"
)

// RSN is a generated ENUMERATED type.
type RSN struct {
	Value int64
}

const (
	RSNV1 int64 = 0
	RSNV2 int64 = 1
)

// Encode implements the MessageEncoder interface for RSN.
func (e *RSN) Encode(w *per.Encoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
	return w.EncodeEnumerated(int64(e.Value), c)
}

// Decode implements the MessageDecoder interface for RSN.
func (e *RSN) Decode(r *per.Decoder) error {

	c := per.EnumeratedConstraints{Extensible: true, RootValues: make([]int64, 2), ExtValues: nil}
	val, err := r.DecodeEnumerated(c)
	if err != nil {
		return err
	}
	e.Value = val
	return nil
}
